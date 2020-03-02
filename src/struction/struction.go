package struction

import (
	"encoding/json"
	"fmt"
)

type MyInt1 int						//这是一个新的类型，虽然跟int一毛一样，新的类型是struction.MyInt1
type MyInt2 = int					//这不是一个新类型，这个MyInt2也是int，只是个别名

func foo()  {

	var a1 MyInt1 = 100
	var a2 MyInt2 = 100;

	fmt.Printf("type of a1 : %T, type of a2 %T\n",a1,a2)
}

func (i MyInt1) toDouble() MyInt1{
	return i * 2
}

func (i *MyInt1) toDouble2(){
	*i = *i * 2
}

type Family struct {
	dad *Person
	mom *Person
}

func (f *Family) oneYearLater(){
	f.dad.age+=1
	f.mom.age +=1
}


/**
结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
*/
type Person struct{
	name string						//也可以不写name，止血string，但是string类型只能有一个（默认用类型名作为属性名）
	city string
	age int
}

func foo2(){
	var js Person //这样已经初始化
	fmt.Println(js)
	js.name = "金松"
	js.city = "北京"
	js.age = 34

	fmt.Println(js)


	//匿名结构体
	var s2 struct{name string;age int}
	s2.name = "zhangsan"
	s2.age = 18
	fmt.Println(s2)

	var p2 = new(Person)
	fmt.Printf("%T",p2)
	fmt.Println(p2,*p2)

	*p2 = js
	p2.age = 20
	fmt.Println(p2,*p2,p2.name,p2.age)


	var p3 = &js
	p3.name = "haha"
	fmt.Println(p3,js)



	p6 := Person{
		name:"zhangsanfeng",
		city :"wudangshan",
		age : 108,
	}

	fmt.Println(&p6)


	p7 := &Person{"zhouzhiruo","emei",100}
	fmt.Println(p7)


	m1 := make(map[string] *Person)

	s1 := []Person{{name: "zhangsanfeng",age:108},{name: "zhouzhiruo",age:100}}

	fmt.Printf("ssss %p,%p\n",&s1[0],&s1[1])


	for _,p := range s1{
		t := p
		fmt.Printf("zzz%p\n",&t)
		m1[p.name] = &t
	}
	fmt.Println(m1)

	for k,v :=range m1{
		fmt.Printf("%s->%#v->%p\n",k,v,&v)
	}

	m1["zhangsanfeng"].age = 1000
	fmt.Printf("%#v\n",m1["zhangsanfeng"])
	fmt.Println(s1)


	m2 := make(map[string] *Person)
	for i:=0;i< len(s1);i++{
		m2[s1[i].name]=&s1[i]
	}

	fmt.Println(m2)
	/**
	index,value := range slice{},range循环遍历时，这个value时selice中每一个元素的拷贝（值内容一样，值地址不一样），在循环中，多次使用&value，这将得到同一个地址（麻辣个蛋的，拷贝的时候，还往同一个地方拷贝）
	在循环体重增加一个临时局部变量，也不能完全解决问题，比如 t :=p,以后在使用 &t,原因一样，多次循环的p的地址一样，虽然每次分配t的地址不一样，那么多次&t自然得到不同的结果，但是&t的地址已经不是原来slice的地址了。改成for i:=0;i<len;i++即可
	 */




	m1["zhangsanfeng"].foo3()

	m1["zhangsanfeng"].setAge(1)
	m1["zhangsanfeng"].foo3()
	m1["zhangsanfeng"].setAge2(2)
	m1["zhangsanfeng"].foo3()

	var i10 MyInt1 = 10
	i10 = i10.toDouble()
	fmt.Println(i10)

	/**
	这个方法很奇怪，这个方法的定义是这样的
	func (i *MyInt1) toDouble2(){
		*i = *i * 2
	}
	这个方法接收者是个指针，但是这里调用的时候，用i10也可以？不用&i10？
	麻辣个蛋的，原来struct的方法，也可以这么搞！定义的时候是指针，调用的时候用对象。
	。。。
	emmmm
	对，也不奇怪，这是一个方法，方法一定属于某一个类型，如果取了地址，就不是这个类型了，就是 *Type了。对对。
	 */
	i10.toDouble2()
	fmt.Println(i10)


	dad := Person{"zhangcuishan","beijing",25}
	mom := Person{"yinsusu","shanghai",22}

	f := Family{&dad,&mom}
	fmt.Println(f.dad.age,f.mom.age)
	f.oneYearLater()
	fmt.Println(f.dad.age,f.mom.age)
}


/**
这是一个方法，所谓方法，相当于java中的成员方法，一定是某一个Struct的方法。
这个p 是方法接受者。
方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
*/
func (p Person) foo3(){
	fmt.Println("Hi ",p.age)
}

/**
golang中都是值传递，如果这里传入指针，那么将实现类似引用传递的效果。
 */
func (p *Person) setAge(age int){
	p.age = age
}

/**
golang中都是值传递，这里传入Person的值，在这个方法内的修改，不会影响到外边，在后边打印，还是原来的值。
 */
func (p Person) setAge2(age int){
	p.age =age
}




/**
以下类似于继承
 */
type Animal struct {
	name string
	Title string `json:"t"`
}

func (a *Animal) say(){
	fmt.Println("Hi, My name is ",a.name)
}

type Dog struct{
	*Animal
}

func (d *Dog) say(){
	//d.Animal.say()		//这样可以实现类似于调用父类的功能
	//fmt.Println("wang wagn wang ",d.name)					//这样就是重载了啊～～～
}

func foo3(){
	dog := Dog{&Animal{"狗东西","首字母是public，小写是private，小写不能序列化"}}
	dog.say()


	data,error := json.Marshal(dog)
	if error != nil{
		fmt.Println(error)
	}
	fmt.Printf("json%s\n",data)


	str := `{"t":"首字母是public，小写是private，小写不能序列化"}`

	p := &Animal{}
	err := json.Unmarshal([]byte(str),p)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(*p)


	s1 := []User{
		User{"zhangsan",3},
		User{"lisi",4},
	}

	demo(s1)
	fmt.Println(s1)

}

type User struct {
	name string
	age int
}

func demo(u [] User){
	u[0].age = 100
}