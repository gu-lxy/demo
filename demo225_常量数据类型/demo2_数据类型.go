package main

import "fmt"

func main() {

	/**
	   数值类型:
	        整数类型:
	             有符号：intbit
	                intN：数字N代表的是二进制数的位数，N是2的幂次方，最小的是2^3,就是int8。比如int8表示的就是8个二进制位能表示int数据范围
	                int8 范围： -128 ~ +127
	                int16 范围：2^4
	                int32 范围: 2^5
	                int64 范围：2^6
	             无符号：uintbit   un:非
	                uint8 ： 0 - 255
	                uint16 : 0 - 65535
	                uint32 :
	        浮点数类型：
	             float32, float64
	        复数：
	    布尔类型: bool
	 */
	var num int8
	num = 18
	fmt.Printf("num的数值是:%d,数据类型是:%T",num,num)

	var age int16
	age = 20
	fmt.Printf("num的数值是:%d,数据类型是:%T",age,age)

    var result bool // 定义一个布尔类型的值
    result = true // 赋值
    fmt.Println(result)

    result = false
	fmt.Println(result)

}
