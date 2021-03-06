package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init()  {
	//模版加载
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8088",
	}
	router.Router()
	if err := server.ListenAndServe();err != nil{
		log.Println(err)
	}
}
