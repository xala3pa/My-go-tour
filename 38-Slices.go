package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    content := make([][]uint8, dy)
    for x,_ := range content {
		content[x] = make([]uint8, dx)
        for y,_ := range content[x] {
        	content[x][y] = uint8(x^y)
        }
    }
    return content
}

func main() {
    pic.Show(Pic)
}