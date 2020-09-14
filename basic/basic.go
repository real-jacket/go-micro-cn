package basic

import (
	"github.com/go-micro-cn/tutorials/basic/config"
	"github.com/go-micro-cn/tutorials/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
