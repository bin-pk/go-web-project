package main

import (
	"go-web-project/controller"
	"go-web-project/core"
	"go-web-project/repos"
	"go-web-project/service"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	cacheCtx := core.CreateCacheContext()

	duckDB := repos.CuteDuckDB{}
	if err := duckDB.Create(repos.CuteDBConfig{
		DriverName: "duckdb",
		DBFilePath: "temp.db",
		Ipaddr:     nil,
		Port:       0,
		Username:   "",
		Password:   "",
		Timeout:    0,
	}); err != nil {
		log.Println("Error:", err)
	}

	// service 등록
	helloService := &service.HelloService{}

	servMap := service.CreateServiceMap()
	servMap.AddService("hello", helloService)

	// 등록한 service 를 웹서버에 등록
	serv := controller.CreateWebServer(controller.CuteWebServerConfig{Port: 8080}, cacheCtx)
	serv.AddServiceMap(servMap)
	serv.AddServiceHandler("/task", controller.GetCode, &service.HelloInput{})

	// static page 등록
	serv.AddStaticFile("/", filepath.Join(dir, "viewer", "index.html"))

	// web server 시작.
	if err := serv.Start(); err != nil {
		log.Fatalln(err)
	}
}
