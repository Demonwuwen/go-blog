package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi)Index(w http.ResponseWriter,r *http.Request)  {
	index := common.Template.Index

	//页面涉及所有数据必须有定义
	//数据查询
	if err := r.ParseForm(); err != nil{
		log.Println("表单获取失败：",err)
		index.WriteError(w, errors.New("表单获取失败"))
		return
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != ""{
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示数量
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path,"/")

	hr, err := service.GetAllIndexInfo(slug,page,pageSize)

	//hr,err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("Index 获取数据出错",err)
		index.WriteError(w, errors.New("系统错误， 请联系管理员!!"))
	}
	index.WriteData(w, hr)
}