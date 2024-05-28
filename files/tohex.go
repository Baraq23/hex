package tohexa

import (
	"fmt"
)

func ToHexadecimal(str string) int {
	hx := Atoi(str)
	fmt.Sprintln(hx)
	return hx
}

func Atoi(str string) int {
	num := 0
	for _, v := range str {
		// fmt.Println(num)
		num = ((num*10) + int(v - '0'))
	}
	fmt.Println(num)
	Hex(num)
	return num
}

func Hex(num int) string {
	
	base16 := "0123456789abcdef"
	hex := ""
	for num/16 != 0 {
		rem := num%16
		hex = string(base16[rem]) + hex
		num = num/16
	}
	fmt.Println(hex)
	return hex
}