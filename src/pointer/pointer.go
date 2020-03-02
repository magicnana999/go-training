package pointer

import "fmt"

func foo(){

	/**
	1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	2.指针变量的值是指针地址。
	3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
	 */


	a := 10					//变量a的值为10
	b := &a					//b是指针类型 *int，b的值是a的地址
	c := *b					//b是指针，*是取这个指针的地址的值赋值给c
	fmt.Println(c)

	set100(a)
	fmt.Println(a)
	set200(&a)
	fmt.Println(a)

	var p1 * string
	if p1 == nil{
		fmt.Println("空的")
	}else{
		fmt.Println("非空")
	}


	//浙江导致panic，p2未分配内存
	//var p2 * int
	//*p2 = 100
	//fmt.Println(*p2)


	/*
	new
	func new(Type) *Type			内置函数
	 1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
	2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
	*/

	p3 :=new(int)					//分配内存，返回*int，用0初始化
	p4 := new (string)			//分配内存，返回*string，用""初始化

	*p3 = 100
	*p4 = "helloworld"
	fmt.Println(*p3)
	fmt.Println(*p4+"_")

	var m1 map[string] int;
	m1 = make(map[string] int,10)
	m1["name"] = 100
	fmt.Println(m1, len(m1),m1["name"])


	var i5 int = 100
	p5 := &i5
	*p5 = 200
	fmt.Println(i5,*p5)

}

func set100(i int){
	i = 100
}

func set200(i *int){
	*i = 200
}
