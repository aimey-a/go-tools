package gproto

import "github.com/aimey-a/go-tools/gtype"

func Int[T gtype.BaseTypeNumber](v T) *T {
	var val = v
	return &val
}

func String(v string) *string {
	var val = v
	return &val
}

func Bool(v bool) *bool {
	var val = v
	return &val
}
