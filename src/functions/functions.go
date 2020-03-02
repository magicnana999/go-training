package functions

import "fmt"

func foo(){
	var x = 10
	var y = 20
	z := tt(func(x, y int) int {
		return x+y
	},x,y)

	fmt.Println(z)
}

func tt(f FuncDouble,x int,y int) int{
	return f(x,y)
}

type FuncDouble func(x,y int) int
