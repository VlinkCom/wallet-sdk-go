package ex_wallet

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func SignToString(obj interface{}) string {
	var str strings.Builder

	// 判断是否是map类型
	if m, ok := obj.(map[string]interface{}); ok {
		// 如果是map，则按字典顺序处理
		var keys []string
		for key := range m {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			if val, ok := m[key]; ok && val != nil {
				str.WriteString(fmt.Sprintf("%s=%v&", key, val))
			}
		}
	} else {
		// 如果是结构体，则通过反射处理
		v := reflect.ValueOf(obj)
		if v.Kind() == reflect.Ptr {
			v = v.Elem() // 解引用指针
		}

		if v.Kind() == reflect.Struct {
			var keys []string
			// 遍历结构体字段
			for i := 0; i < v.NumField(); i++ {
				field := v.Type().Field(i)
				// 通过字段名获取值
				val := v.Field(i).Interface()
				keys = append(keys, field.Name)
				// 检查值是否为 nil
				if val != nil {
					str.WriteString(fmt.Sprintf("%s=%v&", field.Name, val))
				}
			}
			sort.Strings(keys)
		}
	}

	// 删除末尾的 '&'
	result := str.String()
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}
