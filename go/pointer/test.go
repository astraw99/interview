package main

import (
	"encoding/json"
	"fmt"
)

type Girl struct {
	Name       string `json:"name"`
	DressColor string `json:"dress_color"`
}

func (g *Girl) SetColor(color string) {
	fmt.Printf("g1: %p\n", &g) // 取指针的地址
	g.DressColor = color
}
func (g *Girl) GetJson() string {
	data, _ := json.Marshal(&g)
	return string(data)
}
func main() {
	g := &Girl{Name: "yueyue"}
	fmt.Printf("g0: %p\n", &g) // 取指针的地址

	g.SetColor("white")
	fmt.Println(g.GetJson())
}
