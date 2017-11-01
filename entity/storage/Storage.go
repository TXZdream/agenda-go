package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/txzdream/agenda-go/entity/model"
)

// Usage : agenda.Storage{Users: []agenda.User{}, Meetings: []agenda.Meeting{}}
type Storage struct {
	Users    []model.User
	Meetings []model.Meeting
}

var instance *Storage
var once sync.Once

// 单例模式
func GetStorageInstance() *Storage {
	once.Do(func() {
		instance = &Storage{}
	})
	return instance
}

type StorageError string

const (
	// 文件夹--存在/创建
	SucceedCreateDataDir StorageError = "Succeed In Creating Data Dir"
	ExistFileNamedData   StorageError = "Fail To Create Data Dir Because Of Exisiting A File Named \"Data\""
	FailCreateDataDir    StorageError = "Fail To Create Data Dir"
	// 文件--创建
	SucceedCreateDataFile StorageError = "Succeed In Creating Data File"
	FailCreateDataFile    StorageError = "Fail To Create Data File"
	// 文件--读取
	SucceedReadDateFile StorageError = "Succeed In Reading Data File"
	FailReadDataFile    StorageError = "Fail To Read Data File"
	// 文件--获取json数据
	FailGetJsonData StorageError = "Fail To Read Json Data"
	// 文件--写入
	SucceedWriteDataFile StorageError = "Succeed In Writing Data File"
	FailWriteDataFile    StorageError = "Fail To Write Data File"
)

// 判断文件夹是否存在否则创建一个，返回状态+错误信息
func IsExistDataDirOrCreate() (bool, StorageError) {
	file, err := os.Stat(model.DataDirPath)
	if err == nil {
		if file.Mode().IsDir() {
			return true, SucceedCreateDataDir
		}
		return false, ExistFileNamedData
	}
	err = os.Mkdir(model.DataDirPath, os.ModePerm)
	if err != nil {
		return false, FailCreateDataDir
	}
	return true, SucceedCreateDataDir
}

// 判断File是否存在否则创建一个，返回状态+错误信息
func IsExistFileOrCreate(fileName string) (bool, StorageError) {
	file, err := os.Stat(fileName)
	// 文件不存在，则创建
	if os.IsNotExist(err) {
		_, err = os.Create(fileName)
		if err != nil { // 创建失败
			return false, FailCreateDataFile
		}
		return true, SucceedReadDateFile
	}
	// 判断是否是文件
	if file.Mode().IsRegular() {
		return true, SucceedReadDateFile
	}
	return false, FailCreateDataFile
}

// -------------- 判断文件是否存在 ---------------
func IsExistCurrentUserFileOrCreate() (bool, StorageError) {
	return IsExistFileOrCreate(model.CurUserPath)
}

func IsExistUserFileOrCreate() (bool, StorageError) {
	return IsExistFileOrCreate(model.UserDataPath)
}

func IsExistMeetingFileOrCreate() (bool, StorageError) {
	return IsExistFileOrCreate(model.MeetingDataPath)
}

// ---------------------------------------------

// 判断文件夹和文件是否存在
func IsExistDataFilesOrCreate() (bool, StorageError) {
	// 是否存在data文件夹
	result, storageError := IsExistDataDirOrCreate()
	if !result {
		return result, storageError
	}
	// 是否存在curUser文件
	result, storageError = IsExistCurrentUserFileOrCreate()
	if !result {
		return result, storageError
	}
	// 是否存在user文件
	result, storageError = IsExistUserFileOrCreate()
	if !result {
		return result, storageError
	}
	// 是否存在meetng文件
	result, storageError = IsExistMeetingFileOrCreate()
	if !result {
		return result, storageError
	}
	return true, SucceedCreateDataFile
}

// 根据文件名读取文件
func ReadFromFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

// 读取curUser文件，返回成功与否+登陆用户名/错误信息
func (storage *Storage) ReadFromCurrentUserFile() (bool, string) {
	currentUserName, err := ReadFromFile(model.CurUserPath)
	if err != nil {
		return false, string(FailReadDataFile)
	}
	return true, string(currentUserName)
}

// 读取User文件,读取所有用户列表，返回成功与否+错误信息
func (storage *Storage) ReadFromUserFile() (bool, StorageError) {
	usersJson, err := ReadFromFile(model.UserDataPath)
	if err != nil {
		return false, FailReadDataFile
	}
	if len(usersJson) != 0 && json.Unmarshal([]byte(usersJson), &storage.Users) != nil {
		return false, FailGetJsonData
	}
	return true, SucceedReadDateFile
}

// 读取Meeting文件,读取所有会议列表，返回成功与否+错误信息
func (storage *Storage) ReadFromMeetingFile() (bool, StorageError) {
	meetingsJson, err := ReadFromFile(model.MeetingDataPath)
	if err != nil {
		return false, FailReadDataFile
	}
	if len(meetingsJson) != 0 && json.Unmarshal([]byte(meetingsJson), &storage.Meetings) != nil {
		return false, FailGetJsonData
	}
	return true, SucceedReadDateFile
}

// 从文件中读取用户和会议数据
func (storage *Storage) ReadFromDataFile() (bool, StorageError) {
	// 文件是否存在，否则创建
	result, storageError := IsExistDataFilesOrCreate()
	if !result {
		return result, storageError
	}
	// 读取用户列表
	result, storageError = storage.ReadFromUserFile()
	if !result {
		return result, storageError
	}
	// 读取会议列表
	result, storageError = storage.ReadFromMeetingFile()
	if !result {
		return result, storageError
	}
	return true, SucceedReadDateFile
}

// 根据文件名把数据写入文件
func WriteToFile(fileName string, content []byte) bool {
	if ioutil.WriteFile(fileName, content, 0777) != nil {
		return false
	}
	return true
}

// 写入当前用户信息
func (storage *Storage) WriteToCurrentUserFile(CurrentUserName string) bool {
	return WriteToFile(model.CurUserPath, []byte(CurrentUserName))
}

// 写入User.json
func (storage *Storage) WriteUserFile() bool {
	userJson, err := json.Marshal(storage.Users)
	if err != nil {
		return false
	}
	return WriteToFile(model.UserDataPath, userJson)
}

// 写入Meeting.json
func (storage *Storage) WriteMeetingFile() bool {
	meetingJson, err := json.Marshal(storage.Meetings)
	if err != nil {
		return false
	}
	return WriteToFile(model.MeetingDataPath, meetingJson)
}

// 退出登陆，清空当前用户，把当前用户名、用户列表数据和会议列表数据写入
func (storage *Storage) LogOutStorage(CurrentUserName string) (bool, StorageError) {
	instance = nil
	if !storage.WriteToCurrentUserFile(CurrentUserName) {
		return false, FailWriteDataFile
	}
	if !storage.WriteUserFile() {
		return false, FailWriteDataFile
	}
	if !storage.WriteMeetingFile() {
		return false, FailWriteDataFile
	}
	return true, SucceedWriteDataFile
}

// ----------- 对用户列表进行操作，需要把改动写入文件，并返回是否成功 ------------
// 创建用户
func (storage *Storage) CreateUser(user model.User) bool {
	storage.Users = append(storage.Users, user)
	return storage.WriteUserFile()
}

// 根据filter函数查找用户
func (storage *Storage) QueryUsers(filter func(user model.User) bool) []model.User {
	var users []model.User
	for _, tUser := range storage.Users {
		if filter(tUser) {
			users = append(users, tUser)
		}
	}
	return users
}

// 更新用户信息，返回是否是否更新成功
func (storage *Storage) UpdateUser(filter func(user model.User) bool, updatedUser model.User) bool {
	for index, tUser := range storage.Users {
		if filter(tUser) {
			storage.Users[index] = updatedUser
			return storage.WriteUserFile()
		}
	}
	return false
}

// 删除用户
func (storage *Storage) DeleteUser(filter func(user model.User) bool) bool {
	isDeleted := false // 是否进行过删除
	for index, tUser := range storage.Users {
		if filter(tUser) {
			storage.Users = append(storage.Users[:index], storage.Users[index+1:]...)
			isDeleted = true
			break
		}
	}
	return isDeleted && storage.WriteUserFile()
}

// ------------------------------------------------------------------

// ----------- 对会议列表进行操作，需要把改动写入文件，并返回是否成功 ------------
// 创建会议
func (storage *Storage) CreateMeeting(meeting model.Meeting) bool {
	storage.Meetings = append(storage.Meetings, meeting)
	return storage.WriteMeetingFile()
}

// 根据filter函数查找会议
func (storage *Storage) QueryMeetings(filter func(meeting model.Meeting) bool) []model.Meeting {
	var meetings []model.Meeting
	for _, tMeeting := range storage.Meetings {
		if filter(tMeeting) {
			meetings = append(meetings, tMeeting)
		}
	}

	return meetings
}

// 更新会议信息，返回是否是否更新成功
func (storage *Storage) UpdateMeeting(filter func(meeting model.Meeting) bool, updatedMeeting model.Meeting) bool {
	for index, tMeeting := range storage.Meetings {
		if filter(tMeeting) {
			storage.Meetings[index] = updatedMeeting
			return storage.WriteMeetingFile()
		}
	}
	return false
}

// 删除会议
func (storage *Storage) DeleteMeetings(filter func(meeting model.Meeting) bool) bool {
	isDeleted := false // 是否进行过删除
	for index, tMeeting := range storage.Meetings {
		if filter(tMeeting) {
			storage.Meetings = append(storage.Meetings[:index], storage.Meetings[index+1:]...)
			isDeleted = true
		}
	}
	return isDeleted && storage.WriteMeetingFile()
}

// ------------------------------------------------------------------
