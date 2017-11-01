package service

import (
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

// 退出Agenda：将单例指针置nil
func (service *Service) QuitAgenda() {
	service.AgendaStorage = nil
}

// 判断是否能够自动登陆，即curUser中是否存在可登录的用户名
// 存在有效用户名，返回true+用户名，否则返回false+空字符串
func (service *Service) AutoUserLogin() (bool, string) {
	result, currentUserName := service.AgendaStorage.ReadFromCurrentUserFile()
	if !result {
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

// 没有可自动登陆的用户，需要用户输入用户名和密码登陆
func (service *Service) UserLogin(userName string, password string) bool {
	// 对传入的md5密码进行加密
	password = tools.MD5Encryption(password)
	// 根据获得的用户名和密码判断是否存在该用户名
	if len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	})) != 1 {
		return false
	}
	// 将当前用户名写入curUser.txt
	return service.AgendaStorage.WriteToCurrentUserFile(userName)
}

// 用户退出登录，把空字符串写进curuser.txt
func (service *Service) UserLogout() bool {
	return service.AgendaStorage.WriteToCurrentUserFile("")
}

// 用户注册，用户输入信息，判断是否存在同名用户
func (service *Service) UserRegister(userName string, password string,
	email string, phone string) bool {
	// 根据获得的用户名判断是否存在该用户名
	if len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == userName
	})) == 1 {
		return false // 已存在同名用户则注册失败
	}
	// 对传入的md5密码进行加密
	password = tools.MD5Encryption(password)
	return service.AgendaStorage.CreateUser(
		model.User{UserName: userName, Password: password, Email: email, Phone: phone})
}

// 删除用户，判断是否存在同名用户再进行删除
func (service *Service) DeleteUser(userName string, password string) bool {
	password = tools.MD5Encryption(password)
	// 存在同名用户则进行删除操作
	if !service.AgendaStorage.DeleteUser(func(user model.User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	}) {
		return false
	}
	// 删除所有发起会议
	service.DeleteAllMeetings(userName)
	// 退出所有参与会议并删除参与人数为0的会议
	meetings := service.ListAllParticipateMeetings(userName)
	for _, meeting := range meetings {
		service.QuitMeeting(userName, meeting.GetTitle())
	}
	return true
}

// 获取当前用户
func (service *Service) GetCurrentUser(currentUserName string) (bool, model.User) {
	users := service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == currentUserName
	})
	if len(users) != 1 {
		return false, model.User{}
	}
	return true, users[0]
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
	startDate := model.Date{DateTime: model.SetDateByYMDHM(startDateIntArray)}
	endDate := model.Date{DateTime: model.SetDateByYMDHM(endDateIntArray)}
	if !startDate.Before(endDate) {
		return false
	}

	// 再一次由于对原有字符串有容错性，再次转换确保字符串格式规范
	*startDateString = startDate.ToString()
	*endDateString = endDate.ToString()
	return true
}

// 判断单个用户是否已注册
func (service *Service) IsRegisteredUser(userName string) bool {
	return len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == userName
	})) == 1
}

// 判断多个用户是否都已注册
func (service *Service) IsRegisteredUsers(userNames []string) bool {
	return len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		for _, userName := range userNames {
			if userName == user.GetUserName() {
				return true
			}
		}
		return false
	})) == len(userNames) // 判断人数是否一致
}

// 获取时间冲突会议
func (service *Service) GetTimeConflictMeetings(startDateString string, endDateString string) []model.Meeting {
	return service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return (meeting.GetStartDate() <= startDateString && meeting.GetEndDate() > startDateString) ||
			(meeting.GetStartDate() < endDateString && meeting.GetEndDate() >= endDateString) ||
			(meeting.GetStartDate() >= startDateString && meeting.GetEndDate() <= endDateString)
	})
}

// 创建会议 - 检查title是否唯一、时间是否合法、参与者和发起者是否可参加会议椅
func (service *Service) CreateMeeting(sponsor string, title string,
	startDateString string, endDateString string, participators []string) bool {
	// 判断title是否已存在 / 判断时间合法 / 判断发起者和参与者是否都已注册
	if len(service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetTitle() == title
	})) > 0 || !IsValidStartAndEndDateTime(&startDateString, &endDateString) ||
		!service.IsRegisteredUser(sponsor) || !service.IsRegisteredUsers(participators) {
		return false
	}

	// 检查参与者是否与发起者是同一人
	for _, participator := range participators {
		if sponsor == participator {
			return false
		}
	}

	// 获取同时段冲突会议
	timeConflictMeetings := service.GetTimeConflictMeetings(startDateString, endDateString)

	// 判断发起者或参与者是否参与了冲突会议
	for _, tMeeting := range timeConflictMeetings {
		if tMeeting.IsParticipators(sponsor) || tMeeting.GetSponsor() == sponsor { // 检查发起者
			return false
		}
		for _, participator := range participators { // 检查参与者
			if tMeeting.IsParticipators(participator) || tMeeting.GetSponsor() == participator {
				return false
			}
		}
	}

	// 创建会议
	return service.AgendaStorage.CreateMeeting(
		model.Meeting{Title: title, Sponsor: sponsor, Participators: participators, StartDate: startDateString, EndDate: endDateString})
}

// 发起者增加会议参与者 -- 判断参与者是否可参加
func (service *Service) AddParticipatorByTitle(sponsor string, title string, participator string) bool {
	meetings := service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetTitle() == title && meeting.GetSponsor() == sponsor
	})
	// 判断该会议是否存在 / 判断参与者是否已注册 / 参与者就是发起者本人
	if len(meetings) == 0 || !service.IsRegisteredUser(participator) || participator == sponsor {
		return false
	}

	meeting := meetings[0] // 获取该会议
	// 获取同时段冲突会议
	timeConflictMeetings := service.GetTimeConflictMeetings(meeting.GetStartDate(), meeting.GetEndDate())
	// 判断参与者是否参与了冲突会议
	for _, tMeeting := range timeConflictMeetings {
		if tMeeting.IsParticipators(participator) || tMeeting.GetSponsor() == participator {
			return false
		}
	}

	// 增加参与者
	if !meeting.AddParticipator(participator) {
		return false
	}

	// 更新会议内容
	return service.AgendaStorage.UpdateMeeting(func(pMeeting model.Meeting) bool {
		return pMeeting.GetTitle() == meeting.GetTitle()
	}, meeting)
}

// 发起者删除会议参与者 -- 判断删除该参与者是否成功，删除后会议人数是否为0
func (service *Service) DeleteParticipatorByTitle(sponsor string, title string, participator string) bool {
	meetings := service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetTitle() == title && meeting.GetSponsor() == sponsor
	})

	// 判断该会议是否存在 / 删除该参与者是否成功
	if len(meetings) == 0 || !meetings[0].DeleteParticipator(participator) {
		return false
	}

	// 判断成功删除后会议参与者人数
	if meetings[0].GetParticipatorsNumber() > 0 { // 还有参与者，更新会议
		return service.AgendaStorage.UpdateMeeting(func(meeting model.Meeting) bool {
			return meeting.GetTitle() == meetings[0].GetTitle()
		}, meetings[0])
	} else { // 没有参与者，删除会议
		return service.AgendaStorage.DeleteMeetings(func(meeting model.Meeting) bool {
			return meeting.GetTitle() == meetings[0].GetTitle()
		})
	}
}

// 查询会议---通过用户名(用户作为发起者/参与者)和会议title查找
func (service *Service) MeetingQueryByTitle(userName string, title string) []model.Meeting {
	return service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetTitle() == title && (meeting.GetSponsor() == userName ||
			meeting.IsParticipators(userName))
	})
}

// 查询会议---通过usernsme(作为会议发起者和参与者)和会议起始时间查找
func (service *Service) MeetingQueryByUserAndTime(
	userName string, startDateString string, endDateString string) []model.Meeting {

	// 检查字符串是否合法
	if !IsValidStartAndEndDateTime(&startDateString, &endDateString) {
		return []model.Meeting{}
	}
	// 获取用户发起/参与且时间冲突的会议
	return service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return (meeting.GetSponsor() == userName || meeting.IsParticipators(userName)) &&
			((meeting.GetStartDate() <= startDateString && meeting.GetEndDate() >= startDateString) ||
				(meeting.GetStartDate() <= endDateString && meeting.GetEndDate() >= endDateString) ||
				(meeting.GetStartDate() >= startDateString && meeting.GetEndDate() <= endDateString))
	})
}

// 列出该用户发起或参与的所有会议
func (service *Service) ListAllMeetings(userName string) []model.Meeting {
	return service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetSponsor() == userName && meeting.IsParticipators(userName)
	})
}

// 列出该用户发起的所有会议
func (service *Service) ListAllSponsorMeetings(userName string) []model.Meeting {
	return service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetSponsor() == userName
	})
}

// 列出该用户参加的所有会议
func (service *Service) ListAllParticipateMeetings(userName string) []model.Meeting {
	return service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.IsParticipators(userName)
	})
}

// 取消会议--发起者根据title删除会议
func (service *Service) DeleteMeeting(sponsor string, title string) bool {
	return service.AgendaStorage.DeleteMeetings(func(meeting model.Meeting) bool {
		return meeting.GetTitle() == title && meeting.GetSponsor() == sponsor
	})
}

// 退出会议 -- 参与者根据titile退出自己参加的会议安排
// 若退出后人数为0，删除会议
func (service *Service) QuitMeeting(participator string, title string) bool {
	meetings := service.AgendaStorage.QueryMeetings(func(meeting model.Meeting) bool {
		return meeting.GetTitle() == title && meeting.IsParticipators(participator)
	})
	for _, tMeeting := range meetings {
		tMeeting.DeleteParticipator(participator)  // 删除一个会议参与者
		if tMeeting.GetParticipatorsNumber() > 0 { // 会议还有参与者，更新会议数据
			service.AgendaStorage.UpdateMeeting(func(meeting model.Meeting) bool {
				return meeting.GetTitle() == tMeeting.GetTitle()
			}, tMeeting)
		} else { // 会议参与人数为0，删除该会议
			service.AgendaStorage.DeleteMeetings(func(meeting model.Meeting) bool {
				return meeting.GetTitle() == tMeeting.GetTitle()
			})
		}
	}
	return len(meetings) > 0
}

// 清空会议--删除用户自己发起的所有会议
func (service *Service) DeleteAllMeetings(sponsor string) bool {
	return service.AgendaStorage.DeleteMeetings(func(meeting model.Meeting) bool {
		return meeting.GetSponsor() == sponsor
	})
}
