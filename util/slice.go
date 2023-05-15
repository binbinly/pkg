package util

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// SliceShuffle shuffle a slice
func SliceShuffle(s []any) {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// SliceReverse slice 反转
func SliceReverse(s []any) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

// InSliceInt 是否在 slice
func InSliceInt(s int, ss []int) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// InSliceStr 是否在 slice
func InSliceStr(s string, ss []string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// SliceDeleteElem 索引删除 slice 元素
func SliceDeleteElem(i int, s []any) ([]any, error) {
	if len(s) == 0 {
		return nil, errors.New("slice is empty")
	}
	if i < 0 || i >= len(s) {
		return nil, fmt.Errorf("index %d out of range for slice of length %d", i, len(s))
	}
	s[i] = s[len(s)-1]
	s[len(s)-1] = nil
	s = s[:len(s)-1]
	return s, nil
}

// SliceIntJoin 整型数组拼接成字符串
func SliceIntJoin(s []int, sep string) string {
	if len(s) <= 1 {
		return ""
	}
	var str strings.Builder
	str.WriteString(strconv.Itoa(s[0]))
	for i := 1; i < len(s); i++ {
		str.WriteString(sep)
		str.WriteString(strconv.Itoa(s[i]))
	}

	return str.String()
}

// SliceToInt 字符串数组转换成整型数组
func SliceToInt(ss []string) (ii []int) {
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		ii = append(ii, i)
	}
	return ii
}

// SliceBigFilter 过滤切片元素 适合大切片
func SliceBigFilter(a []int, f func(v int) bool) []int {
	for i := 0; i < len(a); i++ {
		if !f(a[i]) {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

// SliceSmallFilter 过滤切片元素 适合小切片
func SliceSmallFilter(a []int, f func(v int) bool) []int {
	ret := make([]int, 0, len(a))
	for _, val := range a {
		if f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

// SliceIntDeduplication 去除重复的元素
func SliceIntDeduplication(a []int) []int {
	if a == nil {
		return nil
	}

	m := make(map[int]struct{}, len(a))
	j := 0
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		a[j] = v
		j++
	}
	return a[:j]
}
