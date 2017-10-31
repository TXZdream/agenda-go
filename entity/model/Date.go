package model

import (
	"time"
	"strings"
	"strconv"
	"fmt"
)

type Date struct {
	DateTime time.Time
}

// 把用户输入的日期字符串解析出年月日时分
// 判断字符串是否符合要求并把字符串分解并转为int数组
// 标准字符串：2012-2-2/11:23
// Usage: var s [5]int
// result = agenda.StringDateTimeToIntArray("2012-2-2/11:23", &s)
// if result == true {.....}
func StringDateTimeToIntArray(dateTimeString string, dateTimeArray *[5]int) bool {
	// 下划线分隔年月日和时分
	splitBySlash := strings.Split(dateTimeString, "/")
	if len(splitBySlash) < 2 {
		return false
	}
	// 中划线分隔年月日
	splitByLineThrough := strings.Split(splitBySlash[0], "-")
	// 冒号分隔时分
	splitByColon := strings.Split(splitBySlash[1], ":")
	if len(splitByLineThrough) != 3 || len(splitByColon) != 2 {
		
		fmt.Println(len(splitByLineThrough))
		fmt.Println(len(splitByColon))
		return false
	}

	var err error
	dateTimeArray[0], err = strconv.Atoi(splitByLineThrough[0])
	if (err != nil) {
		return false
	}
	dateTimeArray[1], err = strconv.Atoi(splitByLineThrough[1])
	if (err != nil) {
		return false
	}
	dateTimeArray[2], err = strconv.Atoi(splitByLineThrough[2])
	if (err != nil) {
		return false
	}
	dateTimeArray[3], err = strconv.Atoi(splitByColon[0])
	if (err != nil) {
		return false
	}
	dateTimeArray[4], err = strconv.Atoi(splitByColon[1])
	if (err != nil) {
		return false
	}

	return true
}

// 根据年月日时分设置time.Date
func SetDateByYMDHM(dateTimeArray [5]int) time.Time {
	loc, err := time.LoadLocation("")
	if err != nil {
		panic(err)
	}
	return time.Date(dateTimeArray[0], (time.Month)(dateTimeArray[1]), dateTimeArray[2], 
			dateTimeArray[3], dateTimeArray[4], 0, 0, loc)
}
// ------------ 判断日期是否合法 ------------
// 判断是否是闰年
func IsLeapYear(year int) bool {
	return year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
}

func IsValidYear(year int) bool {
	return year >= 1000 && year <= 9999
}

func IsValidMonth(month int) bool {
	return month >= (int)(time.January) && month <= (int)(time.December)	
}

// 根据年月判断日是否合法，年月必须先保证合法
func IsValidDayByYM(year, month, day int) bool{
	monthDays := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if IsLeapYear(year) {	// 闰年二月
		monthDays[2] = 29
	}
	return day > 0 && day <= monthDays[month]
}

func IsValidHour(hour int) bool {
	return hour >= 0 && hour < 24
}

func IsValidMinute(minute int) bool {
	return minute >= 0 && minute < 60
}
// 判断DateTime是否合法
func IsValidDateTime(dateTimeArray [5]int) bool {
	return IsValidYear(dateTimeArray[0]) && IsValidMonth(dateTimeArray[1]) && 
		   IsValidDayByYM(dateTimeArray[0], dateTimeArray[1], dateTimeArray[2]) && 
		   IsValidHour(dateTimeArray[3]) && IsValidMinute(dateTimeArray[4])
}
// ---------------------------------------

// ------------- 日期比较 -----------------
func (self Date) Equal(other Date) bool {
	return self.DateTime.Equal(other.DateTime)
}

func (self Date) Before(other Date) bool {
	return self.DateTime.Before(other.DateTime)
}

func (self Date) After(other Date) bool {
	return self.DateTime.After(other.DateTime)
}
// --------------------------------------

// fmt.Println(time.Time)	{2016-07-12 22:10:00 +0000 UTC}
// 返回日期的string形式 : 2016-07-12/22:10
func (date Date) ToString() string {
	return strings.Join(strings.Split(date.DateTime.String()[0:16], " "), "/")
}

