package service
import (
	"github.com/txzdream/agenda-go/entity/model"
	"github.com/txzdream/agenda-go/entity/storage"
)

type Service struct {
	service *storage.Storage
}

// 用户登陆
func (service *Service) UserLogin(userName string, password string) bool {
	return false
}

// 用户注册
func (service *Service) UserRegister(userName string, password string) bool {
	return false
}

// 删除用户
func (service *Service) DeleteUser(userName string, password string) bool {
	return false
}

// 列出会议
func (service *Service) ListAllUsers(userName string, password string) []model.User {
	return []model.User{}
}

// 创建会议
