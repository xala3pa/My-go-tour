package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

func WImpl(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WImpl(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WImpl(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WImpl(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("You Win!!!")
	} else {
		fmt.Println("OH!!!!")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("You win!!!")
	} else {
		fmt.Println("OH!!!!!")
	}
}
