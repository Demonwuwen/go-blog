package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)


type IndexData struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
}




func Router()  {
	//1.页面 2.数据(json)  3.i
	http.HandleFunc("/",views.HTML.Index)


	//http://localhost:8088/c/1 1参数分类
	http.HandleFunc("/c/",views.HTML.Category)
	http.HandleFunc("/login",views.HTML.Login)
	http.HandleFunc("/p/",views.HTML.Detail)
	http.HandleFunc("/writing",views.HTML.Writing)
	http.HandleFunc("/pigeonhole",views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post",api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/",api.API.GetPost)
	http.HandleFunc("/api/v1/post/search",api.API.SearchPost)
	http.HandleFunc("//api/v1/qiniu/token",api.API.QiniuToken)
	http.HandleFunc("/api/v1/login",api.API.Login)
	http.Handle("/resource/",http.StripPrefix("/resource/",
		http.FileServer(http.Dir("public/resource"))))
}
