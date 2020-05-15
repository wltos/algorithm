package sort

import "fmt"

// 参考资料
// https://www.cs.usfca.edu/~galles/visualization/ComparisonSort.html   // 教学网站
// TBS learning   // 框框图教学视频

// BubbleSort 1、冒泡排序
// 比较N X M次，内循环，外循环
func BubbleSort(array []int) {
	var n, count = len(array), 0
	for i := 0; i < n-1; i++ {
		var flag bool
		for j := 0; j < n-1-i; j++ {
			fmt.Printf("--%3d--\n", count)
			count = count + 1

			if array[j+1] < array[j] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp

				flag = true
			}
		}

		if !flag { // 优化效率
			return
		}
	}
}

// QuickSort 2、快速排序
// 比较基准值，比基准值小的放到左边，比基准值大的放到右边
// 方框演示思路：[1] 8 5 0 6 2, 两两交换，最后ij重合和基准交换，并返回ij重合的索引
// a. pivot
// b. 比基准小的
// c. 比基准大的
func QuickSort(array []int, left, right int) {
	if left < right {
		mid := getMid(array, left, right)
		QuickSort(array, left, mid-1)
		QuickSort(array, mid+1, right)
	}
}

func getMid(array []int, left, right int) int {
	pivot := array[left]

	for left < right {
		// 右
		for array[right] >= pivot && left < right {
			right = right - 1
		}
		array[left] = array[right]

		// 左
		for array[left] <= pivot && left < right {
			left = left + 1
		}
		array[right] = array[left]
	}

	array[left] = pivot
	return left
}

// SelectSort 3、选择排序
// 每次换最小值得索引，最后跟第一个交换数据
func SelectSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}

		a[i], a[min] = a[min], a[i]
	}
}

// InsertSort 4、插入排序
// 往前移，移不动了就替换
func InsertSort(a []int) {
	for i := range a {
		preIndex := i - 1
		current := a[i]
		for preIndex >= 0 && a[preIndex] > current {
			a[preIndex+1] = a[preIndex]
			preIndex = preIndex - 1
		}
		a[preIndex+1] = current
	}
}

// 合并排序
// 希尔排序
