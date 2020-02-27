package helloworld

import "testing"

func TestAdd(t *testing.T)  {
	r := add(2,4);
	if r != 6{
		t.Fatalf("不对呀,expect:%d,actual:%d",6,r);
	}
	t.Logf("helloworld.add success")
}