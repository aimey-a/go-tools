package gutil

import "github.com/aimey-a/go-tools/gtype"

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
