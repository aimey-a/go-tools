package gmath

import (
	"math/rand"

	"github.com/aimey-a/go-tools/gtype"
)

// 最大值
func MaxValue[T gtype.BaseTypeNumber](a T, b T) T {
	if a >= b {
		return a
	}
	return b
}

// 最小值
func MinValue[T gtype.BaseTypeNumber](a T, b T) T {
	if a >= b {
		return b
	}
	return a
}

// 随机数
//	min: 左区间
//	max: 右区间
func RandInt(min int, max int) int {
	if min >= max {
		return max
	}
	return rand.Intn(max-min) + min
}
