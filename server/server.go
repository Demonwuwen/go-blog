package server

import (
	"go-blog/router"
	"log"
	"net/http"
)

var App = &MsServer{} 
type MsServer struct {
	
}

func (*MsServer)Start(ip, port string)  {
	server := http.Server{
		Addr: ip+":"+port,
	}
	router.Router()
	if err := server.ListenAndServe();err != nil{
		log.Println(err)
	}
}