package api

import (
	"errors"
	"go-blog/common"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/service"
	"go-blog/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		common.Template.Writing.WriteError(w, err)
		return
	}
	post, err := dao.GetPostById(pid)
	common.Success(w,post)

}
func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//	获取用户id,判断是否登陆
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登陆已过期"))
		return
	}
	uid := claim.Uid
	//Post save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cid := params["categoryId"].(string)
		categoryId, err := strconv.Atoi(cid)
		if err != nil {
			log.Println("转换字符到int失败")
		}
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		ptype := int(postType)
		post := &models.Post{
			//Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       ptype,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		//update
		params := common.GetRequestJsonParam(r)
		cidFloat := params["categoryId"].(float64)
		categoryId:= int(cidFloat)
		if err != nil {
			log.Println("转换字符到int失败")
		}
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		ptype := int(postType)
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       ptype,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)

	}
	common.GetRequestJsonParam(r)
}
