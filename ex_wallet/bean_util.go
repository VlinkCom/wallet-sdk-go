package ex_wallet

import (
	"reflect"
)

// BeanToMap 将结构体转换为 map
func BeanToMap(bean interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(bean)
	// 确保是结构体类型
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // 如果是指针，获取指向的值
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i)

		// 忽略 "privateKey" 和 "publicKey"
		if field.Name == "PrivateKey" || field.Name == "PublicKey" {
			continue
		}

		// 如果字段值为空，则跳过
		if value.IsZero() {
			continue
		}

		// 使用结构体字段名作为 map 的键
		// 结构体的字段名如果是大写，默认是可导出的，因此可以作为 map 的键
		key := field.Tag.Get("json")
		// 将键转为小写字母开头的形式以符合 Java 版本的要求
		result[key] = value.Interface()
	}

	return result
}
