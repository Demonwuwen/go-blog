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

func (*HTMLApi)Category(w http.ResponseWriter,r *http.Request)  {
	categoryTemplate := common.Template.Category
	//http://localhost:8088/c/1 1参数分类
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path,"/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w,errors.New("路径不匹配"))
		return
	}

	if err := r.ParseForm(); err != nil{
		log.Println("表单获取失败：",err)
		categoryTemplate.WriteError(w, errors.New("表单获取失败"))
		return
	}

	pageStr := r.Form.Get("page")
	if pageStr == ""{
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示数量
	pageSize := 10
	categoryResponse,err := service.GetPostsByCategoryId(cId,page,pageSize)
	if err != nil {
		categoryTemplate.WriteError(w,err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
