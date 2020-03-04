package arrayslice

import (
	"fmt"
)

/**
	1. 数组：是同一种数据类型的固定长度的序列。
    2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
    3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
    4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
    for i := 0; i < len(a); i++ {
    }
    for index, v := range a {
    }
    5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
    6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
    7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
    8.指针数组 [n]*T，数组指针 *[n]T。
 */


func foo(){
	var a0 [1] int = [1] int{1}
	fmt.Println(a0,a0[0], len(a0))

	var b = [2] int {10,20}
	fmt.Println(b,b[0], len(b), cap(b))

	var c = [...] int {10,20,30}
	fmt.Println(c,c[0],c[1], len(c), cap(c))

	var d = [...] int{1:200,2:300}
	fmt.Println(d)

	var e = [...] string {0:"",1:"hello"} //如果有元素为指定初始值，那么用String类型的默认值，空串，在打印时显示为空格时因为在fmt.println（数组）时，每个元素中间增加一个空格
	fmt.Println(e, cap(e))

	f := [...] struct{
		name string
		age int
	}{
		{"张三丰",120},
		{"周芷若",100},
	}

	fmt.Println(f,f[0].name)

}


func printArray(p *[5]int){
	p[0] = 100;
	for i,v := range p{				//i index    v value
		fmt.Println(i,v)
	}
}

func callPrintArray(){
	printArray(&[5] int {1,2,3,4,5})
}

func sumArray(arr *[5]int){
	sum :=0
	for _,v := range arr{
		sum +=v;
	}
	fmt.Println(sum)
}

func callSumArray(){
	sumArray(&[5] int {1,2,3,4,5})
}

func soo()  {
	var s1 [] int				// 这样只声明，未赋值，未初始化

	if s1 == nil{
		fmt.Println("s1是nil")
	}else{
		fmt.Println("s1不是nil")
	}

	s2 := [] int{}			//未赋值，初始化为nil
	if s2 == nil{
		fmt.Println("s2是nil")
	}else{
		fmt.Println("s2不是nil")
		fmt.Println(s2)
	}

	var s3 = make([] int,0)		//不知道啥意思
	fmt.Println(s3==nil,s3, len(s3), cap(s3))

	var s4 = make([] int,0,0)								// 后边的00啥意思啊
	fmt.Println(s4==nil,s4, len(s4), cap(s4))

	s5 := [] int {1,2,3,4,5,6}			//这是个slice，后边的中括号中如果是...那么是个数组
	fmt.Println(s5)

	arr6 := [...] int {1,2,3,4,5,6}
	s6 := arr6[1:4]				//前包后不包 [1:4)
	fmt.Println(s6)

	fmt.Println(arr6[1:])
	fmt.Println(arr6[:1])

	var s7 = make([] int,0,0)
	fmt.Println(s7)

	var s8 = make([] int,2,2)
	fmt.Println(s8)
	s8[0] = 0
	s8[1]=1
	fmt.Println(s8)


	var s9 = [] int {1,2,3}
	p := &s9[0]
	*p += 100
	fmt.Println(s9)


	var s10 = []int{1,2,3}
	fmt.Printf("ssss %v\n",s10)

	var s11 = []int {4,5,6}
	fmt.Printf("ssss %v\n",s11)

	var s12 = append(s10,s11...)
	fmt.Println(s12)

	var s13 = append(s12,100,200)
	fmt.Println(s13)


	var s14 = make([] int,2)
	fmt.Println(s14, len(s14), cap(s14))

	s14 = append(s14, 100)
	fmt.Println("s14")
	fmt.Println(s14, len(s14), cap(s14))

	s15 := "Hello World"[1:5]
	fmt.Println(s15)

	str16 := "你好，世界"
	s16 := [] rune(str16)
	fmt.Println(string(s16[0:len(s16)-1]))
}

func zoo(){

	v1 := [] int{1,2,3,4,5}
	fmt.Printf("v1 type is %T\n",v1)

	v2 := v1[0:2]
	fmt.Printf("v2 type is %T\n",v2)

	v3 := make([] int,0)
	v3 = append(v3, 1)
	v3 = append(v3, 2)
	fmt.Printf("v3 type is %T\n",v3)

}
