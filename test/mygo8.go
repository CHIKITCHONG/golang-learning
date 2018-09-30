// 数组

// 数组的类型为 n[T]，其中 n 表示数组中元素的个数，T 表示数组中元素的类型。
// 元素的个数 n 也是数组类型的一部分（我们将在稍后详细讨论） 

package main


import (

	"fmt"

)


/*----------------
var a [3]int 声明了一个长度为 3 的整型数组。数组中的所有元素都被自动赋值为元素类型的 0 值。
比如这里 a 是一个整型数组，因此 a 中的所有元素都被赋值为 0（即整型的 0 值）。运行上面的程序，输出为：[0 0 0]。
------------------*/

// func main() {

// 	var a [3]int  	//int array with length 3
// 	fmt.Println(a)
// }


/*--------------
数组的索引从 0 开始到 length - 1 结束。下面让我们给上面的数组赋一些值。
-----------------*/

// func main() {

// 	var a [3]int
// 	a[0] = 12
// 	a[1] = 78
// 	a[2] = 50
// 	fmt.Println(a)
// }



/*--------------------------

可以利用速记声明（shorthand declaration）的方式来创建同样的数组：

––––––––––––––––––––––––*/
// func main() {

// 	a := [3]int{12, 78, 50}
// 	fmt.Println(a)
// }

/*
在速记声明中，没有必要为数组中的每一个元素指定初始值。
*/

// func main() {
// 	a := [3]int{12}
// 	fmt.Println(a)
// }



/*

在声明数组时你可以忽略数组的长度并用 ... 代替，让编译器为你自动推导数组的长度。比如下面的程序：

*/

// func main() {

// 	a := [...]int{12, 78, 50} 			//... makes the compiler determine the length
// 	fmt.Println(a)
// }



/*

在 Go 中数组是值类型而不是引用类型。这意味着当数组变量被赋值时，
将会获得原数组（译者注：也就是等号右面的数组）的拷贝。新数组中元素的改变不会影响原数组中元素的值。

*/

// func main() {

// 	a := [...]string{"USA", "CHINA", "INDIA", "GERMANY", "FRANCE"}
// 	b := a 			// a copy of a is assigned to b
// 	b[0] = "Sigapore"
// 	fmt.Println("a is ", a)
// 	fmt.Println("b is ", b)
// }

/*输出结果
a is  [USA CHINA INDIA GERMANY FRANCE]
b is  [Sigapore CHINA INDIA GERMANY FRANCE]
*/


/*
数组的长度
内置函数 len 用于获取数组的长度：
*/

// func main() {
// 	a := [...]float64{67.7, 89.8, 21, 78}
// 	fmt.Println("length of a is", len(a))
// }



// for循环可以遍历数组中的元素（索引从0到length(n-1）)

// func main() {
// 	a := [...]float64{67.7, 89.8, 21, 78}
// 	for i := 0; i < len(a); i++ {
// 		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
// 	}
// }



/*

Go 提供了一个更简单，更简洁的遍历数组的方法：使用 range for。range 返回数组的索引和索引对应的值。
让我们用 range for 重写上面的程序（除此之外我们还计算了数组元素的总和）。

*/

// func main() {

// 	a := [...]float64{67.7, 89.8, 21, 78}
// 	sum := float64(0)

// 	for i, v := range a {
// 		fmt.Printf("%d the element of a is %.2f\n", i, v)
// 		sum += v
// 	}
// 	fmt.Println("\nsum of all elements of a",sum)

// }



/*

如果你只想访问数组元素而不需要访问数组索引，则可以通过空标识符来代替索引变量：下面的代码忽略了索引。同样的，也可以忽略值。
*/
// for _,v :=range a {

// } 






/*
---------------------
上面的程序中，第 17 行利用速记声明创建了一个二维数组 a。
第 20 行的逗号是必须的，这是因为词法分析器会根据一些简单的规则自动插入分号。
如果你想了解更多，请阅读：https://golang.org/doc/effective_go.html#semicolons 。
在第 23 行声明了另一个二维数组 b，并通过索引的方式给数组 b 中的每一个元素赋值。这是初始化二维数组的另一种方式。
第 7 行声明的函数 printarray 通过两个嵌套的 range for 打印二维数组的内容。上面程序的输出为：
---------------------
*/


func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}


func main() {

    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    fmt.Println(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    fmt.Println(b)
    printarray(b)

}
/*

输出结果如下
lion tiger 
cat dog 
pigeon peacock 
[[lion tiger] [cat dog] [pigeon peacock]]

[[apple samsung] [microsoft google] [AT&T T-Mobile]]
apple samsung 
microsoft google 
AT&T T-Mobile 

*/


















