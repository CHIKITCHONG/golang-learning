package main

import "fmt"

/*
什么是 map？

Map 是 Go 中的内置类型，它将键与值绑定到一起。可以通过键获取相应的值。

如何创建 map？

可以通过将键和值的类型传递给内置函数 make 来创建一个 map。语法为：make(map[KeyType]ValueType)。（译者注：map 的类型表示为 map[KeyType]ValueType）例如：


personSalary := make(map[string]int)
上面的代码创建了一个名为 personSalary 的 map。其中键的类型为 string，值的类型为 int。
*/


//func main() {
//	var personSalary map[string]int
//	if personSalary == nil {
//		fmt.Println("map is nil. Going to make one.")
//		personSalary = make(map[string]int)
//	}
//}



/*

向 map 中插入元素
插入元素给 map 的语法与数组相似。下面的代码插入一些新的元素给 map personSalary。

*/


//func main() {
//	personSalary := make(map[string]int)
//	personSalary["steve"] = 12000
//	personSalary["jamie"] = 15000
//	personSalary["mike"] = 9000
//	fmt.Println("personSalary map contents:", personSalary)
//}


//也可以在声明时初始化一个数组：


//func main() {
//	personSalary := map[string]int {
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	fmt.Println("personSalary map contents:", personSalary)
//}

/*

我们如何检测一个键是否存在于一个 map 中呢？可以使用下面的语法：
value, ok := map[key]
上面的语法可以检测一个特定的键是否存在于 map 中。如果 ok 是 true，
则键存在，value 被赋值为对应的值。如果 ok 为 false，则表示键不存在。

*/


//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	newEmp := "joe"
//	value, ok := personSalary[newEmp]
//	if ok == true {
//		fmt.Println("Salary of", newEmp, "is", value)
//	} else {
//		fmt.Println(newEmp,"not found")
//	}
//
//}


//在上面的程序中，第 15 行，ok 应该为 false 因为 joe 不存在。因此程序的输出为：
//joe not found


/*

range for 可用于遍历 map 中所有的元素（译者注：这里 range 操作符会返回 map 的键和值）。

*/


//func main() {
//	personSalary := map[string]int{
//		"steve": 12000,
//		"jamie": 15000,
//	}
//	personSalary["mike"] = 9000
//	fmt.Println("All items of a map")
//	for key, value := range personSalary {
//		fmt.Printf("personSalary[%s] = %d\n", key, value)
//	}
//
//}
//



/*

delete(map, key) 用于删除 map 中的 key。delete 函数没有返回值。

*/


func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)

}

