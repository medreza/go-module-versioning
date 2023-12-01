package hello

import "fmt"

type Hello struct {
	Name string
}

func (h *Hello) Say() {
	fmt.Println("Hello, ", h.Name)
}
