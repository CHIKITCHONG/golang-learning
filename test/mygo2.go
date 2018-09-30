// package main

/*

bool 类型表示真假值，只能为 true 或 false。请运行下面的程序：
 */

// func main() {
// 	var a bool = true
// 	b := false
// 	c := a && b
// 	d := a || b
// 	println(a, b, c, d)
// }


/*
Go是强类型的语言，没有隐式的类型提升和转换。让我们通过一个例子说明这意味着什么


在C语言中是完全合法的，但是在Go中却不是。i 的类型是 int 而 j 的类型是 float64，将这两个类型不同的数字相加是非法的。
运行这个程序将会报错：main.go:10: invalid operation: i + j (mismatched types int and float64)。

*/
// package main

// import (

// 	"fmt"

// )

// func main() {

// 	i :=55
// 	j := 67.8
// 	sum := i +j  //int + floate64 not allowed
// 	fmt.println(sum)

// }

/*

为了修复这个错误，我们应该将 i 和 j 转换为同样的类型，在这里让我们将 j 转换为 int。通过 T(v)可以将 v 的值转换为 T 类型 。

*/

// package main

// import (  
//     "fmt"
// )

// func main() {  
//     i := 55      //int
//     j := 67.8    //float64
//     sum := i + int(j) //j is converted to int
//     fmt.Println(sum)
// }



/*

在赋值时情况也是如此，将一个变量赋值给另一个类型不同的变量时必须显式转型。下面的程序说明了这一点：


*/


package main

import (  
    "fmt"
)
//在 var j float64 = float64(i) 这一行，i 被转换为 float64，然后赋值给 j。
//当你尝试将 i 不进行转换直接赋值给 j 时，编译器将报错。


func main() {  
    i := 10
    var j float64 = float64(i) //this statement will not work without explicit conversion
    fmt.Println("j", j)
}

