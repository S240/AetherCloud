package dal

import (
	"github.com/S240/AetherCloud/app/apigw/biz/dal/mysql"
	"github.com/S240/AetherCloud/app/apigw/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
