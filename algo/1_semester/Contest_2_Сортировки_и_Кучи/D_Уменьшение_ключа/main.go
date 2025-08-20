package main

import (
	"bufio"
	"os"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner() *Scanner {
	return &Scanner{r: bufio.NewReaderSize(os.Stdin, 1<<20)}
}

func (s *Scanner) next() string {
	b := make([]byte, 0, 16)
	for {
		c, err := s.r.ReadByte()
		if err != nil {
			return string(b)
		}
		if c > ' ' {
			b = append(b, c)
			break
		}
	}
	for {
		c, err := s.r.ReadByte()
		if err != nil || c <= ' ' {
			break
		}
		b = append(b, c)
	}
	return string(b)
}

func (s *Scanner) nextInt() int {
	num := s.next()
	sign, val, i := 1, 0, 0
	if len(num) > 0 && num[0] == '-' {
		sign = -1
		i = 1
	}
	for ; i < len(num); i++ {
		val = val*10 + int(num[i]-'0')
	}
	return val * sign
}

func (s *Scanner) nextInt64() int64 {
	num := s.next()
	var sign int64 = 1
	i := 0
	var val int64 = 0
	if len(num) > 0 && num[0] == '-' {
		sign = -1
		i = 1
	}
	for ; i < len(num); i++ {
		val = val*10 + int64(num[i]-'0')
	}
	return val * sign
}

func intToString[T int | int8 | int16 | int32 | int64](x T) string {
	if x == 0 {
		return "0\n"
	}
	sign := ""
	if x < 0 {
		sign = "-"
		x = -x
	}
	buf := [20]byte{}
	i := len(buf)
	for x > 0 {
		i--
		buf[i] = byte('0' + x%10)
		x /= 10
	}
	return sign + string(buf[i:]) + "\n"
}

func main() {
	in := NewScanner()
	out := *bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	q := in.nextInt()
	val := make([]int64, q+1) // значения
	heap := make([]int, q+1)  //
	pos := make([]int, q+1)   // позиция i-го щзапроса в куче
	size := 0
	for qi := 1; qi < q+1; qi++ {
		req := in.next()
		switch req {
		case "insert":
			x := in.nextInt64()
			val[qi] = x
			size++
			heap[size] = qi
			pos[qi] = size
			siftUp(val, pos, heap, size)
		case "getMin":
			out.WriteString(intToString(val[heap[1]]))
		case "extractMin":
			i := heap[1]
			heap[1], heap[size] = heap[size], heap[1]
			pos[heap[1]] = 1
			pos[i] = 0
			size--
			if size > 0 {
				siftdown(val, pos, heap, 1, size)
			}
		case "decreaseKey":
			i, diff := in.nextInt(), in.nextInt64()
			val[i] -= diff
			siftUp(val, pos, heap, pos[i])
		}
	}
}

func siftUp(val []int64, pos, heap []int, i int) {
	for i > 1 {
		p := i / 2
		if val[heap[i]] >= val[heap[p]] {
			break
		}
		heap[i], heap[p] = heap[p], heap[i]
		pos[heap[i]], pos[heap[p]] = i, p
		i = p
	}
}

func siftdown(val []int64, pos, heap []int, i, size int) {
	for {
		l, r := 2*i, 2*i+1
		if l > size {
			break
		}
		j := l
		if r <= size && val[heap[r]] < val[heap[l]] {
			j = r
		}
		if val[heap[i]] < val[heap[j]] {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		pos[heap[i]] = i
		pos[heap[j]] = j
		i = j
	}
}
