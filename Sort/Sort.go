package Sort

import "math/rand"


/*****************************************************************************************
	排序算法
		1.冒泡排序
		2.插入排序
		3.快速排序



******************************************************************************************/



// 冒泡排序
func BubbleSort(arr []int) []int {
	tmp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

//插入排序
func InsertSort(arr []int) []int {
	var i, j, tmp int
	for i = 1; i < len(arr); i++ {
		tmp = arr[i]
		for j = i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}

// 快速排序
func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = QuickSort(low_part)
	high_part = QuickSort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}

// 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return Merge(left, right)
}
func Merge(left, right []int) []int {
	result := make([]int, 0, len(left) + len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}


// 堆排序(Heap Sort)
func HeapSort(arr []int) []int{
	i := 0
	tmp := 0
	for i = len(arr) / 2 - 1; i >= 0; i-- {
		arr = sift(arr, i, len(arr))
	}

	for i = len(arr) - 1; i >= 1; i-- {
		tmp = arr[0];
		arr[0] = arr[i];
		arr[i] = tmp;
		arr = sift(arr, 0, i);
	}
	return  arr
}

func sift(arr []int, i int, arrLen int) []int {
	done := false

	tmp       := 0
	maxChild  := 0

	for (i * 2 + 1 < arrLen) && (!done) {
		if i * 2 + 1 == arrLen - 1 {
			maxChild = i * 2 + 1
		} else if arr[i * 2 + 1] > arr[i * 2 + 2] {
			maxChild = i * 2 + 1
		} else {
			maxChild = i * 2 + 2
		}

		if arr[i] < arr[maxChild] {
			tmp = arr[i]
			arr[i] = arr[maxChild]
			arr[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}


