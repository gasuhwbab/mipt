package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Next *Node
	Val  int
}

func NewNode(n int) *Node {
	return &Node{
		Next: nil,
		Val:  n,
	}
}

type Stack struct {
	Head *Node
	size int
}

func NewStack() *Stack {
	return &Stack{
		Head: nil,
		size: 0,
	}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(n int) string {
	if s.size == 0 {
		s.Head = NewNode(n)
		s.size++
		return "ok"
	}
	temp := s.Head
	s.Head = NewNode(n)
	s.Head.Next = temp
	s.size++
	return "ok"
}

func (s *Stack) Pop() (int, string) {
	if s.size == 0 {
		return 0, "error"
	}
	ans := s.Head.Val
	s.Head = s.Head.Next
	s.size--
	return ans, ""
}

func (s *Stack) Back() (int, string) {
	if s.size == 0 {
		return 0, "error"
	}
	return s.Head.Val, ""
}

func (s *Stack) Clear() string {
	s.Head = nil
	s.size = 0
	return "ok"
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	s := NewStack()
	var request string
	fmt.Fscan(in, &request)
	for request != "" {
		switch request {
		case "push":
			var n int
			fmt.Fscan(in, &n)
			fmt.Fprintln(out, s.Push(n))
		case "size":
			fmt.Fprintln(out, s.Size())
		case "pop":
			ans, err := s.Pop()
			if err != "" {
				fmt.Fprintln(out, err)
			} else {
				fmt.Fprintln(out, ans)
			}
		case "back":
			ans, err := s.Back()
			if err != "" {
				fmt.Fprintln(out, err)
			} else {
				fmt.Fprintln(out, ans)
			}
		case "clear":
			fmt.Fprintln(out, s.Clear())
		case "exit":
			fmt.Fprintln(out, "bye")
			return
		}
		request = ""
		fmt.Fscan(in, &request)
	}
}
