package dal

import (
	"github.com/S240/AetherCloud/app/uploadfile/biz/dal/mysql"
	"github.com/S240/AetherCloud/app/uploadfile/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
