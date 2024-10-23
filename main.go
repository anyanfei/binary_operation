package main

import "fmt"

const (
	F1 = 1 << iota
	F2
	F3
	F4
	F5
)

// 二进制操作

func main() {
	/*lastSlice := make([]int, 0, 5)
	t1 := 8405008
	binaryStr := fmt.Sprintf("%b", t1)
	fmt.Println(binaryStr)
	allLen := len(binaryStr)
	for i, v := range binaryStr {
		if v == 49 {
			lastSlice = append(lastSlice, int(math.Pow(2, float64(allLen-i-1))))
		}
	}
	fmt.Println(lastSlice)*/
	t1 := F1 | F3 | F5 // 相加
	findBitWeights(t1)
	fmt.Println(isInNum(t1, F2))
	fmt.Println(isInNum(t1, F5))
	fmt.Println(countBits(t1))
}

// 获取传入数字是否在里面
func isInNum(num int, oneNum int) bool {
	return num&oneNum > 0
}

// 获取数位权切片
func findBitWeights(num int) {
	resultSlice := make([]int, 0, countBits(num))
	bitPos := 0
	for num > 0 {
		if num&1 == 1 {
			bitWeight := 1 << bitPos
			resultSlice = append(resultSlice, bitWeight)
		}
		num >>= 1 // 取一次右移一位
		bitPos++
	}
	fmt.Println("最后得到的切片", resultSlice)
}

func countBits(num int) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}
