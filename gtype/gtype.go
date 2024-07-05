package gtype

type BaseTypeNumber interface {
	~int | ~int8 | ~int32 | ~int64 | ~uint8 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type BaseTypeString interface {
	~string | ~[]byte | ~[]rune
}
