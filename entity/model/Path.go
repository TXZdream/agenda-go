package model

import "os/user"

var DataDirPath string = "/.agenda"
var UserDataPath string = "/users.json"
var MeetingDataPath string = "/meetings.json"
var CurUserPath string = "/curUser.txt"
var LogPath string = "/log/Agenda.log"

func init() {
	user, err := user.Current()
	if err == nil {
		DataDirPath = user.HomeDir + DataDirPath
		UserDataPath = DataDirPath + UserDataPath
		MeetingDataPath = DataDirPath + MeetingDataPath
		CurUserPath = DataDirPath + CurUserPath
		LogPath = DataDirPath + LogPath
	} else {
		panic(err)
	}
}
