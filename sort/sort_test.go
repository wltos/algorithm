package sort

import "testing"

// 测试指令
// go test                                              // 测试
// go test -v sort.go sort_test.go                      // 测试单个文件
// go test -v -run TestBubbleSort sort_test.go sort.go  // 测试指定单元测试用例

func TestBubbleSort(t *testing.T) {
	a := []int{1, 8, 5, 0, 6, 2, 7}
	BubbleSort(a)
	t.Logf("%v", a)
}

func TestQuickSort(t *testing.T) {
	a := []int{1, 8, 5, 0, 6, 2, 7}
	QuickSort(a, 0, len(a)-1)
	t.Logf("%v", a)
}

func TestSelectSort(t *testing.T) {
	a := []int{1, 8, 5, 0, 6, 2, 7}
	SelectSort(a)
	t.Logf("%v", a)
}

func TestInsertSort(t *testing.T) {
	a := []int{1, 8, 5, 0, 6, 2, 7}
	InsertSort(a)
	t.Logf("%v", a)
}
