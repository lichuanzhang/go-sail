package utils

import (
	"fmt"
	"math/rand"

	"github.com/shopspring/decimal"
)

// RandomInt64 在指定范围内取随机整数
//
// start和end同时支持正负数
//
// 结果值区间 ∈ [start, end)
//
// # Note
//
// 若start大于end将panic
//
// # Example:
//
// result := RandomInt64(10, 20)
// //-> 13
//
// result := RandomInt64(-10, 20)
// //-> 3
//
// result := RandomInt64(-20, -10)
// //-> -7
func RandomInt64(start, end int64) int64 {
	if start > end {
		panic(fmt.Errorf("range invalid: start great than end"))
	}
	if start == end {
		return start
	}
	//如果范围都是负值区间
	if start < 0 && end < 0 {
		fixedStart, fixedEnd := 0-start, 0-end
		return 0 - (fixedEnd + rand.Int63n(fixedStart))
	}
	//如果是一正一负
	if start < 0 && end >= 0 {
		fixed := 0 - start
		return rand.Int63n(fixed+end) - fixed
	}
	//起始为0
	if start == 0 {
		return rand.Int63n(end)
	}

	return start
}

// RandomFloat64 在指定范围内取随机浮点数
//
// start和end同时支持正负数
//
// precision为精度，此参数将限定返回值的最大小数位数
//
// 结果值区间 ∈ [start, end)
//
// # Note
//
// 若start大于end将panic
//
// # Example:
//
// result := RandomFloat64(10.10, 20.20, 2)
// //-> 16.22
//
// result := RandomFloat64(-10.10, 20.20, 3)
// //-> -7.222
//
// result := RandomFloat64(-20.20, -10.10101010101, 4)
// //-> -8.1234
func RandomFloat64(start, end float64, precision int) float64 {
	if start > end {
		panic(fmt.Errorf("range invalid: start great than end"))
	}
	if start == end {
		return start
	}
	var scale = 1

	//对start进行放大
	for {
		startScaled := start * float64(scale)
		if startScaled == float64(int64(startScaled)) {
			break
		}
		scale *= 10
	}

	//对end进行放大
	for {
		endScaled := end * float64(scale)
		if endScaled == float64(int64(endScaled)) {
			break
		}
		scale *= 10
	}

	start *= float64(scale)
	end *= float64(scale)

	randInt64 := RandomInt64(int64(start), int64(end))

	return decimal.NewFromInt(randInt64).
		Div(decimal.NewFromInt(int64(scale))).
		Truncate(int32(precision)).InexactFloat64()
}

// Pow 计算x的y次幂
//
// # Note
//
// 若y小于0,将panic
func Pow(x, y int64) int64 {
	if y < 0 {
		panic(fmt.Errorf("y less than zero"))
	}
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}
	var times int64
	for {
		times++
		if times == y {
			break
		}
		x *= x
	}

	return x
}
