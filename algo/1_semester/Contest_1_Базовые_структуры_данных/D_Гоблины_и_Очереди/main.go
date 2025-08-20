package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	next *Node
	val  int
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

type LinkedList struct {
	head *Node
	tail *Node
	mid  *Node
	size int
}

func (l *LinkedList) InsertPrivilegedGoblin(val int) {
	g := NewNode(val)
	if l.size == 0 {
		l.head = g
		l.mid = g
		l.tail = g
		l.size++
		return
	}
	if l.size == 1 {
		l.tail.next = g
		l.tail = l.tail.next
		l.size++
		return
	}
	g.next = l.mid.next
	l.mid.next = g
	l.size++
	if l.size%2 != 0 {
		l.mid = l.mid.next
	}
}

func (l *LinkedList) InsertGoblin(val int) {
	g := NewNode(val)
	if l.size == 0 {
		l.head = g
		l.mid = g
		l.tail = g
		l.size++
		return
	}
	l.tail.next = g
	l.tail = l.tail.next
	l.size++
	if l.size%2 != 0 {
		l.mid = l.mid.next
	}
}

func (l *LinkedList) DeleteGoblin() int {
	ans := l.head.val
	l.head = l.head.next
	l.size--
	if l.size%2 != 0 {
		l.mid = l.mid.next
	}
	return ans
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	l := &LinkedList{}
	for ; N > 0; N-- {
		var request string
		fmt.Fscan(in, &request)
		switch request {
		case "+":
			var i int
			fmt.Fscan(in, &i)
			l.InsertGoblin(i)
		case "-":
			fmt.Fprintln(out, l.DeleteGoblin())
		case "*":
			var i int
			fmt.Fscan(in, &i)
			l.InsertPrivilegedGoblin(i)
		}
	}
}
