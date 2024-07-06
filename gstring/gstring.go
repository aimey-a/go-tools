package gstring

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unsafe"

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
func ExactCutting(str string) []*gtype.StringKeyValue {
	attList := []*gtype.StringKeyValue{}
	for {
		a := IndexOf(str, "{")
		if a != -1 {
			att := &gtype.StringKeyValue{}
			b := IndexOf(str, "}")
			s := string([]rune(str)[a+1 : b])
			index := 0
			for {
				d := IndexOf(s, ",")
				if d != -1 {
					if index == 0 {
						e := string([]rune(s)[:d])
						att.Key = ToNumber[int](e)
						index++
					} else {
						e := string([]rune(s)[:d])
						att.Value = ToNumber[int](e)
						index++
					}
				} else if index == 1 {
					e := string([]rune(s)[d+1:])
					att.Value = ToNumber[int](e)
					att.Weight = -1
					break
				} else {
					if index == 0 {
					} else {
						att.Weight = ToNumber[int](s)
					}
					break
				}
				s = string([]rune(s)[d+1:])
			}
			attList = append(attList, att)
			if len(str) > b+2 {
				str = string([]rune(str)[b+2:])
			} else {
				break
			}
		} else {
			break
		}
	}
	return attList
}

// 切割数据  {1,1,1},{1,1} 类型 已结算概率
func ExactCuttingDropProbability(str string) []*gtype.StringKeyValue {
	attList := []*gtype.StringKeyValue{}
	attProbabilityList := []*gtype.StringKeyValue{}
	probabilityValue := 0
	for {
		a := strings.Index(str, "{")
		if a != -1 {
			att := &gtype.StringKeyValue{}
			b := strings.Index(str, "}")
			s := string([]rune(str)[a+1 : b])
			index := 0
			probability := true
			bk := true
			for {
				d := strings.Index(s, ",")
				if d != -1 {
					if index == 0 {
						e := string([]rune(s)[:d])
						att.Key = ToNumber[int](e)
						index++
					} else {
						e := string([]rune(s)[:d])
						att.Value = ToNumber[int](e)
						index++
					}
				} else if index == 1 {
					e := string([]rune(s)[d+1:])
					att.Value = ToNumber[int](e)
					att.Weight = -1
					probability = false
					break
				} else {
					if index == 0 {
						bk = false
					} else {
						att.Weight = ToNumber[int](s)
					}
					break
				}
				s = string([]rune(s)[d+1:])
			}
			if bk {
				if probability {
					attProbabilityList = append(attProbabilityList, att)
					probabilityValue += att.Weight
				} else {
					attList = append(attList, att)
				}
			}
			if len(str) > b+2 {
				str = string([]rune(str)[b+2:])
			} else {
				break
			}
		} else {
			break
		}
	}
	rand.Seed(time.Now().UnixNano())
	rate := gmath.RandInt(0, probabilityValue+1)
	drop := DropProbability(attProbabilityList, probabilityValue, rate)
	if drop != nil {
		attList = append(attList, drop)
	}
	return attList
}

// 概率处理  data 值  probability 概率  rate 随机值
func DropProbability(data []*gtype.StringKeyValue, probability, rate int) *gtype.StringKeyValue {
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
