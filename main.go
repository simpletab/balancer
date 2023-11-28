// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net"

	"github.com/zehuamama/balancer/proxy"
)

func main() {
	//config, err := ReadConfig("config.yaml")
	//if err != nil {
	//	log.Fatalf("read config error: %s", err)
	//}
	//
	//err = config.Validation()
	//if err != nil {
	//	log.Fatalf("verify config error: %s", err)
	//}

	// 四层代理
	s := &proxy.Server{}
	li, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal(err)
	}
	s.Li = li
	s.Balance = proxy.NewRoundRobin()
	s.Balance.AddNode("10.244.182.4:80", "nginx1")
	s.Balance.AddNode("10.244.182.5:80", "nginx2")
	s.Balance.AddNode("10.244.235.133:80", "nginx3")
	s.Run()

	//router := mux.NewRouter()
	//for _, l := range config.Location {
	//	httpProxy, err := proxy.NewHTTPProxy(l.ProxyPass, l.BalanceMode)
	//	if err != nil {
	//		log.Fatalf("create proxy error: %s", err)
	//	}
	//	// start health check
	//	if config.HealthCheck {
	//		httpProxy.HealthCheck(config.HealthCheckInterval)
	//	}
	//	router.Handle(l.Pattern, httpProxy)
	//}
	//if config.MaxAllowed > 0 {
	//	router.Use(maxAllowedMiddleware(config.MaxAllowed))
	//}
	//svr := http.Server{
	//	Addr:    ":" + strconv.Itoa(config.Port),
	//	Handler: router,
	//}
	//
	//// print config detail
	//config.Print()
	//
	//// listen and serve
	//if config.Schema == "http" {
	//	err := svr.ListenAndServe()
	//	if err != nil {
	//		log.Fatalf("listen and serve error: %s", err)
	//	}
	//} else if config.Schema == "https" {
	//	err := svr.ListenAndServeTLS(config.SSLCertificate, config.SSLCertificateKey)
	//	if err != nil {
	//		log.Fatalf("listen and serve error: %s", err)
	//	}
	//}
}

//一、启动命令改成：'bash', '-c', "while true;do echo 1;sleep 60s;done"，可进入容器查看容器内目录结构。
//参考：https://www.cnblogs.com/levinyinyc/p/15209947.html
