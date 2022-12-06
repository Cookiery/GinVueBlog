package qiniu

import (
	"context"
	"main/commond/errmsg"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// UploadFile
func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false, // Need money, don't use
		UseHTTPS:      false, // Need money, don't use
	}
	putExtra := storage.PutExtra{}
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := uploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}

	return QiniuServer + ret.Key, errmsg.SUCCSE
}
