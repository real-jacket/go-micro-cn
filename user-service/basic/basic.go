package basic

import (
	"github.com/go-micro-cn/tutorials/user-service/basic/config"
	"github.com/go-micro-cn/tutorials/user-service/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
