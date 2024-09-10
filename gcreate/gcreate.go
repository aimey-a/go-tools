package gcreate

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// 生成身份证
func CreateIDCard() {
	// // 随机选择一个地址码
	// areaCode := areaCodes[rand.Intn(len(areaCodes))]

	// // 生成随机出生日期
	// birthDate := generateRandomDate()

	// // 生成随机顺序码
	// sequenceCode := generateSequenceCode()

	// // 拼接前17位
	// idWithoutChecksum := areaCode + birthDate + sequenceCode

	// // 计算校验码
	// checksum := calculateChecksum(idWithoutChecksum)

	// // 拼接成完整的身份证号码
	// return idWithoutChecksum + checksum
}

// 计算校验码
func calculateChecksum(idWithoutChecksum string) string {
	// 系数列表
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	// 校验码对应表
	checksumTable := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	sum := 0
	// 前17位每一位乘以系数相加
	for i := 0; i < 17; i++ {
		num, _ := strconv.Atoi(string(idWithoutChecksum[i]))
		sum += num * weights[i]
	}

	// 取模得到校验码
	checksumIndex := sum % 11
	return checksumTable[checksumIndex]
}

// 生成随机日期 (格式为 YYYYMMDD)
func generateRandomDate() string {
	// 随机生成1970到2000年之间的日期
	year := rand.Intn(31) + 1970 // 1970 到 2000 年
	month := rand.Intn(12) + 1   // 1 到 12 月
	day := rand.Intn(28) + 1     // 1 到 28 日，简化处理，不考虑闰年和不同月份天数

	return fmt.Sprintf("%04d%02d%02d", year, month, day)
}

// 生成随机顺序码 (最后一位表示性别，奇数为男，偶数为女)
func generateSequenceCode() string {
	sequence := rand.Intn(1000) // 000 到 999
	return fmt.Sprintf("%03d", sequence)
}

// 生成手机号
func CreateMobileNumber(val ...int) int {
	var prefix int
	if len(val) == 0 {
		// 常见的手机号码前缀，您可以根据需要扩展这个列表
		prefixes := []int{130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 150, 151, 152, 153, 155, 156, 157, 158, 159, 170, 171, 173, 175, 176, 177, 178, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189}
		// 选择一个随机的前缀
		prefix = prefixes[rand.Intn(len(prefixes))]
	} else {
		prefix = val[0]
	}
	// 生成随机的8位数字，组合成完整的11位手机号码
	number := rand.Intn(100000000)
	phoneNumber := fmt.Sprintf("%v%v", prefix, number)

	phone, _ := strconv.Atoi(phoneNumber)
	return phone
}

// 生成邮件
func CreateEmail(domain string) string {
	// 定义邮箱用户名的字符集
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 生成随机长度的邮箱用户名（长度在5到10之间）
	usernameLength := rand.Intn(6) + 5
	var username strings.Builder
	for i := 0; i < usernameLength; i++ {
		randomChar := charset[rand.Intn(len(charset))]
		username.WriteByte(randomChar)
	}
	// 组合成完整的邮箱地址
	email := fmt.Sprintf("%s@%s", username.String(), domain)
	return email
}
