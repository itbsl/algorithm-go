package sort
//简单选择排序的基本思想:
//第一趟在n个记录中选取最小(大)记录作为有序序列的第一个记录，第二趟在n-1个记录中选取
//最小(大)记录作为有序序列的第二个记录，第i趟在n-i+1个记录中选取最小的记录作为有序序列
//中的第i个记录

//简单选择排序
func SelectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := i+1; j < length; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
