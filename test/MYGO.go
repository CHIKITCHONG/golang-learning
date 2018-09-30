package main

// import (

// 	"fmt"

// )


/*
package main：每个 Go 文件都必须以 package name 语句开头。包（package）提供了代码封装和重用。这里包的名字为 main。

import "fmt"：导入 fmt 包，在 main 函数中将使用这个包打印文本到标准输出。

func main()：main函数是一个特殊的函数，它是 Go 程序的入口点。main 函数必须包含在 main package 中。 { 和 } 表示 main 函数的开始和结束。
 */


/*
语句 var age int 声明了一个类型为 int，名称为 age 的变量。
在这里我们没有给它赋任何值。如果一个变量没有被赋予任何值，Go 会自动将这个变量初始化为其类型的 0值，
比如这里的 age 将被赋值为 0（译者注：int的0值为0）。运行这个程序，将得到如下输出：
*/
// func main() {
// 	var age int
// 	fmt.Println("my age is", age)
// }



/*
多个变量可以在一条语句中声明。
多变量声明的语法为：var name1, name2 type = initialvalue1, initialvalue2，例如：
*/

// func main() {

// 	var width, height int = 100, 50
// 	fmt.Println("width is", width, "heigth is", height)

// }



/*
正如你猜想的那样，如果不指定 width 和 height 的初值，
它们将自动被赋值为 0，也就是说它们将以 0 作为初值：
*/

// func main() {  
//     var width, height int
//     fmt.Println("width is", width, "height is", height)
//     width = 100
//     height = 50
//     fmt.Println("new width is", width, "new height is ", height)
// }




// Go 提供了另一种简洁的声明变量的方式。这种方式称为速记声明（shorthand declaratiion）。速记声明使用 := 操作符来声明变量。
// name := initialvalue


// func main() {

// 	name, age := "naveen", 29
// 	fmt.Println("my name is", name, "age is", age)

// }



/*

速记声明要求在 := 的左边必须至少有一个变量是新声明的。考虑如下程序：

*/


// func main() {  
//     a, b := 20, 30 // declare variables a and b
//     fmt.Println("a is", a, "b is", b)
//     b, c := 40, 50 // b is already declared but c is new
//     fmt.Println("b is", b, "c is", c)
//     b, c = 80, 90 // assign new values to already declared variables b and c
//     fmt.Println("changed b is", b, "c is", c)
// }


/*

 但是当我们运行下面的程序：将会报错：8: no new variables on left side of := 。
这是因为变量 a 和变量 b 都是已经声明过的变量，在 := 左侧并没有新的变量被声明。

*/


// func main() {  
//     a, b := 20, 30 //a and b declared
//     fmt.Println("a is", a, "b is", b)
//     a, b := 40, 50 //error, no new variables
// }


/*

变量可以被赋予运行时产生的值。考虑下面的程序：

*/

// package main
// import (

// 	"fmt"
// 	"math"

// )


// func main() {

// 	a,b := 145.8, 543.8
// 	// c 的值为 a 和 b 的最小值，该值是在运行时计算的。
// 	c := math.Min(a,b)
// 	fmt.Println("minumum value is", c)
// }





