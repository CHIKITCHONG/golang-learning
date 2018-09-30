// rectpros.go

package rectangle


import "math"
import "fmt"


func init() {  
    fmt.Println("rectangle package initialized")
}


func Area(len, wid float64) float64 {

	area := len * wid
	return area

}


func Diagonal(len, wid float64) float64 {

	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal

}
//  如果想访问包外的函数，必须将其首字母大写。