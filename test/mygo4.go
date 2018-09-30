package main

/*

函数声明以关键字 func 开头，后面是函数名字，接着是在 ( 和 ) 之间指定的参数列表，然后是函数的返回类型。
指定参数的语法为参数名称后面跟着参数类型。可以指定任意数量的参数，形式为：
 (parameter1 type, parameter2 type)。最后是由 { 和 } 以及它们之间的代码块组成的函数体。

*/


// 让我们写一个函数，它以单个产品的价格和产品数量作为输入参数，并以总价格（单个产品的价格与产品数量的乘积）作为返回值。
 //如果连续的参数具有相同的类型，我们可以避免每次都写出它们的类型，只需要在结束的时候写一次就可以了，
 //比如：price int, no int 可以写成：price, no int。 于是上面的函数可以写成：
//-----------------

// import (

//  	"fmt"

// )


// func calculateBill(price, no int) int {

// 	var totalPrice = price * no
// 	return totalPrice
// }


// func main() {

// 	price, no := 10, 6
// 	totalPrice := calculateBill(price, no)
// 	fmt.Println("total price is", totalPrice)

// }


// ------------
// 一个函数可以返回多个值。让我们写一个函数 rectProps，它接受一个矩形的长和宽，并返回该矩形的面积和周长。
// 矩形的面积为长与宽的积。周长为长与宽的和的 2 倍。
// 如果一个函数有多个返回值，那么这些返回值应该用小括号()括起来，比如：func rectProps(length, width float64)(float64, float64)
// 接受两个类型为 float64 的参数（length 和 width），并且同样返回两个类型为 float64 的返回值。上面程序的输出为：Area 60.480000 Perimeter 32.800000
// ----------------


// import (  

//     "fmt"
// )

// func rectProps(length, width float64)(float64, float64) {  

//     var area = length * width
//     var perimeter = (length + width) * 2
//     return area, perimeter

// }

// func main() {  

//  	area, perimeter := rectProps(10.8, 5.6)
//     fmt.Printf("Area %f Perimeter %f", area, perimeter) 

// }





// --------------
/*
具名返回值
可以给一个函数的返回值指定名字。如果指定了一个返回值的名字，则可以视为在该函数的第一行中定义了该名字的变量。
上面的 rectProps 函数可以用具名返回值的形式重写如下：
*/
// ------------------

// func rectProps(length, width float64)(area, perimeter float64) {  
//     area = length * width
//     perimeter = (length + width) * 2
//     return //no explicit return value
// }

/*------------
在上面的函数中，area 和 perimeter 是具名返回值。注意 return 语句没有指定任何返回值。因为在函数声明时已经指定 area 和 perimeter 是返回值，在遇到 return 语句时它们会自动从函数中返回。
（译者注：在Go中，有返回值的函数，无论是具名返回值还是普通形式的返回值，函数中必须包含 return 语句。）
*/



/*-----------------------------
空指示符
在 Go 中，下划线（_）表示空指示符（blank identifier）。它可以用于代替任何类型的任何值。让我们看看如何使用空指示符。
我们知道上面定义的函数 rectProps 返回矩形的面积（area）和周长（perimeter）。如果我们只需要获取 area 而想要忽略 perimeter该怎么办呢？这时候就可以使用空指示符。
下面的程序仅接收 rectProps 返回的 area：
------------------------------*/



// import (

// 	"fmt"

// )

// func rectProps(length, width float64) (float64, float64) {

// 	var area = length * width
// 	var perimeter = (length + width) * 2
// 	return area, perimeter

// }

// func main() {

// 	area, _ := rectProps(10.8, 5.6)
// 	fmt.Printf("Area %f", area)

// }













