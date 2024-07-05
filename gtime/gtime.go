package gtime

import (
	"time"

	"github.com/aimey-a/go-tools/gtype"
)

const (
	STRING_TIME_FORMAT_FULL = "2006-01-02 15:04:05 +0800 CST" // 时间格式化模板（标准）
	STRING_TIME_FORMAT_LITE = "2006-01-02 15:04:05"           // 时间格式化模板（简易）
	STRING_TIME_FORMAT_FILE = "2006-01-02_15_04_05"           // 时间格式化模板（文件）
)

const (
	SEC_1   = 1       // 1秒
	SEC_5   = 5       // 5秒
	SEC_10  = 10      // 10秒
	SEC_15  = 15      // 15秒
	SEC_20  = 20      // 20秒
	SEC_25  = 25      // 25秒
	SEC_30  = 30      // 30秒
	SEC_35  = 35      // 35秒
	SEC_40  = 40      // 40秒
	SEC_45  = 45      // 45秒
	SEC_50  = 50      // 50秒
	SEC_55  = 55      // 55秒
	MIN_1   = 60      // 1分钟
	MIN_2   = 120     // 2分钟
	MIN_3   = 180     // 3分钟
	MIN_4   = 240     // 4分钟
	MIN_5   = 300     // 5分钟
	MIN_6   = 360     // 6分钟
	MIN_7   = 420     // 7分钟
	MIN_8   = 480     // 8分钟
	MIN_9   = 540     // 9分钟
	MIN_10  = 600     // 10分钟
	MIN_12  = 720     // 12分钟
	MIN_15  = 900     // 15分钟
	MIN_20  = 1200    // 20分钟
	MIN_25  = 1500    // 25分钟
	MIN_30  = 1800    // 30分钟
	MIN_35  = 2100    // 35分钟
	MIN_40  = 2400    // 40分钟
	MIN_45  = 2700    // 45分钟
	MIN_50  = 3000    // 50分钟
	MIN_55  = 3300    // 55分钟
	HOUR_1  = 3600    // 1小时
	HOUR_2  = 7200    // 2小时
	HOUR_3  = 10800   // 3小时
	HOUR_4  = 14400   // 4小时
	HOUR_5  = 18000   // 5小时
	HOUR_6  = 21600   // 6小时
	HOUR_7  = 25200   // 7小时
	HOUR_8  = 28800   // 8小时
	HOUR_9  = 32400   // 9小时
	HOUR_10 = 36000   // 10小时
	HOUR_11 = 39600   // 11小时
	HOUR_12 = 43200   // 12小时
	HOUR_13 = 46800   // 13小时
	HOUR_14 = 50400   // 14小时
	HOUR_15 = 54000   // 15小时
	HOUR_16 = 57600   // 16小时
	HOUR_17 = 61200   // 17小时
	HOUR_18 = 64800   // 18小时
	HOUR_19 = 68400   // 19小时
	HOUR_20 = 72000   // 20小时
	HOUR_21 = 75600   // 21小时
	HOUR_22 = 79200   // 22小时
	HOUR_23 = 82800   // 23小时
	DAY_1   = 86400   // 1天
	DAY_2   = 172800  // 2天
	DAY_3   = 259200  // 3天
	DAY_4   = 345600  // 4天
	DAY_5   = 432000  // 5天
	DAY_6   = 518400  // 6天
	DAY_7   = 604800  // 7天
	DAY_8   = 691200  // 8天
	DAY_9   = 777600  // 9天
	DAY_10  = 864000  // 10天
	DAY_15  = 1296000 // 15天
	DAY_20  = 1728000 // 20天
	DAY_30  = 2592000 // 30天
)

//获取时间戳(秒)
func GetTimestamp[T gtype.BaseTypeNumber]() T {
	ltime := time.Now().Unix()
	return T(ltime)
}

//获取时间戳(毫秒)
func GetMillisecond[T gtype.BaseTypeNumber]() T {
	ltime := time.Now().UnixMilli()
	return T(ltime)
}

//获取时间戳(微秒)
func GetMicrosecond[T gtype.BaseTypeNumber]() T {
	ltime := time.Now().UnixMicro()
	return T(ltime)
}

//获取时间戳(纳秒)
func GetNanosecond[T gtype.BaseTypeNumber]() T {
	ltime := time.Now().UnixNano()
	return T(ltime)
}

//获取时间
func NowTime() time.Time {
	return time.Now()
}

//获取时间戳转时间
//timestamp 时间戳
func ToTime[T gtype.BaseTypeNumber](timestamp T) time.Time {
	return time.Unix(int64(timestamp), 0)
}

//时间转时间戳
//times 时间
func TimeToUnix[T gtype.BaseTypeNumber](times time.Time) T {
	return T(times.Unix())
}

//获取距离零点的时间（秒）
//timestamp 时间戳
func TimeToZero[T gtype.BaseTypeNumber](timestamp ...T) T {
	var tim T
	if len(timestamp) > 0 {
		tim = timestamp[0]
	} else {
		tim = GetTimestamp[T]()
	}
	targetTime := ToTime(tim)

	tomorrow := targetTime.AddDate(0, 0, 1)
	tomorrowMidnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, targetTime.Location())
	duration := tomorrowMidnight.Sub(targetTime)
	return T(duration.Seconds())
}

//获取零点的时间戳
//timestamp 时间戳
func ZeroTime[T gtype.BaseTypeNumber](timestamp ...T) T {
	var tim T
	if len(timestamp) > 0 {
		tim = timestamp[0]
	} else {
		tim = GetTimestamp[T]()
	}
	targetTime := ToTime(tim)
	return TimeToUnix[T](time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), 0, 0, 0, 0, targetTime.Location()))
}

// 时间格式化
//	timestamp: 时间戳
//	format: 格式化模板
func Format[T gtype.BaseTypeNumber](timestamp T, format string) string {
	return ToTime(timestamp).Format(format)
}
