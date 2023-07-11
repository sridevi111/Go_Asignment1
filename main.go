// You can edit this code!
// Click here and start typing.
package main

import ("fmt"
"golang-jul-2023/calc"
"time"


)


func main() {
	sum := calc.Add(2,3)
	fmt.Println(sum)

	diff :=calc.Diff(4,9)
	fmt.Println(diff)

	fmt.Println(time.Now())
}



