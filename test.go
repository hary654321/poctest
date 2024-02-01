package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p.Name)
}

func main() {
	p := &People{Name: "aaa"}
	p.String()
}
