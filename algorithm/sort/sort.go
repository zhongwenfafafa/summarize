package sort

// 冒泡排序
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		isMove := false
		for j := 0; j < n - 1; j++ {
			if a[j] < a[j + 1] {
				a[j + 1], a[j] = a[j], a[j + 1]
				isMove = true
			}
		}

		if !isMove {
			return
		}
	}
}

// 插入排序
func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}

	//{1,4,6,3,5,9,7}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < a[j] {
				a[j + 1] = a[j]
			} else {
				break
			}
		}

		a[j + 1] = value
	}
}

// 选择排序
func SelectSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		a[i], a[minIndex] = a[minIndex], a[i]
	}
}