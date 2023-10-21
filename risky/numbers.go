package risky

import (
	"strconv"
)

func Int(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func Uint(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 64)
	return uint(i)
}

func Bool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func Float(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func Float64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
