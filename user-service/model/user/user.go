package user

import (
	"fmt"
	"sync"

	proto "github.com/go-micro-cn/tutorials/user-service/proto/user"
)

var (
	s *service
	m sync.RWMutex
)

// service 服务
type service struct{}

// Service 用户服务类
type Service interface {
	// QueryUserByName 根据用户名获取用户
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
