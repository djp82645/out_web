package utils

import (
	"reflect"
	"strconv"
	"unsafe"
)

func String2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pbytes.Data = stringHeader.Data
	pbytes.Len = stringHeader.Len
	pbytes.Cap = stringHeader.Len

	return b
}

func IsNil(i interface{}) (is_nil bool) {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// interface转int
func Interface2Int(inter interface{}) int {
	switch inter := inter.(type) {
	case int:
		return inter
	case int64:
		return int(inter)
	case float64:
		return int(inter)
	case string:
		inter_value, err := strconv.Atoi(inter)
		if err != nil {
			return 0
		}
		return inter_value
	default:
		return 0
	}
}
