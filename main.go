package main

import (
	"fmt"
	"github.com/Violet-Lemon/hexlet-go/v2/greeting"
	"github.com/fatih/color"
)

func main() {
	color.Cyan(greeting.Hello())
	color.Blue("I'm coloured text.")
	color.Magenta("Good luck, dear friend!")
	fmt.Println(greeting.Hello())
}
