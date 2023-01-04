package sliutil

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Contain 检查值是否在切片中
func Contain[T comparable](slice []T, value T) bool {
	set := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	_, ok := set[value]
	return ok
}

// ContainSubSlice 检查切片是否包含子切片
func ContainSubSlice[T comparable](slice, subSlice []T) bool {
	for _, v := range subSlice {
		if !Contain(slice, v) {
			return false
		}
	}

	return true
}

// Chunk 创建一个元素切片，按照 size 的长度分成几组。
func Chunk[T any](slice []T, size int) [][]T {
	var result [][]T

	if len(slice) == 0 || size <= 0 {
		return result
	}

	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			var tmp []T
			tmp = append(tmp, v)
			result = append(result, tmp)
		}
		return result
	}

	// divide slice equally
	divideNum := length/size + 1
	for i := 0; i < divideNum; i++ {
		if i == divideNum-1 {
			if len(slice[i*size:]) > 0 {
				result = append(result, slice[i*size:])
			}
		} else {
			result = append(result, slice[i*size:(i+1)*size])
		}
	}

	return result
}

// Compact 创建一个删除零值( false、nil、0、"")的切片。
func Compact[T any](slice []T) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if !reflect.DeepEqual(v, nil) &&
			!reflect.DeepEqual(v, false) &&
			!reflect.DeepEqual(v, "") &&
			!reflect.DeepEqual(v, 0) {
			result = append(result, v)
		}
	}
	return result
}

// Concat 创建一个新的切片连接切片与任何其他切片和或值。
func Concat[T any](slice []T, values ...[]T) []T {
	result := append([]T{}, slice...)

	for _, v := range values {
		result = append(result, v...)
	}

	return result
}

// Difference 差集 在 slice中 而不在 mappedSlice 中
func Difference[T comparable](slice, comparedSlice []T) []T {
	var result []T
	for _, v := range slice {
		if !Contain(comparedSlice, v) {
			result = append(result, v)
		}
	}

	return result
}

// DifferenceBy 差集
// slice和comparedSlice先执行iteratee函数
// 比较在slice中而不在mappedSlice中
func DifferenceBy[T comparable](slice []T, comparedSlice []T, iteratee func(index int, item T) T) []T {
	originSliceAfterMap := Map(slice, iteratee)
	comparedSliceAfterMap := Map(comparedSlice, iteratee)

	result := make([]T, 0)
	for i, v := range originSliceAfterMap {
		if !Contain(comparedSliceAfterMap, v) {
			result = append(result, slice[i])
		}
	}

	return result
}

// DifferenceWith accepts comparator which is invoked to compare elements of slice to values. The order and references of result values are determined by the first slice. The comparator is invoked with two arguments: (arrVal, othVal).
func DifferenceWith[T any](slice []T, comparedSlice []T, comparator func(value, otherValue T) bool) []T {
	result := make([]T, 0)

	getIndex := func(arr []T, item T, comparison func(v1, v2 T) bool) int {
		index := -1
		for i, v := range arr {
			if comparison(item, v) {
				index = i
				break
			}
		}
		return index
	}

	for i, v := range slice {
		index := getIndex(comparedSlice, v, comparator)
		if index == -1 {
			result = append(result, slice[i])
		}
	}

	return result
}

// Equal 检查两个切片是否相等：长度相同且所有元素的顺序和值相等
func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// EqualWith 检查两个切片是否与比较器函数相等
func EqualWith[T, U any](slice1 []T, slice2 []U, comparator func(T, U) bool) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v1 := range slice1 {
		v2 := slice2[i]
		if !comparator(v1, v2) {
			return false
		}
	}

	return true
}

// Every 如果切片中的所有值都通过判断函数，则返回 true。
func Every[T any](slice []T, predicate func(index int, item T) bool) bool {
	if predicate == nil {
		panic("predicate func is missing")
	}

	var currentLength int
	for i, v := range slice {
		if predicate(i, v) {
			currentLength++
		}
	}

	return currentLength == len(slice)
}

// None 如果切片中的所有值都不通过判断函数，则返回 true。
func None[T any](slice []T, predicate func(index int, item T) bool) bool {
	if predicate == nil {
		panic("predicate func is missing")
	}

	var currentLength int
	for i, v := range slice {
		if !predicate(i, v) {
			currentLength++
		}
	}

	return currentLength == len(slice)
}

// Some 如果切片中有值通过判断函数，则返回 true。
func Some[T any](slice []T, predicate func(index int, item T) bool) bool {
	if predicate == nil {
		panic("predicate func is missing")
	}

	for i, v := range slice {
		if predicate(i, v) {
			return true
		}
	}
	return false
}

// Filter 过滤 通过判断函数的值
func Filter[T any](slice []T, predicate func(index int, item T) bool) []T {
	if predicate == nil {
		panic("predicate func is missing")
	}

	result := make([]T, 0)
	for i, v := range slice {
		if predicate(i, v) {
			result = append(result, v)
		}
	}
	return result
}

// Count 计数 通过判断函数的值
func Count[T any](slice []T, predicate func(index int, item T) bool) int {
	if predicate == nil {
		panic("predicate func is missing")
	}

	if len(slice) == 0 {
		return 0
	}

	var count int
	for i, v := range slice {
		if predicate(i, v) {
			count++
		}
	}

	return count
}

// GroupBy 遍历切片的元素，每个元素将按条件分组，返回两个切片
func GroupBy[T any](slice []T, groupFn func(index int, item T) bool) ([]T, []T) {
	if groupFn == nil {
		panic("groupFn func is missing")
	}

	if len(slice) == 0 {
		return make([]T, 0), make([]T, 0)
	}

	groupB := make([]T, 0)
	groupA := make([]T, 0)

	for i, v := range slice {
		ok := groupFn(i, v)
		if ok {
			groupA = append(groupA, v)
		} else {
			groupB = append(groupB, v)
		}
	}

	return groupA, groupB
}

// GroupWith 使用迭代函数进行分组 返回一个map
func GroupWith[T any, U comparable](slice []T, iteratee func(T) U) map[U][]T {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	result := make(map[U][]T)

	for _, v := range slice {
		key := iteratee(v)
		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		result[key] = append(result[key], v)
	}

	return result
}

// Find 遍历 slice 的元素，返回第一个通过判断函数的元素。
func Find[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	if predicate == nil {
		panic("predicate func is missing")
	}

	if len(slice) == 0 {
		return nil, false
	}

	index := -1
	for i, v := range slice {
		if predicate(i, v) {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, false
	}

	return &slice[index], true
}

// FindLast 遍历 slice 的元素，返回最后一个通过判断函数的元素。
func FindLast[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	if predicate == nil {
		panic("predicate func is missing")
	}

	if len(slice) == 0 {
		return nil, false
	}

	index := -1
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(i, slice[i]) {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, false
	}

	return &slice[index], true
}

// Flatten 将切片展平一层
func Flatten(slice any) any {
	sv := sliceValue(slice)

	var result reflect.Value
	if sv.Type().Elem().Kind() == reflect.Interface {
		result = reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, sv.Len())
	} else if sv.Type().Elem().Kind() == reflect.Slice {
		result = reflect.MakeSlice(sv.Type().Elem(), 0, sv.Len())
	} else {
		return result
	}

	for i := 0; i < sv.Len(); i++ {
		item := reflect.ValueOf(sv.Index(i).Interface())
		if item.Kind() == reflect.Slice {
			for j := 0; j < item.Len(); j++ {
				result = reflect.Append(result, item.Index(j))
			}
		} else {
			result = reflect.Append(result, item)
		}
	}

	return result.Interface()
}

// FlattenDeep 展平切片递归
func FlattenDeep(slice any) any {
	sv := sliceValue(slice)
	st := sliceElemType(sv.Type())
	tmp := reflect.MakeSlice(reflect.SliceOf(st), 0, 0)
	result := flattenRecursive(sv, tmp)
	return result.Interface()
}

func flattenRecursive(value reflect.Value, result reflect.Value) reflect.Value {
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		kind := item.Kind()

		if kind == reflect.Slice {
			result = flattenRecursive(item, result)
		} else {
			result = reflect.Append(result, item)
		}
	}

	return result
}

// ForEach 遍历 slice 的元素并为每个元素调用迭代函数
func ForEach[T any](slice []T, iteratee func(index int, item T)) {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	for i, v := range slice {
		iteratee(i, v)
	}
}

// Map 通过 iteratee 函数运行 slice 的每个元素来创建一个值切片。
func Map[T any, U any](slice []T, iteratee func(index int, item T) U) []U {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	result := make([]U, len(slice), cap(slice))
	for i, v := range slice {
		result[i] = iteratee(i, v)
	}

	return result
}

// Reduce 通过 iteratee 函数运行 slice 的每个元素来创建一个值切片。
func Reduce[T any](slice []T, iteratee func(index int, item1, item2 T) T, initial T) T {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	if len(slice) == 0 {
		return initial
	}

	result := iteratee(0, initial, slice[0])

	tmp := make([]T, 2)
	for i := 1; i < len(slice); i++ {
		tmp[0] = result
		tmp[1] = slice[i]
		result = iteratee(i, tmp[0], tmp[1])
	}

	return result
}

// InterfaceSlice 将参数转换为接口切片。
func InterfaceSlice(slice any) []any {
	sv := sliceValue(slice)
	if sv.IsNil() {
		return nil
	}

	result := make([]any, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		result[i] = sv.Index(i).Interface()
	}

	return result
}

// StringSlice 将参数转换为字符串片段。
func StringSlice(slice any) []string {
	v := sliceValue(slice)

	out := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		v, ok := v.Index(i).Interface().(string)
		if !ok {
			panic("invalid element type")
		}
		out[i] = v
	}

	return out
}

// IntSlice 将参数转换为 int 切片。
func IntSlice(slice any) []int {
	sv := sliceValue(slice)

	out := make([]int, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		v, ok := sv.Index(i).Interface().(int)
		if !ok {
			panic("invalid element type")
		}
		out[i] = v
	}

	return out
}

// DeleteAt 删除切片  start index to end index - 1.
func DeleteAt[T any](slice []T, start int, end ...int) []T {
	size := len(slice)

	if start < 0 || start >= size {
		return slice
	}

	if len(end) > 0 {
		end := end[0]
		if end <= start {
			return slice
		}
		if end > size {
			end = size
		}

		slice = append(slice[:start], slice[end:]...)
		return slice
	}

	if start == size-1 {
		slice = slice[:start]
	} else {
		slice = append(slice[:start], slice[start+1:]...)
	}
	return slice
}

// Drop 创建一个切片，当 n > 0 时从开头删除 n 个元素，或者当 n < 0 时从结尾删除 n 个元素
func Drop[T any](slice []T, n int) []T {
	size := len(slice)

	if size == 0 || n == 0 {
		return slice
	}

	if math.Abs(float64(n)) >= float64(size) {
		return []T{}
	}

	if n < 0 {
		return slice[0 : size+n]
	}

	return slice[n:size]
}

// InsertAt 将值或其他切片插入到索引处的切片中。
func InsertAt[T any](slice []T, index int, value any) []T {
	size := len(slice)

	if index < 0 || index > size {
		return slice
	}

	if v, ok := value.(T); ok {
		slice = append(slice[:index], append([]T{v}, slice[index:]...)...)
		return slice
	}

	if v, ok := value.([]T); ok {
		slice = append(slice[:index], append(v, slice[index:]...)...)
		return slice
	}

	return slice
}

// UpdateAt 更新索引处的切片元素。
func UpdateAt[T any](slice []T, index int, value T) []T {
	size := len(slice)

	if index < 0 || index >= size {
		return slice
	}
	slice = append(slice[:index], append([]T{value}, slice[index+1:]...)...)

	return slice
}

// Unique 删除切片中的重复元素。
func Unique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// here no use map filter. if use it, the result slice element order is random, not same as origin slice
	var result []T
	for i := 0; i < len(slice); i++ {
		v := slice[i]
		skip := true
		for j := range result {
			if v == result[j] {
				skip = false
				break
			}
		}
		if skip {
			result = append(result, v)
		}
	}

	return result
}

// UniqueBy 对切片的每个项目调用 iteratee func，然后删除重复项。
func UniqueBy[T comparable](slice []T, iteratee func(item T) T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	var result []T
	for _, v := range slice {
		val := iteratee(v)
		result = append(result, val)
	}

	return Unique(result)
}

// Union 多个切片合并并去重
func Union[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}

	// append all slices, then unique it
	var allElements []T

	for _, slice := range slices {
		allElements = append(allElements, slice...)
	}

	return Unique(allElements)
}

// Intersection 去交集并去重
func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
	}

	var result []T

	reducer := func(sliceA, sliceB []T) []T {
		hashMap := make(map[T]int)
		for _, val := range sliceA {
			hashMap[val] = 1
		}

		out := make([]T, 0)
		for _, val := range sliceB {
			if v, ok := hashMap[val]; v == 1 && ok {
				out = append(out, val)
				hashMap[val]++
			}
		}
		return out
	}

	result = reducer(slices[0], slices[1])

	reduceSlice := make([][]T, 2)
	for i := 2; i < len(slices); i++ {
		reduceSlice[0] = result
		reduceSlice[1] = slices[i]
		result = reducer(reduceSlice[0], reduceSlice[1])
	}

	return result
}

// SymmetricDifference 交集函数的相反运算
func SymmetricDifference[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
	}

	result := make([]T, 0)

	intersectSlice := Intersection(slices...)

	for i := 0; i < len(slices); i++ {
		slice := slices[i]
		for _, v := range slice {
			if !Contain(intersectSlice, v) {
				result = append(result, v)
			}
		}

	}

	return Unique(result)
}

// Reverse 切片元素反转
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle 切片元素打乱
func Shuffle[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range rand.Perm(len(slice)) {
		result[i] = slice[v]
	}

	return result
}

// SortByField 按字段返回排序的切片
// 切片元素应为结构体，字段类型应为 int、uint、string 或 bool
// 默认sortType是升序（asc），如果是降序，设置sortType为desc
func SortByField(slice any, field string, sortType ...string) error {
	sv := sliceValue(slice)
	t := sv.Type().Elem()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", slice)
	}

	// Find the field.
	sf, ok := t.FieldByName(field)
	if !ok {
		return fmt.Errorf("field name %s not found", field)
	}

	// Create a less function based on the field's kind.
	var compare func(a, b reflect.Value) bool
	switch sf.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Int() > b.Int() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Int() < b.Int() }
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Uint() > b.Uint() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Uint() < b.Uint() }
		}
	case reflect.Float32, reflect.Float64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Float() > b.Float() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Float() < b.Float() }
		}
	case reflect.String:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.String() > b.String() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.String() < b.String() }
		}
	case reflect.Bool:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Bool() && !b.Bool() }
		} else {
			compare = func(a, b reflect.Value) bool { return !a.Bool() && b.Bool() }
		}
	default:
		return fmt.Errorf("field type %s not supported", sf.Type)
	}

	sort.Slice(slice, func(i, j int) bool {
		a := sv.Index(i)
		b := sv.Index(j)
		if t.Kind() == reflect.Ptr {
			a = a.Elem()
			b = b.Elem()
		}
		a = a.FieldByIndex(sf.Index)
		b = b.FieldByIndex(sf.Index)
		return compare(a, b)
	})

	return nil
}

// Without 创建一个不包括所有给定值的切片
func Without[T comparable](slice []T, values ...T) []T {
	if len(values) == 0 || len(slice) == 0 {
		return slice
	}

	out := make([]T, 0, len(slice))
	for _, v := range slice {
		if !Contain(values, v) {
			out = append(out, v)
		}
	}

	return out
}

// IndexOf 返回在切片中第一次出现值的索引或返回 -1
func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}

// LastIndexOf 返回在切片中最后一次出现值的索引或返回 -1
func LastIndexOf[T comparable](slice []T, value T) int {
	for i := len(slice) - 1; i > 0; i-- {
		if value == slice[i] {
			return i
		}
	}

	return -1
}

// ToSlicePointer 返回指向可变参数转换切片的指针
func ToSlicePointer[T any](value ...T) []*T {
	out := make([]*T, len(value))
	for i := range value {
		out[i] = &value[i]
	}
	return out
}

// ToSlice 返回可变参数转换的切片
func ToSlice[T any](value ...T) []T {
	out := make([]T, len(value))
	copy(out, value)
	return out
}

// AppendIfAbsent 仅不存在附加值
func AppendIfAbsent[T comparable](slices []T, value T) []T {
	if !Contain(slices, value) {
		slices = append(slices, value)
	}
	return slices
}

// sliceValue 返回切片的反射值
func sliceValue(slice any) reflect.Value {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid slice type, value of type %T", slice))
	}
	return v
}

// FunctionValue return the reflect value of a function
func FunctionValue(function any) reflect.Value {
	v := reflect.ValueOf(function)
	if v.Kind() != reflect.Func {
		panic(fmt.Sprintf("Invalid function type, value of type %T", function))
	}
	return v
}

// CheckSliceCallbackFuncSignature Check func sign :  s :[]type1{} -> func(i int, data type1) type2
// see https://coolshell.cn/articles/21164.html#%E6%B3%9B%E5%9E%8BMap-Reduce
func CheckSliceCallbackFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	//Check it is a function
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	// first input param type should be int
	if fn.Type().In(0) != reflect.TypeOf(1) {
		return false
	}
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

// sliceElemType 获取切片元素类型
func sliceElemType(reflectType reflect.Type) reflect.Type {
	for {
		if reflectType.Kind() != reflect.Slice {
			return reflectType
		}

		reflectType = reflectType.Elem()
	}
}

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInt64ToString 将 int64 切片格式化为字符串，例如：n1,n2,n3。
func JoinInt64ToString(is []int64) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteByte(',')
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}

// SplitStringToInt64 将字符串拆分为 int64 切片。
func SplitStringToInt64(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}
	sArr := strings.Split(s, ",")
	res := make([]int64, 0, len(sArr))
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

// NilSliceToEmptySlice 递归地将 nil 切片设置为空切片
func NilSliceToEmptySlice(inter interface{}) interface{} {
	// original input that can't be modified
	val := reflect.ValueOf(inter)

	switch val.Kind() {
	case reflect.Slice:
		newSlice := reflect.MakeSlice(val.Type(), 0, val.Len())
		if !val.IsZero() {
			// iterate over each element in slice
			for j := 0; j < val.Len(); j++ {
				item := val.Index(j)

				var newItem reflect.Value
				switch item.Kind() {
				case reflect.Struct:
					// recursively handle nested struct
					newItem = reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(item.Interface())))
				default:
					newItem = item
				}

				newSlice = reflect.Append(newSlice, newItem)
			}

		}
		return newSlice.Interface()
	case reflect.Struct:
		// new struct that will be returned
		newStruct := reflect.New(reflect.TypeOf(inter))
		newVal := newStruct.Elem()
		// iterate over input's fields
		for i := 0; i < val.NumField(); i++ {
			newValField := newVal.Field(i)
			valField := val.Field(i)
			switch valField.Kind() {
			case reflect.Slice:
				// recursively handle nested slice
				newValField.Set(reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(valField.Interface()))))
			case reflect.Struct:
				// recursively handle nested struct
				newValField.Set(reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(valField.Interface()))))
			default:
				newValField.Set(valField)
			}
		}

		return newStruct.Interface()
	case reflect.Map:
		// new map to be returned
		newMap := reflect.MakeMap(reflect.TypeOf(inter))
		// iterate over every key value pair in input map
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			// recursively handle nested value
			newV := reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(v.Interface())))
			newMap.SetMapIndex(k, newV)
		}
		return newMap.Interface()
	case reflect.Ptr:
		// dereference pointer
		return NilSliceToEmptySlice(val.Elem().Interface())
	default:
		return inter
	}
}
