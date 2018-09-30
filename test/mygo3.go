package main

/*

常量（constant）表示固定的值，比如：5，-89，"I love Go"，67.89 等等。
关键字 const 修饰的名字为常量，不能被重新赋予任何值。 因此下面的程序会报错：cannot assign to a。
*/

// func main() {

// 	const a = 55
// 	a = 89			//reassigment not allowed

// }


/*

常量的值必须在编译期间确定。因此不能将函数的返回值赋给常量，因为函数调用发生在运行期。
会报错：error main.go:11: const initializer math.Sqrt(4) is not a constant.
*/


// import (  
//     "fmt"
//     "math"
// )

// func main() {  
//     fmt.Println("Hello, playground")
//     var a = math.Sqrt(4)					//allowed
//     const b = math.Sqrt(4)				//not allowed
// }


/*

字符串常量最简单的常量，通过了解字符串常量可以更好的理解常量的概念。
在Go中任何用双引号（"）括起来的值都是字符串常量，比如 "Hello World"，"Sam" 都是字符串常量。
字符串常量是什么类型呢？答案是 无类型（untyped）。
像 "Hello World" 这样的字符串没有任何类型

*/


// const hello = "Hello World"      			// hello没有任何类型


// 下面的代码中，变量 name 被赋值为一个无类型的常量 "Sam"，它是如何工作的呢？
//-----------------------
// func main() {  
//     fmt.Println("Hello, playground")
//     var name = "Sam"  // 译者注：这里编译器需要推导出 name 的类型，
//                       // 那么它是如何从无类型的常量 "Sam" 中获取类型的呢？
//     fmt.Printf("type %T value %v", name, name)

// }

//-----------------
// 答案是无类型常量有一个默认的类型，当且仅当代码中需要无类型常量提供类型时，它才会提供该默认类型。
// 在语句 var name = "Sam" 中，name需要一个类型，因此它从常量 "Sam" 中获取其默认类型：string。
// （译者注：这里可以理解为常量是无类型的，但是在需要类型的场合，编译器会根据常量的值和上下文将常量转换为相应的类型。）

//--------------------
// 有没有办法创建一个有类型（typed）的常量？答案是：Yes。下面的代码创建了一个有类型常量。
//-------------------------
// const typedhello string = "Hello World" 




//Go是强类型语言。在赋值时混合使用类型是不允许的。让我们通过以下代码说明这是什么意思。
// ------------

func main() {  
        var defaultName = "Sam" //allowed
        type myString string
        var customName myString = "Sam" //allowed
        customName = defaultName //not allowed

}

// ------------------
//在上面的代码中，我们首先创建了一个变量 defaultName 并且赋值为常量 "Sam"。常量 "Sam" 的默认类型为 string，
// 因此赋值之后，defaultName 的类型亦为 string。
// 下一行我们创建了一个新的类型 myString，它是 string 的别名。
// （译者注：可以通过 type NewType Type 的语法来创建一个新的类型。类似 C 语言的 typedef 。）
// 接着我们创建了一个名为 customName 的 myString 类型的变量，并将常量 "Sam" 赋给它。因为常量 "Sam" 
// 是无类型的所以可以将它赋值给任何字符串变量。因此这个赋值是合法的，customName 的类型是 myString。
// 现在我们有两个变量：string 类型的 defaultName 和 myString 类型的 customName。尽管我们知道 myString 
// 是 string 的一个别名，但是Go的强类型机制不允许将一个类型的变量赋值给另一个类型的变量。因此， customName = defaultName 这个赋值是不允许的，编译器会报错：main.go:10: cannot use defaultName (type string) as type myString in assignment
// -----------------



/*

布尔常量与字符串常量（在概念上）没有区别。布尔常量只包含两个值：true 和 false。
字符串常量的规则也同样适用于布尔常量，这里不再赘述。下面的代码解释了布尔常量：

*/

package main

func main() {  
    const trueConst = true
    type myBool bool
    var defaultBool = trueConst //allowed
    var customBool myBool = trueConst //allowed
    defaultBool = customBool //not allowed
}

// ------------

// 数值表达式
// 数值常量可以在表达式中自由的混合和匹配，仅当将它们赋值给变量或者在代码中明确需要类型的时候，才需要他们的类型。

package main

import (  
    "fmt"
)

func main() {  
    var a = 5.9/8
    fmt.Printf("a's type %T value %v",a, a)
}


