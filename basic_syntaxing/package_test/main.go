package main

import (
	"github.com/joho/godotenv"
	"gopkg.in/gookit/color.v1"
)

func main() {
	godotenv.Load()

	color.Red.Println("This is color red!")
	color.Blue.Println("This is blue!")
}
