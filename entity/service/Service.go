package entity

import (
	model "github.com/txzdream/agenda-go/entity/model"
	storage "github.com/txzdream/agenda-go/entity/storage"
)

type Service struct {
	AgendaStorage *storage.Storage
}

// 开启Agenda，获取单例的storage
// 获取数据文件各种信息 --- 返回true/false&返回报错信息
// Usage: var k service.Service  service.StartAgenda(&k)
func StartAgenda(service *Service) (bool, storage.StorageError) {
	service.AgendaStorage = storage.GetStorageInstance()
	return service.AgendaStorage.ReadFromDataFile()
}

// 判断是否能够自动登陆，即curUser中是否存在有效用户名
// 存在有效用户名，返回true+用户名，否则返回false+空字符串
func AutoUserLogin(service *Service) () (bool, string) {
	result, currentUserName := service.AgendaStorage.ReadFromCurrentUserFile
	if !result {
		return false, ""
	}
	// 判断是否存在该用户名
	if len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.UserName == currentUserName
	})) == 1 {
		return true, currentUserName
	}
	return false, ""
}

// 获取当前用户名


// ?? 退出Agenda
func (service *Service) QuitAgenda() {
	// return false
}

// 用户登陆
func (service *Service) UserLogin(userName string, password string) (bool, storage.StorageError) {
	return service.AgendaStorage.ReadFromDataFile()
}

// 用户注册
func (service *Service) UserRegister(userName string, password string) bool {
	return false
}

// 删除用户
func (service *Service) DeleteUser(userName string, password string) bool {
	return false
}

// 列出用户
func (service *Service) ListAllUsers(userName string, password string) []model.User {
	return []model.User{}
}

// 创建会议
func (service *Service) CreateMeeting(sponsor string, title string, 
					startDate string, endDate string, participators []model.User) bool {
	return false
}

// 查找会议---通过title查找
func (service *Service) MeetingQueryByTitle(userName string, title string) []model.Meeting {
	return []model.Meeting{}
}

// 查找会议---通过usernsme(作为会议发起者和参与者)和会议起始时间查找
func (service *Service) MeetingQueryByUserAndTime(userName string, startDate string, endDate string) []model.Meeting {
	return []model.Meeting{}
}

// 列出该用户发起或参与的所有会议
func (service *Service) ListAllMeetings(userName string) []model.Meeting {
	return []model.Meeting{}
}

// 列出该用户发起的所有会议
func (service *Service) ListAllSponsorMeetings(userName string, password string) bool {
	return false
}

// 列出该用户参加的所有会议
func (service *Service) ListAllParticipateMeetings(userName string, password string) bool {
	return false
}

// 删除发起者sponsor题目title会议
func (service *Service) DeleteMeeting(sponsor string, title string) bool {
	return false
}

// 删除sponsor所有会议
func (service *Service) DeleteAllMeetings(sponsor string) bool {
	return false
}


