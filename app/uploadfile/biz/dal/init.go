package dal

import (
	"github.com/S240/AetherCloud/app/uploadfile/biz/dal/mysql"
	"github.com/S240/AetherCloud/app/uploadfile/biz/dal/redis"
	"github.com/S240/AetherCloud/app/uploadfile/biz/dal/minio"
)

func Init() {
	redis.Init()
	mysql.Init()
	minio.Init()
}
