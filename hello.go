package hello

import "fmt"

type Hello struct {
	Name string
}

func (h *Hello) SayHello() {
	fmt.Println("Hello, ", h.Name)
}
