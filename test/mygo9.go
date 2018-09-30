package main
//切片
import "fmt"

/*
元素类型为 T 的切片表示为： []T。
通过 a[start:end] 这样的语法创建了一个从 a[start] 到 a[end -1] 的切片。
在上面的程序中，第 9 行 a[1:4] 创建了一个从 a[1] 到 a[3] 的切片。因此 b 的值为：[77 78 79]。
*/


/*
下面是创建切片的另一种方式：
*/

//func main() {
//
//	c := []int{76, 77, 78, 79, 80}
//	fmt.Println(c)
//	// 在上面的程序中，第 9 行 c := []int{6, 7, 8} 创建了一个长度为 3 的 int 数组，并返回一个切片给 c。
//}

/*

切片本身不包含任何数据。它仅仅是底层数组的一个上层表示。对切片进行的任何修改都将反映在底层数组中。

*/

//func main() {
//
//	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
//	dslice := darr[2:5]
//	fmt.Println("array before", darr)
//	for i := range dslice {
//		dslice[i]++
//	}
//	fmt.Println("array after", darr)
//}

/*

当若干个切片共享同一个底层数组时，对每一个切片的修改都会反映在底层数组中。

*/

//func main() {
//	numa := [3]int{78, 79, 80}
//	nums1 := numa[:] //creates a slice which contains all elements of the array
//	nums2 := numa[:]
//	nums1[0] = 100
//	nums2[0] = 99
//	fmt.Println(numa)
//}



/*

切片的长度是指切片中元素的个数。切片的容量是指从切片的起始元素开始到其底层数组中的最后一个元素的个数。

（译者注：使用内置函数 cap 返回切片的容量。）

让我们写一些代码来更好地理解这一点。

*/

//func main() {
//	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
//	fruitslice := fruitarray[1:3]
//	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
//}



/*

内置函数 func make([]T, len, cap) []T 可以用来创建切片，该函数接受长度和容量作为参数，返回切片。
容量是可选的，默认与长度相同。使用 make 函数将会创建一个数组并返回它的切片。
用 make 创建的切片的元素值默认为 0 值。

*/


//func main() {
//	i := make([]int, 5, 5)
//	fmt.Println(i)
//}



/*

追加元素到切片

我们已经知道数组是固定长度的，它们的长度不能动态增加。而切片是动态的，可以使用内置函数 append 添加元素到切片。
append 的函数原型为：append(s []T, x ...T) []T。

x …T 表示 append 函数可以接受的参数个数是可变的。这种函数叫做变参函数。

你可能会问一个问题：如果切片是建立在数组之上的，而数组本身不能改变长度，那么切片是如何动态改变长度的呢？
实际发生的情况是，当新元素通过调用 append 函数追加到切片末尾时，如果超出了容量，append 内部会创建一个新的数组。
并将原有数组的元素被拷贝给这个新的数组，最后返回建立在这个新数组上的切片。这个新切片的容量是旧切片的二倍
(译者注：当超出切片的容量时，append 将会在其内部创建新的数组，该数组的大小是原切片容量的 2 倍。
最后 append 返回这个数组的全切片，即从 0 到 length - 1 的切片)。下面的程序使事情变得明朗：

 */




//func main() {
//	cars := []string{"Ferrari", "Honda", "Ford"}
//	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
//	cars = append(cars, "Toyota")
//	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
//}


/*
在上面的程序中，cars 的容量开始时为 3。在第 10 行我们追加了一个新的元素给 cars，并将 append(cars, "Toyota")
的返回值重新复制给 cars。现在 cars 的容量翻倍，变为 6。上面的程序输出为：
*/



//切片的 0 值为 nil。一个 nil 切片的长度和容量都为 0。可以利用 append 函数给一个 nil 切片追加值。


//func main() {
//	var names [] string
//	if names == nil {
//		fmt.Println("names, here")
//		names = append(names, "john", "alice", "alan")
//		fmt.Println(names)
//	}
//}


/*

可以使用 ... 操作符将一个切片追加到另一个切片末尾：

*/

//func main() {
//
//	vaggies := []string{"potato", "branjal", "tomatoes"}
//	fruits := []string{"apples", "oranges"}
//	foods := append(vaggies, fruits...)
//	fmt.Println("foods:", foods)
//
//}




/*

同数组一样，切片也可以有多个维度。

*/


func main() {
	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	fmt.Println(pls)
	for _, v1 := range pls {
		//fmt.Println(v1)
		for k, v2 := range v1 {
			fmt.Println(k, v2)
		}
		fmt.Printf("\n")
	}
}








