package gutil

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/aimey-a/go-tools/gtype"
)

// 分页  返回数据和总长度
func Pagination[T any](pageNumber, pageSize int, list []T) ([]T, int) {
	size := pageNumber * pageSize
	sizeOne := ((pageNumber - 1) * pageSize)
	total := len(list)
	if sizeOne > total {
		return []T{}, total
	} else if size > total {
		return list[sizeOne:], total
	} else {
		return list[sizeOne:size], total
	}
}

// 反向分页  返回数据和总长度
func ReversePagination[T any](pageNumber, pageSize int, list []T) ([]T, int) {
	total := len(list)
	// 计算反向分页的起始索引
	startIndex := total - pageSize*pageNumber
	if startIndex < 0 {
		startIndex = 0
	}
	// 确保结束索引不越界
	endIndex := startIndex + pageSize
	if endIndex > total {
		endIndex = total
	}
	// 反向切片以获取所需页面的数据
	reversedList := make([]T, 0, endIndex-startIndex)
	for i := endIndex - 1; i >= startIndex; i-- {
		reversedList = append(reversedList, list[i])
	}
	return reversedList, total
}

// 分页排行
func Paging[T gtype.BaseTypeNumber, R gtype.BaseTypeNumber](pageNumber, pageSize, i T) R {
	return R(((pageNumber - 1) * pageSize) + (i + 1))
}

//验证电子邮件
func ValidateEmail(str string) bool {
	// 定义电子邮件的正则表达式
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	re := regexp.MustCompile(pattern)

	// 使用MatchString方法检测email是否匹配pattern
	return re.MatchString(str)
}

// CheckIDCard 校验身份证号码是否有效
func CheckIDCard(idCard string) bool {
	// 正则表达式验证身份证号码格式
	matched, _ := regexp.MatchString(`^\d{17}(\d|x)$`, idCard)
	if !matched {
		return false
	}

	// 将身份证号码转换为大写，以防x是小写
	idCard = strings.ToUpper(idCard)

	// 定义权重数组和校验码对应的值
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkSum := 0

	// 计算前17位数字的加权和
	for i, r := range idCard[:17] {
		var digit int
		digit, _ = strconv.Atoi(string(r))
		checkSum += digit * weights[i]
	}

	// 获取校验码
	checkCode := "10X98765432"
	expectedCheckDigit := checkCode[checkSum%11]

	// 比较最后一位是否与计算出的校验码相符
	return expectedCheckDigit == idCard[17]
}

// CheckMobileNumber 验证手机号码格式是否正确
func CheckMobileNumber(phoneNumber string) bool {

	// 匹配中国大陆手机号码的正则表达式
	pattern := `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
	matched, _ := regexp.MatchString(pattern, phoneNumber)
	return matched
}
