package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Println("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Println("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Println("Type: %T Value: %v\n", z, z)
}

// Không được phép dùng dấu ';' để kết thúc câu lệnh

// Khi export một package, tên hàm phải viết hoa chữ cái đầu tiên

// Khi convert về kiểu dữ liệu nhỏ hơn so với kiểu dữ liệu ban đầu sẽ bị bỏ đi phần thập phân

// Conts không thể gán với một cú pháp

