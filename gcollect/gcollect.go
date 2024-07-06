package gcollect

import "github.com/aimey-a/go-tools/gtype"

// 从数组/切片中索引元素（[]any）
//	arr: 数组/切片
//	ele: 元素
func IndexForAny[T gtype.BaseTypeNumber](arr []T, ele T) int {
	if arr != nil {
		for k, v := range arr {
			if v == ele {
				return k
			}
		}
	}
	return -1
}

// 数组/切片是否存在元素（[]any）
//	arr: 数组/切片
//	ele: 元素
func ContainsForAny[T gtype.BaseTypeNumber](arr []T, ele T) bool {
	if arr != nil {
		return IndexForAny(arr, ele) >= 0
	}
	return false
}

// 从数组/切片中移除元素（[]any）
//	arr: 数组/切片
//	ele: 元素
func RemoveForAny[T gtype.BaseTypeNumber](arr []T, ele T) []T {
	if arr != nil {
		for {
			idx := IndexForAny(arr, ele)
			if idx >= 0 {
				arr = append(arr[:idx], arr[idx+1:]...)
			} else {
				break
			}
		}
	}
	return arr
}

// 从数组/切片中移除元素（[]any）
//	arr: 数组/切片
//	index: 索引
func DeleteForAny(arr []any, index int) []any {
	if arr != nil {
		if index < len(arr) {
			arr = append(arr[:index], arr[index+1:]...)
		}
	}
	return arr
}

// 在数组/切片中新增元素（[]any）
//	arr: 数组/切片
//	ele: 元素
func AppendForAny(arr []any, ele any) []any {
	return append(arr, ele)
}
