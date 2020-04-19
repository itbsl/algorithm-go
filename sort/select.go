package sort
//简单选择排序的基本思想:
//第一趟在n个记录中选取最小(大)记录作为有序序列的第一个记录，第二趟在n-1个记录中选取
//最小(大)记录作为有序序列的第二个记录，第i趟在n-i+1个记录中选取最小的记录作为有序序列
//中的第i个记录

//在逆序情况下
//元素的比较次数: n(n-1)/2
//元素的移动次数为: 3(n-1)

//在正序情况下
//元素的比较次数: n(n-1)
//元素的移动次数为: 0

//简单选择排序
func SelectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i+1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
