package varconst

import "fmt"

func varandconst(a,b int){
	fmt.Printf("a:%d,b:%d",a,b)

	var a1,a2 int
	fmt.Printf("a:%d,b:%d",a1,a2)

	var (
		a3 string = "hello"
		a4 bool = true
	)

	fmt.Println(a3,a4)

	c := 100
	d := true

	fmt.Println(c,d)

}
