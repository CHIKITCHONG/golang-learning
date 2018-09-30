package main

/*
什么是指针
指针是存储一个变量的内存地址的变量。

在上图中，变量 b 的值是 156，存储在地址为 0x1040a124 的内存中。
变量 a 存储了变量 b 的地址。现在可以说 a 指向 b。
*/

//指向类型 T 的指针用 *T 表示。


////func main() {
////	b := 255
////	// *int 是指针变量,因为b是255，所以要用*int
////	var a *int = &b
////	fmt.Printf("Type of a is %T\n", a)
////	fmt.Println("address of b is", a)
////
//}


// 指针的空值为 nil，在下面的程序中，b 的初始值为 nil。接着将 a 的地址赋值给 b。程序的输出为：


//func main() {
//	a := 25
//	var b *int
//	if b == nil {
//		fmt.Println("b is", b)
//		b = &a
//		fmt.Println("b after initialization is", b)
//	}
//}

/*

指针解引用
解引用指针的意思是通过指针访问被指向的值。指针 a 的解引用表示为：*a。
让我们通过一个程序看一下它是怎么工作的。

*/


//func main() {
//	b := 255
//	a := &b
//	fmt.Println("address of b is", a)
//	fmt.Println("value of b is", *a)
//}



/*

让我们再写一个程序，该程序使用指针改变 b 的值。

*/


//func main() {
//	b := 255
//	a := &b
//	fmt.Println("address of b is", a)
//	fmt.Println("value of b is", *a)
//	*a++
//	fmt.Println("new value of b is", b)
//}
/*
在上面的程序中，我们将 a 指向的值自增 1，这样做也改变了 b 的值，因为 a 指向 b。因此 b 的值变为 256。程序的输出为：
（理解：因为*a解了引用所以 *a = b，那么*a++ 所以b++）
*/






/*

传递指针给函数
在下面的程序中，第 14 行，我们将指向 a 的指针 b 传递给函数 change。
在函数 change 内部，第 8 行，通过解引用修改了 a 的值。程序的输出如下：
*/


//func change(val *int) {
//
//	*val = 55
//}
//
//
//func main() {
//
//	a := 58
//	fmt.Println("value of a before function call is",a)
//	b := &a
//	change(b)
//	fmt.Println("value of a after function call is", a)
//}


/*

不要传递指向数组的指针给函数，而是使用切片
假设我们需要通过函数修改一个数组。一个办法是将数组的指针作为参数传递给函数。


在上面的程序中，第13行，数组 a 的地址传递给了函数 modify。
在第8行的 modify 函数中，我们通过解引用的方式将数组的第一个元素赋值为 90。程序输出为：[90 90 91]。
*/


//func modify(arr *[3]int) {
//	(*arr)[0] = 90
//}
//
//
//func main() {
//	a := [3]int{89, 90, 91}
//	modify(&a)
//	fmt.Println(a)
//}

/*

在上面的程序中，第13行，数组 a 的地址传递给了函数 modify。
在第8行的 modify 函数中，我们通过解引用的方式将数组的第一个元素赋值为 90。程序输出为：[90 90 91]。

a[x] 是 (*a)[x] 的简写，因此上面的程序中，(*arr)[0] 可以替换为 arr[0]。让我们用这种简写方式重写上面的程序：

*/


//func modify(arr *[3]int) {
//	arr[0] = 90
//}
//
//
//func main() {
//	a := [3]int{89, 90, 91}
//	modify(&a)
//	fmt.Println(a)
//}


/*

虽然可以通过传递数组指针给函数的方式来修改原始数组的值，但这在 Go 中不是惯用的方式，我们可以使用切片完成同样的事情。

让我们用切片的方式重写上面的程序：

*/


//func modify(sls []int) {
//
//	sls[0] = 90
//}
//
//
//func main() {
//
//	a := [3]int{89, 90, 91}
//	modify(a[:])
//	fmt.Println(a)
//}


/*

Go 不支持其他语言（比如C）中的指针运算。
下面的程序将报错：main.go:6: invalid operation: p++ (non-numeric type *[3]int)

*/


//func main() {
//	b := [...]int{109, 110, 111}
//	p := &b
//	p++
//}