package gtype

import "sync"

type BaseTypeNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint16 | ~uint8 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type BaseTypeString interface {
	~string | ~[]byte | ~[]rune
}

type StringKeyValue struct {
	//主键
	Key int
	//值
	Value int
	//权重
	Weight int
}

type GSyncMap struct {
	//数据
	Smap sync.Map
	//长度
	Length int
}
