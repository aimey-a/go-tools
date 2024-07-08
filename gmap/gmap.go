package gmap

import (
	"sync"
	"sync/atomic"
)

type GSyncMap struct {
	//数据
	smap sync.Map
	//长度
	length int64
}

// 深拷贝函数
func DeepCopyMap(src map[any]any) map[any]any {
	dst := make(map[any]any)
	for key, value := range src {
		switch v := value.(type) {
		case map[any]any:
			dst[key] = DeepCopyMap(v) // 递归拷贝嵌套的map
		default:
			dst[key] = v
		}
	}
	return dst
}

func (val *GSyncMap) Load(key any) (actual any, ok bool) {
	return val.smap.Load(key)
}
func (val *GSyncMap) Store(key any, value any) {
	val.LoadOrStore(key, value)
}
func (val *GSyncMap) Delete(key any) {
	val.LoadAndDelete(key)
}

func (val *GSyncMap) LoadAndDelete(key any) (actual any, loaded bool) {
	actual, loaded = val.smap.LoadAndDelete(key)
	if loaded {
		atomic.AddInt64(&val.length, -1)
	}
	return
}

func (val *GSyncMap) LoadOrStore(key any, value any) (actual any, loaded bool) {
	actual, loaded = val.smap.LoadOrStore(key, value)
	if !loaded {
		atomic.AddInt64(&val.length, 1)
	}
	return
}

func (val *GSyncMap) Range(f func(key any, value any) bool) {
	val.smap.Range(f)
}

func (val *GSyncMap) Len() (idx int) {
	return int(val.length)
}
