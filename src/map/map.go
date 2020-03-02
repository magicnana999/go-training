package _map

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func foo() {

	m1 := make(map[string]int, 2)
	m1["age"] = 100
	m1["grade"] = 200
	m1["foo"] = 300
	m1["f2"] = 400
	m1["ff3"] = 500

	fmt.Println(m1)
	fmt.Printf("type is %T\n", m1)

	value, exist := m1["age22"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("无数据")
	}

	for k1,v1 := range m1{
		fmt.Println(k1,v1)
	}

	delete(m1, "f2")
	for k2,_ := range m1{
		fmt.Println(k2)
	}

	for k3 := range m1{
		fmt.Println(k3)
	}


	rand.Seed(time.Now().UnixNano())
	m2 := make(map[string]int,100)

	for i :=0;i<2;i++{
		key := fmt.Sprintf("stu%02d",i)
		value := rand.Int()
		m2[key] = value
	}
	fmt.Println(m2)


	s1 := make([] string,0, len(m2))
	//上边make中，len=2，才能创建2个元素的内存，后续的cap>len的部分，只能append
	fmt.Println(s1,len(s1), cap(s1))

	for k,_ := range m2{
		s1 = append(s1, k)
	}
	fmt.Println(s1)

	sort.Strings(s1)
	for index,k := range s1{
		fmt.Println(index,k,m2[k])
	}

	s3 := make([]map[string]int,0,5)
	for i := 0;i<5;i++{
		m3 := make(map[string] int,2)
		m3["age"] = 100 * i + 100
		m3["score"] = 200 * i+200
		s3 = append(s3, m3)
	}

	fmt.Println(s3)


	m4 := make(map[string][]string,2)
	s4_1 := make([]string,0,2)
	s4_1 = append(s4_1, "haha","hehe")
	s4_2 := make([]string,0,2)
	s4_2 = append(s4_2,"gaga","yaya")
	m4["age"] = s4_1
	m4["score"] = s4_2
	fmt.Println(m4)
}
