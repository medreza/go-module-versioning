package hello

import (
	"fmt"
	"strings"
)

type Hello struct {
	Name string
}

func (h *Hello) Shout() {
	fmt.Println("HELLO, ", strings.ToUpper(h.Name))
}

func (h *Hello) ShoutLouder() {
	fmt.Println("HELLO, ", strings.ToUpper(h.Name), "!!!")
}
