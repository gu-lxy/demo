package main

import "fmt"

func main() {
	// 常量组
	const (
		x int = 10 // x常量，int类型，数值10
		//如果常量组中不指定常量的数据类型和初始值，那么该常量则与上一行非空常量类型和数值相同
		y
		z
		//y int = 10 // y常量，int类型，数值10
		//z int = 10 // z常量, int类型, 数值10
		name1 string = "yuhongwei"
		sex1           // sex1的值是 “yuhongwei”
		//sex1 string = "男"
	)
	//fmt.Println(x,y,z)
	//fmt.Println(sex1)


	//一种特殊的常量：iota
	// 失眠：属羊 。1只羊，2只羊，3只羊....
	// 计数：1、2、3、4..
    // 计算机当中，go语言程序当中如何去计数？
    const (
       a = iota  // iota的默认值是0，从0开始计数
       b = iota  // iota 会在0的基础上 +1， 此时iota的值是1
       c = iota  // iota 会在1的基础上 +1,  此时iota的值是2
	)
	fmt.Println(a)  // a的值是0
	fmt.Println(b)  // b的值是1
	fmt.Println(c)  // c的值是2

	//创建另外一组常量
	//iota的第二个特点： 当iota遇到const定义常量的时候，iota的值会再次初始为0
	//失眠： 1、2、3、4... 0、1、2、3... 0、1、2、3、..
	const (
		age = iota //
		age1 = iota
		age2 = iota
	)
	fmt.Println("age的值是：",age)
	fmt.Println("age1的值是：",age1)
	fmt.Println("age2的值是：",age2)

	/**
	const (
		num  = iota
		num1 = iota
		num2 = iota
	)
	 */

	//fmt.Println( num + num1 ) // 结果 ? 1： num值是0 ，num1的值1
	//fmt.Println(num1 * num2) // 结果 ? 2 ： num1的值1， num2的值2


    //iota的变形写法
	const (
		num  = iota  // iota值恢复为0
		num1         // 等同于 num1 = iota  1
		num2         // 等同于 num2 = iota   2
	)
	fmt.Println(num)
	fmt.Println(num1) // 结果 ?
	fmt.Println(num2) // 结果 ?


	const (
		num3 = iota        // 0
		name = "yuhongwei"  // iota+1 ： iota = 1
		num4 = iota  // iota = 2
		address = "山东省泰安市泰山风景区"   // iota + 1: iota 3
		num5  = iota  // iota+ 1 :iota 4
		num6  //iota 5
	)
	fmt.Println(num3,num4,num5,num6) // 输出结果？

}
