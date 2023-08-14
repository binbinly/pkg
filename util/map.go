package util

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// MapBuildQuery map => url query
func MapBuildQuery(m map[string]any) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var params []string
	for _, k := range keys {
		switch v := m[k].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, string:
			params = append(params, fmt.Sprintf("%s=%v", k, v))
		default:
			j, _ := json.Marshal(v)
			params = append(params, fmt.Sprintf("%s=%v", k, string(j)))
		}
	}
	return strings.Join(params, "&")
}
