package api

import (
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//接收用户名和密码 返回对应json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	log.Println("username:",userName)
	log.Println("passwd:",passwd)
	loginRes,err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w,err)
		return
	}
	common.Success(w,loginRes)
}
