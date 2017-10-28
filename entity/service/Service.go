package entity

import (
	//"fmt"
)

type Service struct {
	AgendaStorage *Storage
}

type ServiceStat int
const (
	LOGIN ServiceStat = 
	LOGOUT ServiceStat
)

// ??开启Agenda
// 获取单例的storage
// 获取数据文件各种信息--start失败false&返回报错信息&未登录
// 成功则判断是否返回true&报错信息&是否登陆状态
func StartAgenda(service *Service) (bool, StorageError, ServiceStat) {
	service.AgendaStorage = GetStorageInstance()
	return false
}

// ?? 退出Agenda
func (service *Service) QuitAgenda() {
	// return false
}

// 用户登陆
func (service *Service) UserLogin(userName string, password string) (bool, StorageError) {
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

// 列出用户
func (service *Service) ListAllUsers(userName string, password string) []User {
	return []User{}
}

// 创建会议
func (service *Service) CreateMeeting(sponsor string, title string, 
					startDate string, endDate string, participators []User) bool {
	return false
}

// 查找会议---通过title查找
func (service *Service) MeetingQueryByTitle(userName string, title string) []Meeting {
	return []Meeting{}
}

// 查找会议---通过usernsme(作为会议发起者和参与者)和会议起始时间查找
func (service *Service) MeetingQueryByUserAndTime(userName string, startDate string, endDate string) []Meeting {
	return []Meeting{}
}

// 列出该用户发起或参与的所有会议
func (service *Service) ListAllMeetings(userName string) []Meeting {
	return []Meeting{}
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


