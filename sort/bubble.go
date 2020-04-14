package sort

//复杂度O(n平方):
// *在最好情况下(正序),元素的交换次数为0，比较次数为n-1
// *在最坏情况下(正序),元素的交换次数为n(n-1)/2,比较次数为n(n-1)/2
//空间复杂度O(1):
// *只需要一个辅助单元进行交换
//使用情况:
// *元素数目少,或者元素的初始化序列基本有序

//冒泡排序
func BubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		//做一个标记,如果在本次排序过程中,没有出现交换的情况,说明排序已经有序,
		//可以直接跳出循环,减少不必要的比较
		mark := true
		for j := 0; j < (length-i-1); j++ {
			if arr[j] > arr[j+1] {
				mark = false //出现了交换情况，则将标记改为false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		//mark值未出现变化，排序已经有序，直接退出循环即可
		if mark {
			return
		}
	}
}