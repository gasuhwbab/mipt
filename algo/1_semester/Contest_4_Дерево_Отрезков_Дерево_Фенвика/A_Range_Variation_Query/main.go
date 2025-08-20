// В начальный момент времени последовательность 𝑎𝑛 задана следующей формулой: 𝑎𝑛=(𝑛2mod12345)+(𝑛3mod23456).
// Требуется много раз отвечать на запросы следующего вида:

// Найти разность между максимальным и минимальным значениями среди элементов 𝑎𝑖,𝑎𝑖+1,…,𝑎𝑗.
// Присвоить элементу 𝑎𝑖 значение 𝑗.

// Входные данные
// Первая строка входного файла содержит натуральное число 𝑘 — количество запросов (1≤𝑘≤100000).
// Следующие 𝑘 строк содержат запросы, по одному на строке. Запрос номер 𝑖 описывается двумя целыми числами 𝑥𝑖, 𝑦𝑖.

// Если 𝑥𝑖>0, то требуется найти разность между максимальным и минимальным значениями среди элементов 𝑎𝑥𝑖,…,𝑎𝑦𝑖.
// При этом 1≤𝑥𝑖≤𝑦𝑖≤100000.

// Если 𝑥𝑖<0, то требуется присвоить элементу 𝑎|𝑥𝑖| значение 𝑦𝑖. В этом случае −100000≤𝑥𝑖≤−1 и |𝑦𝑖|≤100000.

// Выходные данные
// Для каждого запроса первого типа в выходной файл требуется вывести одну строку, содержащую разность между
// максимальным и минимальным значениями на соответствующем отрезке.
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 100000
	INF = int(1<<60) - 1
)

type Node struct {
	isFull      bool
	minV, maxV  int
	left, right *Node
}

func initVal(i int) int {
	x := int64(i)
	return int((x*x)%12345 + (x*x*x)%23456)
}

func ensure(n **Node, tl, tr int) {
	if *n == nil {
		*n = &Node{}
	}
	if (*n).isFull {
		return
	}
	if tl == tr {
		val := initVal(tl)
		(*n).maxV, (*n).minV = val, val
		(*n).isFull = true
		return
	}
	tm := (tl + tr) / 2
	ensure(&(*n).left, tl, tm)
	ensure(&(*n).right, tm+1, tr)
	(*n).minV = min((*n).left.minV, (*n).right.minV)
	(*n).maxV = max((*n).left.maxV, (*n).right.maxV)
	(*n).isFull = true
}

func update(n **Node, tl, tr, pos, val int) {
	if *n == nil {
		*n = &Node{}
	}
	if tl == tr {
		(*n).minV, (*n).maxV = val, val
		(*n).isFull = true
		return
	}
	tm := (tl + tr) / 2
	if pos <= tm {
		update(&(*n).left, tl, tm, pos, val)
	} else {
		update(&(*n).right, tm+1, tr, pos, val)
	}
	(*n).isFull = false
}

func query(n **Node, tl, tr, l, r int) (int, int) {
	if l > r {
		return -INF, INF
	}
	if l == tl && r == tr {
		ensure(n, tl, tr)
		return (*n).maxV, (*n).minV
	}
	if *n == nil {
		*n = &Node{}
	}
	tm := (tl + tr) / 2
	lmn, rmn := INF, INF
	lmx, rmx := -INF, -INF
	lmx, lmn = query(&(*n).left, tl, tm, l, min(tm, r))
	rmx, rmn = query(&(*n).right, tm+1, tr, max(tm+1, l), r)
	return max(lmx, rmx), min(lmn, rmn)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k int
	fmt.Fscan(in, &k)
	root := &Node{maxV: -INF, minV: INF, left: nil, right: nil}
	for ; k > 0; k-- {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if x > 0 {
			mx, mn := query(&root, 1, N, x, y)
			fmt.Fprintln(out, mx-mn)
		} else {
			update(&root, 1, N, -x, y)
		}
	}
}
