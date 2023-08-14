package util

import (
	"sort"
)

// MapSortString 对 map 进行排序
func MapSortString(m map[string]any) map[string]any {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ret := make(map[string]any, len(m))
	for _, k := range keys {
		ret[k] = m[k]
	}

	return ret
}
