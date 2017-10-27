package agenda
import (
	"fmt"
	"sync"
	"os"
	"io/ioutil"
)

// Usage : agenda.Storage{Users: []agenda.User{}, Meetings: []agenda.Meeting{}}
type Storage struct {
	Users []User
	Meetings []Meeting
}

var instance *Storage
var once sync.Once

// 单例模式
func GetStorageInstance() *Storage {
	once.Do(func() {
		instance = &Storage{Users: []User{}, Meetings: []Meeting{}}
	})
	return instance
}

type StorageError string
const (
	SucceedCreateDataDir = "Succeed In Creating Data Dir"
	ExistFileNamedData = "Fail To Create Data Dir Because Of Exisiting A File Named \"Data\""
	FailCreateDataDir = "Fail To Create Data Dir"
	SucceedCreateUserFile = "Succeed In Creating User File"
	FailCreateUserFile = "Fail To Create User File"
	SucceedCreateMeetingFile = "Succeed In Creating Meeting File"	
	FailCreateMeetingFile = "Fail To Create Meeting File"
	SucceedCreateDataFiles = "Succeed In Creating Data Files"
	FailReadUserFile = "Fail To Read From User File"
	FailReadMeetingFile = "Fail To Read From User File"
)


// 判断Data文件夹是否存在否则创建一个
// 返回状态+错误信息
func IsExistDataDirOrCreate() (bool, StorageError) {
	file, err := os.Stat(DataDirPath)
	if err == nil {
		if file.Mode().IsDir() {
			return true, SucceedCreateDataDir
		}
		return false, ExistFileNamedData
	}
	err = os.Mkdir(DataDirPath, os.ModeDir)
	if err != nil {
		return false, FailCreateDataDir
	}
	return true, SucceedCreateDataDir
}

// 判断User File是否存在否则创建一个
// 返回状态+错误信息
func IsExistUserFileOrCreate() (bool, StorageError) {
	file, err := os.Stat(UserDataPath)
	// 文件不存在，则创建
	if os.IsNotExist(err) { // 创建失败
		_, err = os.Create(UserDataPath)
		if err != nil {
			return false, FailCreateUserFile
		}
		return true, SucceedCreateUserFile
	}
	// 判断是否是文件
	if file.Mode().IsRegular() {
		return true, SucceedCreateUserFile
	}
	return false, FailCreateUserFile
}

// 判断Meeting File是否存在否则创建一个
// 返回状态+错误信息
func IsExistMeetingFileOrCreate() (bool, StorageError) {
	file, err := os.Stat(MeetingDataPath)
	// 文件不存在，则创建
	if os.IsNotExist(err) {
		_, err = os.Create(MeetingDataPath)
		if err != nil {	// 创建失败
			return false, FailCreateMeetingFile
		}
		return true, SucceedCreateMeetingFile
	}
	// 判断是否是文件
	if file.Mode().IsRegular() {
		return true, SucceedCreateMeetingFile
	}
	return false, FailCreateMeetingFile
}

// 判断users.json和meetings.json两个文件是否存在
func IsExistDataFilesOrCreate() (bool, StorageError) {
	// 是否存在data文件夹
	result, storageError := IsExistDataDirOrCreate()
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
	return true, SucceedCreateDataFiles
}

// 从文件中读取数据
func (storage *Storage)ReadFromFile() (bool, StorageError) {
	// 文件是否存在，否则创建
	result, storageError := IsExistDataFilesOrCreate()
	if !result {
		return result, storageError
	}
	// 读取文件数据
	content, err := ioutil.ReadFile(UserDataPath)
	if err != nil {
		return false, FailReadUserFile
	}
	fmt.Printf("File contents: %s", content)
	content, err = ioutil.ReadFile(MeetingDataPath)
	if err != nil {
		return false, FailReadMeetingFile
	}
	fmt.Printf("File contents: %s", content)
	fmt.Printf("ll")
	return true, SucceedCreateDataFiles
}

func WriteUserFile(b1 []byte) bool {
	if ioutil.WriteFile(UserDataPath, b1, 0644) != nil {
		return false
	}
	return true
}

func WriteMeetingFile(b1 []byte) bool {
	if ioutil.WriteFile(MeetingDataPath, b1, 0644) != nil {
		return false
	}
	return true
}


