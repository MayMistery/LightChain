package server

import (
	"MayMistery/LightChain/cmd"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func FileServer(cfg cmd.Config) {
	// 获取当前执行文件目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	rootPath := filepath.Join(wd, "file")

	// 如果file目录不存在,则创建
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		err := os.MkdirAll(rootPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	http.Handle("/", http.FileServer(http.Dir(rootPath)))

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	log.Printf("Listening on %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
