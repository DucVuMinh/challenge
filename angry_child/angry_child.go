package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

// Complete the angryChildren function below.
func angryChildren(k int32, packets []int32) int64 {
	// final formular: Min f:= 1->(n-1) { sum(i = 1->(k-1) : [k-i)*i*(a_(j+i+1) - a_(j+i)]  }   )
	// sử dụng dynamic programming tính từng phần tử sum(i = 1->(k-1) : [k*i*(a_(j+i+1) - a_(j+i)] và sum(i = 1->(k-1) : [i*i*(a_(j+i+1) - a_(j+i)]
	packetsSorted := []int{}
	for _,v := range packets {
		packetsSorted = append(packetsSorted, int(v))
	}
	sort.Ints(packetsSorted)
	sub := []int{}
	nk := int(k)
	for i, v := range packetsSorted {
		if i == 0 {
			continue
		}
		sub = append(sub, v-packetsSorted[i-1])
	}
	var left []int64
	var right []int64
	var index int
	var leftItem int64
	var rightItem int64

	for index = 1; index < int(k); index++ {
		leftItem += int64(index * sub[index-1])
		rightItem += int64(index * index) * int64(sub[index-1])
	}
	left = append(left, leftItem)
	right = append(right, rightItem)
	count := 1
	for index = nk; index < (len(packets) -1); index++ {
		leftItem = leftItem + int64((nk-1)*sub[index-1] - (packetsSorted[index-1] - packetsSorted[index-nk]))
		left = append(left, leftItem)
		rightItem = rightItem - 2*left[count-1] +   int64((nk-1)*(nk-1))*int64(sub[index-1] ) + int64((packetsSorted[index-1] - packetsSorted[index-nk]))
		right = append(right, rightItem)
		count++
	}
	min := int64(math.Pow(2, 60))
	for i := 0; i < count; i++ {
		resTemp := int64(nk)*left[i] - right[i]
		min = int64(math.Min(float64(min), float64(resTemp)))
	}
	return min
}

func main() {
	//v := angryChildren(5,[]int32{4504,1520,5857,4094,4157,3902,822,6643,2422,7288,8245,9948,2822,1784,7802,3142,9739,5629,5413,7232})
	//v := angryChildren(5,[]int32{10,10,10,100,300,200,1000,20,30})
	//fmt.Println(v)
	i := 121113346195225384
	fmt.Println(math.Pow(2,57))
	fmt.Println(math.Log2(float64(i)))
	fmt.Println(float64(i))
	//reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	//
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	//writer := bufio.NewWriterSize(stdout, 1024*1024)
	//
	//nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//checkError(err)
	//n := int32(nTemp)
	//
	//kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//checkError(err)
	//k := int32(kTemp)
	//
	//var packets []int32
	//
	//for i := 0; i < int(n); i++ {
	//	packetsItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//	checkError(err)
	//	packetsItem := int32(packetsItemTemp)
	//	packets = append(packets, packetsItem)
	//}
	//
	//result := angryChildren(k, packets)
	//
	//fmt.Fprintf(writer, "%d\n", result)
	//
	//writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
