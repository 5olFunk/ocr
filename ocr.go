package main

import (
	"fmt"
	"gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("~/Documents/ocr/test.jpg")


	fmt.Println("about to call SetRectangle()")
	client.SetRectangle(100,100,20,20)
	text, _ := client.Text()
	fmt.Println(text)

	a := box{topLeft: point{1,1}}
	fmt.Println(a)
	
}

type box struct {
	topLeft point
	topRight point
	bottomLeft point
	bottomRight point
}

type point struct {
	x int
	y int
}
