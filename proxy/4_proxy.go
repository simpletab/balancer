package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
)

type TCPProxy struct {
}

type Server struct {
	Li      net.Listener
	Balance Policy
}

func (s *Server) Run() {
	for {
		c, err := s.Li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			remoteAddr := c.RemoteAddr()
			backendIp := s.Balance.PickNode(remoteAddr.String())
			serverConn, err := net.Dial("tcp", backendIp) // 自己主动发起新的连接到后端 server
			if err != nil {
				log.Fatal(err)
				c.Close()
				return
			}
			fmt.Println("获取到了新连接", remoteAddr, backendIp)
			go func() {
				_, err := io.Copy(serverConn, c)
				if err != nil {
					fmt.Println(err, 1)
				}
				c.Close()
				serverConn.Close()
				fmt.Println("结束1", err)
			}()
			go func() {
				_, err := io.Copy(c, serverConn)
				if err != nil {
					fmt.Println(err, 2)
				}
				c.Close()
				serverConn.Close()
				fmt.Println("结束2", err)
			}()
		}(c)
	}
}
