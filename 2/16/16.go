package main

import "fmt"

func main() {
	arr := []int{23, 42, 11, 32, 9, 55}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		index := excute(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}
}

func excute(arr []int, left, right int) int {
	point := left      //设置基准值
	index := point + 1 //当前比较值
	for i := index; i <= right; i++ {
		//如果当前值小于基准值 则将值写入index列
		if arr[i] < arr[point] {
			//交换
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}

	arr[point], arr[index-1] = arr[index-1], arr[point]
	return index - 1
}
