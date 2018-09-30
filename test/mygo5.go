/*
init 函数可用于执行初始化任务，也可用于在执行开始之前验证程序的正确性。
一个包的初始化顺序如下：
包级别的变量首先被初始化
接着 init 函数被调用。一个包可以有多个 init 函数（在一个或多个文件中），它们的调用顺序为编译器解析它们的顺序。
如果一个包导入了另一个包，被导入的包先初始化。
尽管一个包可能被包含多次，但是它只被初始化一次。
下面让我们对我们的程序做一些修改来理解 init 函数。
首先在 rectprops.go 中添加一个 init 函数：
*/


// ---------------



/*
使用空指示符
在 Go 中只导入包却不在代码中使用它是非法的。如果你这么做了，编译器会报错。
这样做的原因是为了避免引入过多未使用的包而导致编译时间的显著增加。
将 geometry.go 中的代码替换为如下代码：
*/


// 示例 eg


//geometry.go
// package main 


// import (   

//      "geometry/rectangle" //importing custom package

// )
// func main() {

// }

/*

上面的程序将会报错：geometry.go:6: imported and not used: "geometry/rectangle"
但是在开发过程中，导入包却不立即使用它是很常见的。可以用空指示符（_）来处理这种情况。
下面的代码可以避免抛出上面的错误：

*/
package main

import ( 

    "test/rectangle" 

)

var _ = rectangle.Area //error silencer

func main() {

}

/*

var _ = rectangle.Area 这一行屏蔽了错误。
我们应该跟踪这些“错误消音器”（error silencer),在开发结束时，我们应该去掉这些“错误消音器”，
并且如果没有使用相应的包，这些包也应该被一并移除。
因此，建议在 import 语句之后的包级别中写“错误消音器”。


有时我们导入一个包只是为了确保该包初始化的发生，而我们不需要使用包中的任何函数或变量。
例如，我们也许需要确保 rectangle 包的 init 函数被调用而不打算在代码中的任何地方使用这个包。
空指示符仍然可以处理这种情况，像下面的代码一样：
*/




/*

有init初始化函数时，在包上引用_ 上面的写法和下面的写法都可以

*/


// package main 


// import (   

//      _  "test/rectangle" 

// )


// func main() {

// }
