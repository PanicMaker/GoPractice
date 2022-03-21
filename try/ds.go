package try

import (
	"fmt"
	"math"
)

/*
	获取 Go 中基本数据类型的范围
*/
func dataLength() string {
	fmt.Println(math.MaxInt8)    // byte
	fmt.Println(math.MaxInt16)   // short
	fmt.Println(math.MaxInt32)   // int
	fmt.Println(math.MaxInt64)   // long
	fmt.Println(math.MaxFloat32) // float
	fmt.Println(math.MaxFloat64) // double

	return "data length"
}
