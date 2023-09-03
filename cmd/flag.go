package cmd

import (
	"flag"
	"log"
	"net"
	"strconv"
)

func Flag(cfg *Config) {
	flag.StringVar(&cfg.Dir, "d", ".", "directory to serve")
	flag.StringVar(&cfg.Host, "h", "0.0.0.0", "listen host")
	flag.StringVar(&cfg.hostNet, "n", "", "listen network interface")

	flag.Parse()

	flag.IntVar(&cfg.Port, "p", 9999, "listen port")

	args := flag.Args()
	if len(args) > 0 {
		portStr := args[len(args)-1]
		usrPort, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatal("Invalid port number")
		}
		cfg.Port = usrPort
	}

	if cfg.hostNet != "" {
		ifi, err := net.InterfaceByName(cfg.hostNet)
		if err != nil {
			log.Fatal(err)
		}

		addrs, err := ifi.Addrs()

		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok { //&& !ipnet.IP.IsLoopback()
				if ipnet.IP.To4() != nil {
					cfg.Host = ipnet.IP.String()
					// 第一个IPv4地址
					break
				}
			}
		}
	}
}
