package gtime

import "time"

type baseType interface {
	int | int8 | int32 | int64 | uint8 | uint32 | uint64 | float32 | float64
}

//获取时间戳
func GetTimestamp[T baseType]() T {
	ltime := time.Now().Unix()
	return T(ltime)
}
