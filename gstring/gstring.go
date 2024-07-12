package gstring

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/aimey-a/go-tools/gcollect"
	"github.com/aimey-a/go-tools/gmath"
	"github.com/aimey-a/go-tools/gtype"
)

// 字符串转int
func ToNumber[T gtype.BaseTypeNumber](str string) T {
	itr, _ := strconv.Atoi(str)
	return T(itr)
}

// int转字符串
func ToString[T gtype.BaseTypeNumber](itr T) string {
	return fmt.Sprint(itr)
}

// 保留指定位数的小数（参数为float32或float64）（默认保留两位小数）
//	fixed: 指定小数点位数
func ToFixed(float any, fixed ...int) string {
	tfixed := 2
	if len(fixed) == 1 {
		tfixed = fixed[0]
	}
	if tfixed < 0 {
		tfixed = 0
	}
	if float != nil {
		switch float.(type) {
		case float32, float64:
			return fmt.Sprintf("%."+ToString(tfixed)+"f", float)
		}
	}
	return ""
}

// 字符串分割
//	sep: 分割符
func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

// 找到指定字符的索引
func IndexOf(str string, of string) int {
	return strings.Index(str, of)
}

// 找到指定字符的索引（后）
func LastIndexOf(str string, of string) int {
	return strings.LastIndex(str, of)
}

// 是否以指定字符起始
func StartWith(str string, of string) bool {
	return strings.HasPrefix(str, of)
}

// 是否以指定字符结束
func EndWith(str string, of string) bool {
	return strings.HasSuffix(str, of)
}

// 是否包含指定字符
func Contains(str string, of string) bool {
	return strings.Contains(str, of)
}

// 是否为空
func IsEmpty(str string) bool {
	return str == ""
}

// 截取
func Sub(str string, from int, to int) string {
	rs := []rune(str)
	length := len(rs)
	if from < 0 || to < 0 || from > to {
		return ""
	}
	if to > length {
		to = length
	}
	return string(rs[from:to])
}

// 替换所有指定字符
func Replace(str string, from string, to string) string {
	return strings.ReplaceAll(str, from, to)
}

// 剔除多余的空格
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// 字符串转字节数组
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 字节数组转字符串
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 字符串格式化
func Format(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

//切割数据  {1,1,1},{1,1}  类型
func ExactCutting(str string, decollatorStr ...string) (attList []*gtype.StringKeyValue) {
	strRune := []rune(str)
	openingBrace := []rune("{")[0]
	ClosingBrace := []rune("}")[0]
	//切割符
	var decollator rune
	if len(decollatorStr) > 0 {
		decollator = []rune(decollatorStr[0])[0]
	} else {
		decollator = []rune(",")[0]
	}
	for {
		opening := gcollect.IndexForAny(strRune, openingBrace)
		if opening != -1 {
			att := &gtype.StringKeyValue{}
			att.Weight = -1
			closing := gcollect.IndexForAny(strRune, ClosingBrace)
			if opening+1 > closing {
				fmt.Println("失败", opening, closing)
				return
			}
			s := strRune[opening+1 : closing]
			index := 0
			for {
				decollatorIndex := gcollect.IndexForAny(s, decollator)
				if decollatorIndex != -1 {
					val := string(s[:decollatorIndex])
					switch index {
					case 0:
						att.Key = ToNumber[int](val)
					case 1:
						att.Value = ToNumber[int](val)
					case 2:
						att.Weight = ToNumber[int](val)
					}
					index++
				} else {
					val := string(s[decollatorIndex+1:])
					switch index {
					case 0:
						att.Key = ToNumber[int](val)
					case 1:
						att.Value = ToNumber[int](val)
					case 2:
						att.Weight = ToNumber[int](val)
					}
					break
				}
				s = s[decollatorIndex+1:]
			}
			attList = append(attList, att)
			if len(strRune) > closing+2 {
				strRune = strRune[closing+2:]
			} else {
				break
			}
		} else {
			break
		}
	}
	return
}

// 切割数据  {1,1,1},{1,1} 类型 已结算概率
func ExactCuttingProbability(str string, decollatorStr ...string) (attList []*gtype.StringKeyValue) {
	strRune := []rune(str)
	openingBrace := []rune("{")[0]
	ClosingBrace := []rune("}")[0]
	//切割符
	var decollator rune
	if len(decollatorStr) > 0 {
		decollator = []rune(decollatorStr[0])[0]
	} else {
		decollator = []rune(",")[0]
	}
	var probabilityValue = 0
	var attProbabilityList []*gtype.StringKeyValue
	for {
		opening := gcollect.IndexForAny(strRune, openingBrace)
		if opening != -1 {
			att := &gtype.StringKeyValue{}
			att.Weight = -1
			closing := gcollect.IndexForAny(strRune, ClosingBrace)
			if opening+1 > closing {
				fmt.Println("失败", opening, closing)
				return
			}
			s := strRune[opening+1 : closing]
			index := 0
			for {
				decollatorIndex := gcollect.IndexForAny(s, decollator)
				if decollatorIndex != -1 {
					val := string(s[:decollatorIndex])
					switch index {
					case 0:
						att.Key = ToNumber[int](val)
					case 1:
						att.Value = ToNumber[int](val)
					case 2:
						att.Weight = ToNumber[int](val)
					}
					index++
				} else {
					val := string(s[decollatorIndex+1:])
					switch index {
					case 0:
						att.Key = ToNumber[int](val)
					case 1:
						att.Value = ToNumber[int](val)
					case 2:
						att.Weight = ToNumber[int](val)
					}
					break
				}
				s = s[decollatorIndex+1:]
			}
			if att.Weight > -1 {
				attProbabilityList = append(attProbabilityList, att)
				probabilityValue += att.Weight
			} else {
				attList = append(attList, att)
			}

			if len(strRune) > closing+2 {
				strRune = strRune[closing+2:]
			} else {
				break
			}
		} else {
			break
		}
	}
	rand.Seed(time.Now().UnixNano())
	rate := gmath.RandInt(0, probabilityValue+1)
	drop := Probability(attProbabilityList, probabilityValue, rate)
	if drop != nil {
		attList = append(attList, drop)
	}
	return
}

// 概率处理  data 值  probability 概率  rate 随机值
func Probability(data []*gtype.StringKeyValue, probability, rate int) *gtype.StringKeyValue {
	rateNum := 0
	if probability == 0 {
		if len(data) > 0 {
			return data[0]
		} else {
			return nil
		}
	}
	if len(data) == 1 {
		return data[0]
	}
	for _, v := range data {
		if v.Weight+rateNum > rate {
			if v.Key != 0 && v.Value != 0 {
				info := &gtype.StringKeyValue{}
				info.Key = v.Key
				info.Value = v.Value
				info.Weight = v.Weight
				return info
			}
		} else {
			rateNum += v.Weight
		}
	}
	return nil
}

//根据分隔符和拼接符 将字符串首字母大小写
//str 数据  concatenator 拼接符  LowerOrUpper 是否大小写 true 为小写 false 文瑞大写   decollatorList 分隔符
func CapitalizeSndSplice(str, concatenator string, LowerOrUpper bool, decollatorList ...string) string {
	var decollator string
	if len(decollatorList) > 0 {
		decollator = decollatorList[0]
	} else {
		decollator = " "
	}
	b := strings.Split(str, decollator)
	var c strings.Builder // 使用 strings.Builder 来提高性能
	for i, v := range b {
		if LowerOrUpper {
			c.WriteString(strings.ToLower(v))
		} else {
			c.WriteString(strings.ToUpper(v))
		}
		if i != len(b)-1 {
			c.WriteString(concatenator)
		}
	}
	return c.String()
}

// 去除超字符集
func RemovehyIllegalCharacter(name string) string {
	playerName := ""
	runes := []rune(name)
	rule := true
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if len([]byte(string(r))) < 4 {
			if r != 32 {
				playerName += string(r)
			}
		}
	}
	if rule && len(playerName) != len(name) {
		if playerName == "" {
			rule = false
		}
	}
	return playerName
}

//替换超字符集
func ReplaceIllegalCharacter(name string) string {
	playerName := ""
	runes := []rune(name)
	rule := true
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		ok := false
		if len([]byte(string(r))) < 4 {
			if r != 32 {
				ok = true
				playerName += string(r)
			}
		}
		if !ok {
			playerName += "?"
		}
	}
	if rule && len(playerName) != len(name) {
		if playerName == "" {
			rule = false
		}
	}
	return playerName
}
