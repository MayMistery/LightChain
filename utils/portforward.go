package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
)

func ConnectSSH() *ssh.Client {
	// 建立SSH连接
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("passwd"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", "remoteIP:22", config)
	if err != nil {
		fmt.Println("Dial ERROR:", err)
		return nil
	}
	return conn
}

func RemotePortForword(ini Ini) {
	conn := ConnectSSH()

	// 远程端口转发:remote_server:8080到localhost:80
	go func() {
		listener, err := conn.Listen("tcp", "0.0.0.0:6666")
		if err != nil {
			fmt.Println("Listen ERROR:", err)
			return
		}

		for {
			remoteConn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept ERROR:", err)
				continue
			}
			go func(remoteConn net.Conn) {
				localConn, err := net.Dial("tcp", "127.0.0.1:9000")
				if err != nil {
					fmt.Println("Dial ERROR:", err)
					return
				}
				defer localConn.Close()
				fmt.Println("Forwarded connection from", remoteConn.RemoteAddr(), " to ", localConn.RemoteAddr())
				handleForwardedConn(remoteConn, localConn)
			}(remoteConn)
		}
	}()

	// 保持SSH连接
	conn.Wait()
}

func LocalPortForward() {

	conn := ConnectSSH()

	//本地端口转发:localhost:8080到remote_server:80
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Listen ERROR:", err)
		return
	}
	go func() {
		for {
			localConn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept ERROR:", err)
				continue
			}
			go func(localConn net.Conn) {
				remoteConn, err := conn.Dial("tcp", "127.0.0.1:80")
				if err != nil {
					fmt.Println("Dial ERROR:", err)
					return
				}
				defer remoteConn.Close()
				fmt.Println("Forwarded connection from", localConn.RemoteAddr(), "to", remoteConn.RemoteAddr())
				handleForwardedConn(localConn, remoteConn)
			}(localConn)
		}
	}()
}

func DynamicPortForward() {

}

func handleForwardedConn(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()

	// 建立goroutine转发src->dest和dest->src的数据
	go func() {
		if _, err := io.Copy(dest, src); err != nil {
			fmt.Println(err)
		}
	}()
	if _, err := io.Copy(src, dest); err != nil {
		fmt.Println(err)
	}
}
