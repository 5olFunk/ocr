package main

import (
	"fmt"
	g "gosseract"
)

func main() {
	client := g.NewClient()
	defer client.Close()
	client.SetImage("~/Documents/go/src/ocr/test.jpg")

	rects := []g.Rectangle{
		g.Rectangle{0, 0, 650, 100},
		g.Rectangle{0, 100, 650, 100},
		g.Rectangle{183, 8, 73, 38},
	}
	client.SetRectangles(rects)
	regions, _ := client.TextRegions()
	for _, region := range regions {
		fmt.Println("[%d, %d, %d, %d] = \"%s\"",
			region.Region.Left,
			region.Region.Top,
		)
	}
	fmt.Println(regions[0].Text)

}

type box struct {
	topLeft     point
	topRight    point
	bottomLeft  point
	bottomRight point
}

type point struct {
	x int
	y int
}
