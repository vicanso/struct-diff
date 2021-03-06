package dash

import (
	"errors"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

func getType(i interface{}) (valueType string) {
	switch i.(type) {
	case string:
		valueType = "string"
		break
	case bool:
		valueType = "bool"
		break
	case int:
		valueType = "int"
		break
	case int8:
		valueType = "int8"
		break
	case int16:
		valueType = "int16"
		break
	case int32:
		valueType = "int32"
		break
	case int64:
		valueType = "int64"
	case uint:
		valueType = "uint"
		break
	case uint8:
		valueType = "uint8"
		break
	case uint16:
		valueType = "uint16"
		break
	case uint32:
		valueType = "uint32"
		break
	case uint64:
		valueType = "uint64"
		break
	case uintptr:
		valueType = "uintptr"
		break
	case float32:
		valueType = "float32"
		break
	case float64:
		valueType = "float64"
		break
	case complex64:
		valueType = "complex64"
		break
	case complex128:
		valueType = "complex128"
		break
	}
	return
}

// ToMap convert struct to map
func ToMap(s interface{}) (data map[string]interface{}) {
	switch s.(type) {
	case map[string]interface{}:
		return s.(map[string]interface{})
	}
	return structs.Map(s)
}

// Difference get the diff between current and original，
// they should be the same struct type
func Difference(current, original interface{}) (diff map[string]interface{}) {
	map1 := ToMap(current)
	map2 := ToMap(original)
	diff = make(map[string]interface{})
	for name, v1 := range map1 {
		v2 := map2[name]
		if !reflect.DeepEqual(v1, v2) {
			diff[name] = v1
		}
	}
	return
}

// FindStringIndex get the index of string
func FindStringIndex(collection []string, value string) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesString check the value(string) is in collection
func IncludesString(collection []string, value string) bool {
	return FindStringIndex(collection, value) != -1
}

// FindIntIndex get the index of int
func FindIntIndex(collection []int, value int) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesInt check the value(int) is in collection
func IncludesInt(collection []int, value int) bool {
	return FindIntIndex(collection, value) != -1
}

// FindUint8Index get the index of uint8
func FindUint8Index(collection []uint8, value uint8) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesUint8 check the value(uint8) is in collection
func IncludesUint8(collection []uint8, value uint8) (found bool) {
	return FindUint8Index(collection, value) != -1
}

// FindUint16Index get the index of uint16
func FindUint16Index(collection []uint16, value uint16) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesUint16 check the value(uint16) is in collection
func IncludesUint16(collection []uint16, value uint16) bool {
	return FindUint16Index(collection, value) != -1
}

// FindUint32Index get the index of uint32
func FindUint32Index(collection []uint32, value uint32) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesUint32 check the value(uint32) is in collection
func IncludesUint32(collection []uint32, value uint32) bool {
	return FindUint32Index(collection, value) != -1
}

// FindUint64Index get the index of uint64
func FindUint64Index(collection []uint64, value uint64) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesUint64 check the value(uint64) is in collection
func IncludesUint64(collection []uint64, value uint64) bool {
	return FindUint64Index(collection, value) != -1
}

// FindInt8Index get the index of int8
func FindInt8Index(collection []int8, value int8) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesInt8 check the value(int8) is in collection
func IncludesInt8(collection []int8, value int8) bool {
	return FindInt8Index(collection, value) != -1
}

// FindInt16Index get the index of int16
func FindInt16Index(collection []int16, value int16) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesInt16 check the value(int16) is in collection
func IncludesInt16(collection []int16, value int16) bool {
	return FindInt16Index(collection, value) != -1
}

// FindInt32Index get the index of int32
func FindInt32Index(collection []int32, value int32) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesInt32 check the value(int32) is in collection
func IncludesInt32(collection []int32, value int32) bool {
	return FindInt32Index(collection, value) != -1
}

// FindInt64Index get the index of int64
func FindInt64Index(collection []int64, value int64) int {
	index := -1
	for i, v := range collection {
		if index != -1 {
			return index
		}
		if v == value {
			index = i
		}
	}
	return index
}

// IncludesInt64 check the value(int64) is in collection
func IncludesInt64(collection []int64, value int64) bool {
	return FindInt64Index(collection, value) != -1
}

// Pick get the data from struct
func Pick(current interface{}, fields []string) (data map[string]interface{}) {
	data = make(map[string]interface{})
	m := ToMap(current)
	for k, v := range m {
		if IncludesString(fields, k) {
			data[k] = v
		}
	}
	return
}

// IsString check the value type is string
func IsString(i interface{}) bool {
	return getType(i) == "string"
}

// IsInt check the value type is int
func IsInt(i interface{}) bool {
	return getType(i) == "int"
}

// IsInt8 check the value type is int8
func IsInt8(i interface{}) bool {
	return getType(i) == "int8"
}

// IsInt16 check the value type is int16
func IsInt16(i interface{}) bool {
	return getType(i) == "int16"
}

// IsInt32 check the value type is int32
func IsInt32(i interface{}) bool {
	return getType(i) == "int32"
}

// IsInt64 check the value type is int64
func IsInt64(i interface{}) bool {
	return getType(i) == "int64"
}

// IsBool check the value type is bool
func IsBool(i interface{}) bool {
	return getType(i) == "bool"
}

// IsUint check the value type is uint
func IsUint(i interface{}) bool {
	return getType(i) == "uint"
}

// IsUint8 check the value type is uint8
func IsUint8(i interface{}) bool {
	return getType(i) == "uint8"
}

// IsUint16 check the value type is uint16
func IsUint16(i interface{}) bool {
	return getType(i) == "uint16"
}

// IsUint32 check the value type is uint32
func IsUint32(i interface{}) bool {
	return getType(i) == "uint32"
}

// IsUint64 check the value type is uint64
func IsUint64(i interface{}) bool {
	return getType(i) == "uint64"
}

// IsUintptr check the value type is uintptr
func IsUintptr(i interface{}) bool {
	return getType(i) == "uintptr"
}

// IsFloat32 check the value type is float32
func IsFloat32(i interface{}) bool {
	return getType(i) == "float32"
}

// IsFloat64 check the value type is float64
func IsFloat64(i interface{}) bool {
	return getType(i) == "float64"
}

// IsComplex64 check the value type is complex64
func IsComplex64(i interface{}) bool {
	return getType(i) == "complex64"
}

// IsComplex128 check the value type is complex128
func IsComplex128(i interface{}) bool {
	return getType(i) == "complex128"
}

// GetString get the first string param from arguments
func GetString(args ...interface{}) (value string, found bool) {
	for _, v := range args {
		if !found && IsString(v) {
			value = v.(string)
			found = true
		}
	}
	return
}

// GetBool get the first bool param from arguments
func GetBool(args ...interface{}) (value bool, found bool) {
	for _, v := range args {
		if !found && IsBool(v) {
			value = v.(bool)
			found = true
		}
	}
	return
}

// GetInt get the first int param from arguments
func GetInt(args ...interface{}) (value int, found bool) {
	for _, v := range args {
		if !found && IsInt(v) {
			value = v.(int)
			found = true
		}
	}
	return
}

// GetInt8 get the first int8 param from arguments
func GetInt8(args ...interface{}) (value int8, found bool) {
	for _, v := range args {
		if !found && IsInt8(v) {
			value = v.(int8)
			found = true
		}
	}
	return
}

// GetInt16 get the first int16 param from arguments
func GetInt16(args ...interface{}) (value int16, found bool) {
	for _, v := range args {
		if !found && IsInt16(v) {
			value = v.(int16)
			found = true
		}
	}
	return
}

// GetInt32 get the first int32 param from arguments
func GetInt32(args ...interface{}) (value int32, found bool) {
	for _, v := range args {
		if !found && IsInt32(v) {
			value = v.(int32)
			found = true
		}
	}
	return
}

// GetInt64 get the first int64 param from arguments
func GetInt64(args ...interface{}) (value int64, found bool) {
	for _, v := range args {
		if !found && IsInt64(v) {
			value = v.(int64)
			found = true
		}
	}
	return
}

// GetUint8 get the first uint8 param from arguments
func GetUint8(args ...interface{}) (value uint8, found bool) {
	for _, v := range args {
		if !found && IsUint8(v) {
			value = v.(uint8)
			found = true
		}
	}
	return
}

// GetUint16 get the first uint16 param from arguments
func GetUint16(args ...interface{}) (value uint16, found bool) {
	for _, v := range args {
		if !found && IsUint16(v) {
			value = v.(uint16)
			found = true
		}
	}
	return
}

// GetUint32 get the first uint32 param from arguments
func GetUint32(args ...interface{}) (value uint32, found bool) {
	for _, v := range args {
		if !found && IsUint32(v) {
			value = v.(uint32)
			found = true
		}
	}
	return
}

// GetUint64 get the first uint64 param from arguments
func GetUint64(args ...interface{}) (value uint64, found bool) {
	for _, v := range args {
		if !found && IsUint64(v) {
			value = v.(uint64)
			found = true
		}
	}
	return
}

// Fill fill the dest with source
func Fill(dest interface{}, m map[string]interface{}, args ...interface{}) error {
	firstLetterUpper, _ := GetBool(args...)
	structValue := reflect.ValueOf(dest).Elem()
	for k, v := range m {
		name := k
		if firstLetterUpper {
			name = strings.ToUpper(k[0:1]) + k[1:]
		}
		structFieldValue := structValue.FieldByName(name)
		if !structFieldValue.IsValid() {
			return errors.New("No such field: " + k + " in dest")
		}
		if !structFieldValue.CanSet() {
			return errors.New("Cannot set " + k + " field value")
		}
		structFieldType := structFieldValue.Type()
		val := reflect.ValueOf(v)
		if structFieldType != val.Type() {
			return errors.New("Provided value type didn't match dist " + k + " type")
		}
		structFieldValue.Set(val)
	}
	return nil
}

// UniqInt The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UniqInt(collection []int) (uniq []int) {
	for _, v := range collection {
		if !IncludesInt(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqInt8 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqInt8(collection []int8) (uniq []int8) {
	for _, v := range collection {
		if !IncludesInt8(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqInt16 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqInt16(collection []int16) (uniq []int16) {
	for _, v := range collection {
		if !IncludesInt16(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqInt32 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqInt32(collection []int32) (uniq []int32) {
	for _, v := range collection {
		if !IncludesInt32(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqInt64 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqInt64(collection []int64) (uniq []int64) {
	for _, v := range collection {
		if !IncludesInt64(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqUint8 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqUint8(collection []uint8) (uniq []uint8) {
	for _, v := range collection {
		if !IncludesUint8(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqUint16 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqUint16(collection []uint16) (uniq []uint16) {
	for _, v := range collection {
		if !IncludesUint16(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqUint32 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqUint32(collection []uint32) (uniq []uint32) {
	for _, v := range collection {
		if !IncludesUint32(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqUint64 The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqUint64(collection []uint64) (uniq []uint64) {
	for _, v := range collection {
		if !IncludesUint64(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}

// UinqString The first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func UinqString(collection []string) (uniq []string) {
	for _, v := range collection {
		if !IncludesString(uniq, v) {
			uniq = append(uniq, v)
		}
	}
	return
}
