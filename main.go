package main

import (
	"MayMistery/LightChain/cmd"
	"MayMistery/LightChain/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type progressWriter struct {
	Total   int64
	Writer  io.Writer
	Current int64
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	n, err := pw.Writer.Write(p)
	pw.Current += int64(n)

	fmt.Fprintf(pw.Writer, "%d/%d\n", pw.Current, pw.Total)

	return n, err
}

func main() {
	var cfg cmd.Config

	cmd.Flag(&cfg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			utils.IndexHtml(w)
		}

		url := r.URL.Query().Get("r")
		if url == "" {
			w.WriteHeader(400)
			return
		}

		tmpDir := "./tmp"
		if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
			// 不存在则创建
			if err := os.MkdirAll(tmpDir, 0755); err != nil {
				w.WriteHeader(500)
				return
			}
		}

		// 生成临时文件名
		tmpFile := filepath.Join(tmpDir, filepath.Base(url))

		// 下载url到临时文件
		resp, err := http.Get(url)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer resp.Body.Close()

		contentLength := resp.ContentLength

		f, err := os.Create(tmpFile)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer f.Close()

		// 给客户端发送下载进度
		w.Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)

		progress := &progressWriter{
			Total:  resp.ContentLength,
			Writer: io.MultiWriter(w, f),
		}

		_, err = io.Copy(progress, resp.Body)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// 发送文件给客户端
		http.ServeFile(w, r, tmpFile)

		log.Fatal(http.ListenAndServe(":8080", nil))

		//path := r.URL.Path[1:]
		//f, err := os.Open(cfg.Dir + "/" + path)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//defer f.Close()
		//
		//io.Copy(w, f)
	})

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	log.Printf("Listening on %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
