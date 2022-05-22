package api

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-blog/common"
	"go-blog/config"
	"net/http"
)

// 存储相关功能的引入包只有这两个，后面不再赘述

func (*Api) QiniuToken(w http.ResponseWriter, r *http.Request) {
	//自定义凭证有效期（示例2小时，Expires单位为秒，为上传凭证的有效时间）
	bucket := "demon-goblog"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(w, upToken)

}
