package main

import (
	"fmt"
	"github.com/Violet-Lemon/hexlet-go/greeting"
	"github.com/fatih/color"
)

func main() {
	color.Cyan(greeting.Hello())
	color.Blue("I'm coloured text.")
	color.Magenta("Good luck!")
	fmt.Println(greeting.Hello())
}
