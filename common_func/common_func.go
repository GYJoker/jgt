package common_func

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// StrIsEmpty 判断是否是空字符串
func StrIsEmpty(str string) bool {
	return len(str) == 0
}

// StrIsNotEmpty 判断是否不是空字符串
func StrIsNotEmpty(str string) bool {
	return len(str) > 0
}

// StrToUint64 将字符串转换为uint64
func StrToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// GetMapKeys 获取map的key
func GetMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func ArrayToString[T string | int | float64 | uint64 | int64](array []T) string {
	if len(array) == 0 {
		return ""
	}

	str := ""
	for _, s := range array {
		str += fmt.Sprintf("%v,", s)
	}

	return str[:len(str)-1]
}

func SplitStringToIntArray(str string) ([]int, error) {
	return splitAndConvert(str, func(s string) (int, error) {
		return strconv.Atoi(s)
	})
}

func SplitStringToUint64Array(str string) ([]uint64, error) {
	return splitAndConvert(str, func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 64)
	})
}

func SplitStringToFloat64Array(str string) ([]float64, error) {
	return splitAndConvert(str, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
}

func SplitStringToStringArray(str string) ([]string, error) {
	return splitAndConvert(str, func(s string) (string, error) {
		return s, nil
	})
}

func splitAndConvert[T any](str string, convert func(string) (T, error)) ([]T, error) {
	if len(str) == 0 {
		return make([]T, 0), nil
	}
	split := strings.Split(str, ",")
	array := make([]T, 0, len(split))

	for _, s := range split {
		v, err := convert(s)
		if err != nil {
			return nil, err
		}
		array = append(array, v)
	}

	return array, nil
}

// FormatStrMoney 将数字转换为金额格式
func FormatStrMoney(money int64) string {
	return fmt.Sprintf("%.2f", float64(money)/100.0)
}

// ArrContains 判断数组是否包含某个元素
func ArrContains[T string | int | float64 | int64](arr []T, val T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// ArrIsEqual 比较两个数组是否相等
func ArrIsEqual[T string | int | float64 | int64](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for index, value := range arr1 {
		if arr2[index] != value {
			return false
		}
	}

	return true
}

// GetNickNameByPhone 获取昵称
func GetNickNameByPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}

	return phone[:3] + "****" + phone[len(phone)-4:]
}

// ValueToJsonStr 将任意类型的值转换为json字符串
func ValueToJsonStr(value interface{}) string {
	if value == nil {
		return ""
	}

	marshal, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(marshal)
}

// MaxInt64 获取最大值
func MaxInt64(x, y int64) int64 {
	return ThreeWayOperator(x > y, x, y)
}

// CalRateToStr 计算百分比 输出字符串
func CalRateToStr(num, total int64) string {
	if total == 0 {
		return "0.00%"
	}
	return fmt.Sprintf("%.2f%%", float64(num)/float64(total)*100)
}

// StrVal 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func StrVal(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// ArrayInGroupsOf 将数组分割为多个数组
func ArrayInGroupsOf(arr []uint64, num int64) [][]uint64 {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]uint64{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]uint64, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

// ThreeWayOperator 三目运算
func ThreeWayOperator[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// UInt64RemoveRepeatedElement 去除uint64数组中重复的元素
func UInt64RemoveRepeatedElement(arr []uint64) (newArr []uint64) {
	newArr = make([]uint64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
