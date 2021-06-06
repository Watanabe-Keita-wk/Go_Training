package main

import "fmt"

/*
＜自分解釈＞
structは同じフィールドを束ね、
interfaceはメソッドを束ねて管理する。
*/
type Greeter interface {
	Hello()
}

func hello(greeter Greeter) {
	greeter.Hello()
}

type Human struct {
	Name string
}
func (human Human) Hello() {
	fmt.Printf("私は %s です\n", human.Name)
}

type Cat struct{}
func (cat Cat) Hello() {
	fmt.Println("にゃー")
}


func main() {
	var greeter Greeter = Human{"Keita"}
	hello(greeter)

	greeter = Cat{}
	hello(greeter)
}
/*
私は Keita です
にゃー
*/