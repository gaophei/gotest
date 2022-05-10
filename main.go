package main

import (
	"fmt"
	"math"
)

//先易后难，打印实心菱形，边长为n
//行i从下到上为-(n-1)到n-1
//列j从左到右为-(n-1)到n-1
//|i| + |j| <= n-1
func lingxing01(n int) {
	for i := -(n - 1); i <= n-1; i++ {
		for j := -(n - 1); j <= n-1; j++ {
			if int(math.Abs(float64(i)))+int(math.Abs(float64(j))) <= n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//先易后难，打印空心菱形
//当|i| + |j| = n-1 是为*，其余全部为空格
func lingxing02(n int) {
	for i := -(n - 1); i <= n-1; i++ {
		for j := -(n - 1); j <= n-1; j++ {
			if int(math.Abs(float64(i)))+int(math.Abs(float64(j))) == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var num int
	fmt.Println("plz input a int number:")
	fmt.Scanln(&num)
	lingxing01(num)
	fmt.Println("===========")
	lingxing02(num)
}
