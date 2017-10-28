package entity

import (
	"fmt"
	model "github.com/txzdream/agenda-go/entity/model"
	storage "github.com/txzdream/agenda-go/entity/storage"
	tools "github.com/txzdream/agenda-go/entity/tools"
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

// 判断是否能够自动登陆，即curUser中是否存在可登录的用户名
// 存在有效用户名，返回true+用户名，否则返回false+空字符串
func (service *Service) AutoUserLogin() (bool, string) {
	result, currentUserName := service.AgendaStorage.ReadFromCurrentUserFile()
	if !result {
		fmt.Println(result, currentUserName)
		return false, ""
	}
	// 根据获得的用户名判断是否存在该用户名
	if len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == currentUserName
	})) == 1 {
		return true, currentUserName
	}
	return false, ""
}

// ?? 退出Agenda
func (service *Service) QuitAgenda() {
	// return false
}

// 没有可自动登陆的用户，需要用户输入用户名和密码登陆
func (service *Service) UserLogin(userName string, password string) bool {
	// 对传入的md5密码进行加密
	password = tools.MD5Encryption(password)
	// 根据获得的用户名和密码判断是否存在该用户名
	return len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	})) == 1
}

// 用户注册，用户输入信息，判断是否存在同名用户
func (service *Service) UserRegister(userName string, password string, 
									 email string, phone string) bool {
	// 根据获得的用户名判断是否存在该用户名
	if len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == userName
	})) == 1 {
		return false		// 已存在同名用户则注册失败
	}
	// 对传入的md5密码进行加密
	password = tools.MD5Encryption(password)
	return service.AgendaStorage.CreateUser(model.User{UserName: userName, Password: password, Email: email, Phone: phone})
}

// 删除用户，判断是否存在同名用户再进行删除
func (service *Service) DeleteUser(userName string, password string) bool {
	password = tools.MD5Encryption(password)
	// 根据获得的用户名判断是否存在该用户名
	if len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	})) != 1 {
		return false		// 不存在同名用户则删除失败
	}
	// 存在同名用户则进行删除操作
	return service.AgendaStorage.DeleteUser(func(user model.User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	})
}

// 列出所有用户
func (service *Service) ListAllUsers() []model.User {
	return service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return true
	})
}


// 判断会议起始时间是否合法
func IsValidStartAndEndDateTime(startDateString *string, endDateString *string) bool {
	// 判断时间字符串是否符合格式要求：2012-2-2/11:23并解析字符串为Int数组
	var startDateIntArray [5]int
	var endDateIntArray [5]int
	if !model.StringDateTimeToIntArray(*startDateString, &startDateIntArray) || 
		!model.StringDateTimeToIntArray(*endDateString, &endDateIntArray) {
		return false
	}

	// 判断时间数字是否合法
	if !model.IsValidDateTime(startDateIntArray) || !model.IsValidDateTime(endDateIntArray) {
		return false
	}

	// 判断开始时间是否小于结束时间
	startDate := model.Date{model.SetDateByYMDHM(startDateIntArray)}
	endDate := model.Date{model.SetDateByYMDHM(endDateIntArray)}
	if !startDate.Before(endDate) {
		return false
	}

	// 再一次由于对原有字符串有容错性，再次转换确保字符串格式规范
	*startDateString = startDate.ToString()
	*endDateString = endDate.ToString()
	return true
}

// 创建会议
func (service *Service) CreateMeeting(sponsor string, title string, 
					startDateString string, endDateString string, participators []string) bool {
	// 判断title是否已存在
	if len(service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetTitle() == title
	})) > 0 {
		return false
	}

	// 判断时间合法
	if !IsValidStartAndEndDateTime(&startDateString, &endDateString) {
		return false
	}

	// ----- 判断参与者是否有其他同时段的会议 -----
	// 获取同时段冲突会议
	timeConflictMeetings := service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return (meeting.GetStartDate() <= startDateString && meeting.GetEndDate() > startDateString) ||
		(meeting.GetStartDate() < endDateString && meeting.GetEndDate() >= endDateString) ||
		(meeting.GetStartDate() >= startDateString && meeting.GetEndDate() <= endDateString)
	})
	for _, tMeeting := range timeConflictMeetings {
		if tMeeting.IsParticipators(sponsor) || tMeeting.GetSponsor() == sponsor {
			return false
		}
		for _, participator := range participators {
			if tMeeting.IsParticipators(participator) || tMeeting.GetSponsor() == participator {
				return false
			}
		}
	}
	// ------------------------------------------

	// 创建会议
	return true
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


