package toolx

import (
	"fmt"
	"time"
)

const (
	YYYY           = "2006"
	YYYYMM         = "2006-01"
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDHHMM   = "2006-01-02 15:04"
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
	HHMMSS         = "15:04:05"
	HHMM           = "15:04"
)

// FormatTime 格式化时间
func FormatTime(time time.Time) string {
	return time.Format(YYYYMMDDHHMMSS)
}

// ParseTime 将时间字符串转为Time
func ParseTime(str string) (time.Time, error) {
	return time.Parse(YYYYMMDDHHMMSS, str)
}

// 获取相差时间秒
func GetSecDiffer(startTime, endTime time.Time) int64 {
	var sec int64
	if startTime.Before(endTime) {
		sec = endTime.Unix() - startTime.Unix()
	}
	return sec
}

// 判断是否时同一天
func IsSameDay(t1 int64, t2 int64) bool {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	return time1.YearDay() == time2.YearDay()
}

// 倒计时
func CountDown(t int) string {
	day := t / 86400
	hour := t % 86400 / 3600
	min := t % 3600 / 60
	sec := t % 60

	timeArr := []int{day, hour, min, sec}
	timeDesc := []string{"天", "小时", "分钟", "秒"}

	var res string
	ok := false
	for i := 0; i < len(timeArr); i++ {
		if ok || timeArr[i] != 0 {
			res += fmt.Sprintf("%d%s", timeArr[i], timeDesc[i])
		}
		if timeArr[i] != 0 {
			ok = true
		}
	}
	return res
}

// 获取多久之前时间
func TimeDiffer(setTime time.Time) string {
	var rtime string

	sec := GetSecDiffer(setTime, time.Now())
	switch {
	case sec < 5: // 5秒之前
		rtime = "1秒之前"
	case sec < 10: // 10秒之前
		rtime = "5秒之前"
	case sec < 30: // 30秒之前
		rtime = "10秒之前"
	case sec < 60: // 1分钟之前
		rtime = "30秒之前"
	case sec < 60*5: // 5分钟之前
		rtime = "1分钟之前"
	case sec < 60*10: // 10分钟之前
		rtime = "5分钟之前"
	case sec < 60*30: // 30分钟之前
		rtime = "10分钟之前"
	case sec < 60*60: // 1个小时之前
		rtime = "30分钟之前"
	case sec < 60*60*2: // 2个小时之前
		rtime = "1个小时之前"
	case sec < 60*60*3: // 3个小时之前
		rtime = "2个小时之前"
	case sec < 60*60*4: // 4个小时之前
		rtime = "3个小时之前"
	case sec < 60*60*5: // 5个小时之前
		rtime = "4个小时之前"
	case sec < 60*60*6: // 6个小时之前
		rtime = "5个小时之前"
	case sec < 60*60*7:
		rtime = "6个小时之前"
	case sec < 60*60*8:
		rtime = "7个小时之前"
	case sec < 60*60*9:
		rtime = "8个小时之前"
	case sec < 60*60*10:
		rtime = "9个小时之前"
	case sec < 60*60*11:
		rtime = "10个小时之前"
	case sec < 60*60*12:
		rtime = "11个小时之前"
	case sec < 60*60*13:
		rtime = "12个小时之前"
	case sec < 60*60*14:
		rtime = "13个小时之前"
	case sec < 60*60*15:
		rtime = "14个小时之前"
	case sec < 60*60*16:
		rtime = "15个小时之前"
	case sec < 60*60*17:
		rtime = "16个小时之前"
	case sec < 60*60*18:
		rtime = "17个小时之前"
	case sec < 60*60*19:
		rtime = "18个小时之前"
	case sec < 60*60*20:
		rtime = "19个小时之前"
	case sec < 60*60*21:
		rtime = "20个小时之前"
	case sec < 60*60*22:
		rtime = "21个小时之前"
	case sec < 60*60*23:
		rtime = "22个小时之前"
	case sec < 60*60*24:
		rtime = "23个小时之前"
	case sec < 60*60*24*2:
		rtime = "24个小时之前"
	case sec < 60*60*24*3:
		rtime = "2天之前"
	case sec < 60*60*24*4: // 4天之前
		rtime = "3天之前"
	case sec < 60*60*24*5: // 5天之前
		rtime = "4天之前"
	case sec < 60*60*24*6: // 6天之前
		rtime = "5天之前"
	case sec < 60*60*24*15: // 半个月之前
		rtime = "5天之前"
	case sec < 60*60*24*30: // 1个月之前
		rtime = "半个月之前"
	case sec < 60*60*24*30*2: // 2个月之前
		rtime = "1个月之前"
	case sec < 60*60*24*30*3: // 3个月之前
		rtime = "2个月之前"
	case sec < 60*60*24*30*6: // 半年之前
		rtime = "3个月之前"
	case sec < 60*60*24*30*12: // 1年之前
		rtime = "半年之前"
	case sec < 60*60*24*30*12*2: // 2年之前
		rtime = "1年之前"
	default: // 时间太久
		rtime = setTime.Format(YYYYMMDD)
	}
	return rtime
}
