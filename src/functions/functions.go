package functions

import (
	"bufio"
	"fmt"
	"os"
)

func foo() {
	var x = 10
	var y = 20
	z := tt(func(x, y int) int {
		return x + y
	}, x, y)

	fmt.Println(z)

	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)

	anyType(100)

	r1, r2 := return2()
	fmt.Println(r1, r2)

	fmt.Println(隐士返回(1, 2))

	return
}

func foo2() {
	defer1()
	defer2()
	defer3()
}

func foo3() {
	c := a()
	c()
	c()
	c()

	a()
	匿名函数1()
	匿名函数2()

	值传递还是引用传递()
	传递参数int()
	参数指针Struct()
	参数指针map()
	参数指针Array()
	参数指针Slice()
	//传递slice到新的局部作用于导致扩容不会影响原值()

}

func 参数指针Slice(){
	println()
	i := []int{100}
	fmt.Printf("%p %p %v len(i)=%v cap(i)=%v\n",&i,i,i,len(i),cap(i))
	参数指针Slice1(i)
	参数指针Slice2(&i)
	参数指针Slice3(i)
	参数指针Slice4(&i)
	fmt.Printf("%p %p %v len(i)=%v cap(i)=%v\n",&i,i,i,len(i),cap(i))
}

func 参数指针Slice1(i []int){
	i[0] +=1
	fmt.Printf("%p %p %v len(i)=%v cap(i)=%v\n",&i,i,i,len(i),cap(i))
}

func 参数指针Slice2(i *[]int){
	(*i)[0] +=1
	fmt.Printf("%p %p %v len(i)=%v cap(i)=%v\n",&i,i,i,len(*i),cap(*i))
}

func 参数指针Slice3(i []int){
	i = append(i, 100,200)
	fmt.Printf("%p %p %v len(i)=%v cap(i)=%v\n",&i,i,i,len(i),cap(i))
}

func 参数指针Slice4(i *[]int){
	*i = append(*i, 300,400,500,600)
	fmt.Printf("%p %p %v len(i)=%v cap(i)=%v\n",&i,i,i,len(*i),cap(*i))
}

func 参数指针Array(){
	println()
	i := [1]int{100}
	fmt.Printf("%p %p %v\n",&i,&i[0],i)
	参数指针Array1(i)
	参数指针Array2(&i)
	fmt.Printf("%p %p %v\n",&i,&i[0],i)
}

func 参数指针Array1(i [1]int){
	i[0] +=1
	fmt.Printf("%p %p %v\n",&i,&i[0],i)
}

func 参数指针Array2(i *[1]int){
	i[0] +=2
	fmt.Printf("%p %p %p %v\n",&i,i,&i[0],i)
}

//closure 闭包
func a() func() int {
	a := 0
	fmt.Println("outer", a, &a)
	b := func() int {
		a++
		fmt.Println("inner", a, &a)
		return a
	}
	return b
}

func 传递slice到新的局部作用于导致扩容不会影响原值() {
	s := []int{1}
	传递slice到新的局部作用于导致扩容不会影响原值1(s)
	fmt.Println(s)
}

func 传递slice到新的局部作用于导致扩容不会影响原值1(i []int) {
	i[0] = 100
	i = append(i, 200)
	fmt.Println(i)
}

type Person struct {
	name string
	age  int
}

func 参数指针Struct() {
	p := Person{"张三丰", 100}
	fmt.Printf("%p %v\n", &p, p)
	参数指针Struct1(p)
	参数指针Struct2(&p)
	fmt.Printf("%p %v\n", &p, p)
}

func 参数指针Struct1(p Person) {
	p.age += 1
	fmt.Printf("%p %v\n", &p, p)
}

func 参数指针Struct2(p *Person) {
	p.age += 2
	fmt.Printf("%p %p %v\n", &p, p, *p)
}

func 传递参数int() {
	i := 100
	fmt.Printf("%p %v\n", &i, i)
	传递参数int1(i)
	传递参数int2(&i)
	fmt.Printf("%p %v\n", &i, i)

}

func 传递参数int1(i int) {
	i += 1
	fmt.Printf("%p %v\n", &i, i)
}

func 传递参数int2(i *int) {
	*i += 2
	fmt.Printf("%p %p %v\n", &i, i, *i)
}

func 参数指针map() {
	fmt.Println()
	m := make(map[string]string)
	m["name"] = "张三丰"
	m["age"] = "100多了"
	m["wife"] = "没媳妇"
	fmt.Printf("%p %p %v \n", &m, m,m)
	参数指针map1(m)
	参数指针map2(&m)
	fmt.Printf("%p %p %v \n", &m, m,m)

}

func 参数指针map1(m map[string]string) {
	m["age"] = "200多了"
	m["wife"] = m["wife"] +"郭襄"
	fmt.Printf("%p %p %v \n", &m, m,m)
}

func 参数指针map2(m *map[string]string) {
	(*m)["age"] = "300多了"
	(*m)["wife"] = (*m)["wife"]+"李莫愁"
	fmt.Printf("%p %p %v\n", &m, m, *m)

}

func 值传递还是引用传递() {

	type Person struct {
		age int
	}

	i := 100
	a := [2]int{1, 2}
	s := make([]int, 2)
	s[0] = 1
	p := Person{1}

	println(s)
	fmt.Printf("在头部 i[%p->%d],a[%p,%p->%v],s[%p->%p,%p->%v],p[%p->%v]\n", &i, i, &a[0], &a[1], a, &s, &s[0], &s[1], s, &p, p)

	/**
	外部参数 0xc00001c358 -> 100
	分配内存 0xc00001c3a0，把100拷贝到0xc00001c3a0
	*/
	func(i int) {
		i++
		fmt.Printf("传入int类型 [%p->%v]\n", &i, i)
	}(i)

	/**
	外部参数 0xc0000a6268 -> 100
	分配内存 0xc0000a0030，把i的地址（0xc0000a6268）拷贝到0xc0000a0030，but，0xc0000a6268这个地址存储的值是100，值拷贝，拷贝i的地址（也很好理解，你在调用时，显示的指定了i的地址）
	*/
	func(i *int) {
		*i++
		fmt.Printf("传入int指针 [%p->%v->%v]\n", &i, i, *i)
	}(&i)

	/**
	外部参数 0xc00001c360 -> [0xc00001c360,0xc00001c368 -> 1,2]
	分配内存 0xc00001c3d0，0xc00001c3d8，把数组[0]和[1]的值分别复制过来，值拷贝，拷贝数组所有元素的值
	*/
	func(i [2]int) {
		i[0] = 100
		i[1] = 200
		fmt.Printf("传入[2]int类型 [%p->%p,%p->%v]\n", &i, &i[0], &i[1], i)
	}(a)

	/**
	外部参数 0xc0000a6270-> [0xc0000a6270,0xc0000a6278 -> 1,2]
	分配内存 0xc0000a0038，还有一个看不到，把数组的[0xc0000a6270->0]和[0xc0000a6278->1]的地址分别复制过来，值拷贝，拷贝数组所有元素的地址，but，这些所有元素的地址上存储的值还是1，2
	*/
	func(i *[2]int) {
		temp := [2]int{100, 200}
		*i = temp
		fmt.Printf("传入[2]int指针 [%p->%p,%p->%v]\n", &i, &i[0], &i[1], i)
	}(&a)

	/**
	待定
	*/
	func(i []int) {
		println(i)
		i[1] = 200
		i = append(i, 300, 400)
		fmt.Printf("传入 []int类型 [%p->%p,%p->%v]\n", &i, &i[0], &i[1], i)
		println(i)
	}(s)

	//f6 :=func(i *[]int){
	//	*i = append(*i,100)
	//	fmt.Printf("在f6内部 [%p->%#v]\n",&i,i)
	//}
	//f5(s)
	//f6(&s)

	//f7 := func(i Person){
	//	i.age += 100
	//	fmt.Printf("在f7内部 [%p->%#v]\n",&i,i)
	//}

	//f8 := func(i *Person){
	//	i.age += 200
	//	fmt.Printf("在f8内部 %T [%p->%#v]\n",i,&i,i)
	//}

	//f7(p)
	//f8(&p)

	fmt.Printf("在尾部 i[%p->%d],a[%p,%p->%v],s[%p->%p,%p->%v],p[%p->%v]\n", &i, i, &a[0], &a[1], a, &s, &s[0], &s[1], s, &p, p)
	fmt.Println(i, a, s, p)

	//切片扩容原理()
	//切片存储结构()

}

func 切片存储结构() {
	i := []int{1, 2}
	println(i)
	fmt.Println(&i[0], "->", i[0])
	fmt.Println(&i[1], "->", i[1])
}

func 切片扩容原理() {
	ss := []int{5}
	fmt.Println(ss)
	println(ss)
	for index := 0; index < len(ss); index++ {
		print(&ss[index])
	}
	println(" ======= 5 ======")

	ss = append(ss, 7)
	fmt.Println(ss)
	println(ss)
	for index := 0; index < len(ss); index++ {
		print(&ss[index])
	}
	println(" ======= 7 ======")

	ss = append(ss, 9)
	fmt.Println(ss)
	println(ss)
	for index := 0; index < len(ss); index++ {
		print(&ss[index])
	}
	println(" ======= 9 ======")

	ss = append(ss, 11)
	fmt.Println(ss)
	println(ss)
	for index := 0; index < len(ss); index++ {
		print(&ss[index])
	}
	println(" ======= 11 ======")

	ss = append(ss, 12)
	fmt.Println(ss)
	println(ss)
	for index := 0; index < len(ss); index++ {
		print(&ss[index])
	}
	println(" ======= 12 ======")

}

func 匿名函数1() {
	a := 100
	func() {
		a++
		fmt.Println("匿名函数1", a, &a)
	}()
	a++
	fmt.Println("匿名函数1", a, &a)
}

func 匿名函数2() {
	a := 100
	f := func() {
		a++
		fmt.Println("匿名函数2", a, &a)
	}
	a++
	f()
	fmt.Println("匿名函数2", a, &a)
}

/**
 1. 关键字 defer 用于注册延迟调用。
 2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
 3. 多个defer语句，按先进后出的方式执行。
 4. defer语句中的变量，在defer声明时就决定了。


   1. 关闭文件句柄
   2. 锁资源释放
   3. 数据库连接释放

这个函数演示2个东西，defer的参数在定义时固定，多个defer先进后出
*/
func defer1() {
	fmt.Println("===================   defer   =====================")
	n := 0
	defer fmt.Println("0 in func defer n=", n) //1这个defer先进来，在下边的defer之后执行，此行在定义时，n已固定为0，即使在return之前执行时n已经变化，这里仍然是0
	n++
	defer fmt.Println("1 in func defer n=", n) //2这个deffer后进来，在上边的defer之前执行，此行在定义时，n已固定为1，即使在return之前执行时n已经变化，这里仍然时1
	n++
	fmt.Println("2 in func return n=", n) //这两行不是defer，在上边的2个defer之前执行，所以顺序321
	fmt.Println("===================   defer   =====================")

}

func defer2() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func defer3() (e error) {
	fmt.Println("start read")
	filepath := "/Users/jinsong/1.txt"
	f, e := os.Open(filepath)
	if e != nil {
		fmt.Println(e)
		return e
	}

	if f != nil {
		defer func() {
			if e := f.Close(); e != nil {
				fmt.Println(e)
				fmt.Printf("defer close %s err %v\n", filepath, e)
			} else {
				fmt.Println("defer close ok")
			}
		}()
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("end read")

	return e

}

/**
defer 后边的是个匿名函数（是不是叫闭包？还没学到呢），func(){}这个叫定义，后边的()叫调用，所以这是定义一个匿名函数+调用
defer 后边必须是一个函数调用
这个方法演示defer后边可以来个匿名函数
*/
func f1() (r int) {
	fmt.Println("f1 - 0 r=", r)
	defer func() {
		r++
		fmt.Println("f1 inner r=", r)
	}()
	fmt.Println("f1 - 1 r=", r)
	return r
}

/**
1 defer后边可以定义并执行一个匿名函数
2 defer 的参数在定义时固定
3 defer可以在函数执行的最后，更改返回值
* 多个defer先进后出
*/
func f2() (r int) {
	r = 100
	fmt.Println("f2 - 0 r=", r, &r) // 第一输出100,注意r的地址都是同样的，匿名函数内部直接引用外部变量的地址（指针）
	defer func(p int) {
		p += 100
		fmt.Println("f2 inner1 p=", p, &p)
		//第四输出200,为什么这里时200，不是400？因为这里的p作为参数传递进来，golang中的传递都是值传递，这个p从100复制而来，两个不是同一个地址，由fmt可知。
		// 下边的defer不一样，下边不是值传递，下边是直接引用外部变量的地址，所以可以更改。
	}(r)

	defer func() {
		r += 100
		fmt.Println("f2 inner2 r=", r, &r) //第三输出300		这个函数的返回值，在defer中被更改为300
	}()

	r += 100
	fmt.Println("f2 - 1 r=", r, &r) //第二输出200
	return r

}

func f3() (r int) {
	r = 100
	//这样不能编译，因为 返回值指定了r，相当于已经声明了 var r int = 0,或者 r := 0,
	// 下边的 r := 100相当于声明一个变量r，并赋值为100，在同一个作用于内
	fmt.Println("f3 - 0 r=", r, &r)
	return r
}

func tttt() {
	var a int = 100
	//a := 200
	fmt.Println(a)
}

func tttt2() {
	b1 := 100
	//b1 := 200
	fmt.Println(b1)
}

/**
z在返回定义中指定，在局部赋值（不能:=),
如果在返回定义中未指定，则必须显示返回
*/
func 隐士返回(x, y int) (z int) {
	z = x + y
	return
}

func return2() (int, int) {
	return 100, 200
}

func anyType(n interface{}) {
	fmt.Println(n)
}

/**
在默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

注意2：map、slice、chan、指针、interface默认以引用的方式传递。

不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）

Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。


*/
func swap(x, y *int) {

	temp := *x
	*x = *y
	*y = temp

}
func tt(f FuncDouble, x int, y int) int {
	return f(x, y)
}

type FuncDouble func(x, y int) int
