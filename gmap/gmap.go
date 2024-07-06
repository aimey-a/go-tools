package gmap

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
