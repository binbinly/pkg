package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ParseAmount 金额分转元，保留两位小数输出
func ParseAmount(amount int) float64 {
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(amount)/100), 64)
	return f
}

// FormatAmount 格式化金额为分，入库保存
func FormatAmount(amount float64) int {
	return int(math.Ceil(amount * 100))
}

// FormatResUrl 格式化图片资源完整路劲
func FormatResUrl(dfs, url string) string {
	return strings.Join([]string{dfs, url}, "/")
}
