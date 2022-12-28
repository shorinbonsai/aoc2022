package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	// "strings"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filey := ReadFile("input.txt")
	var sums []int
	index := 0
	sum := 0
	for i, numb := range filey {
		if i == len(filey)-1 {
			numb2 := ToInt(numb)
			sum += numb2
			sums = append(sums, sum)
			index += 1
		} else if numb == "" {
			sums = append(sums, sum)
			index += 1
			sum = 0
		} else {
			numb2 := ToInt(numb)
			sum += numb2
		}

	}
	max := 0
	for _, i := range sums {
		if i > max {
			max = i
		}
	}

	println(max)
	// sort.Ints(sums)
	// sort.Reverse(sort.IntSlice(sums))
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	newsum := sums[0] + sums[1] + sums[2]
	println(newsum)
	return
}
