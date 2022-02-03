package main

import (
	"errors"
	"fmt"
	"os"
)

//结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array   [4]int
	head    int
	tail    int
}

func (q *CircleQueue) Push(val int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.array[q.tail] = val
	q.tail = (q.tail + 1) % q.maxSize
	return nil
}

func (q *CircleQueue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val := q.array[q.head]
	q.head = (q.head + 1) & q.maxSize
	return val, nil
}

func (q *CircleQueue) ListQueue() {
	fmt.Println("环形队列如下")
	//取出当前队列有多少个元素
	size := q.maxSize
	if size == 0 {
		fmt.Println("队列为空")
	}
	// 设计一个辅助变量指向head
	tempHead := q.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, q.array[tempHead])
		tempHead = (tempHead + 1) % q.maxSize
	}
	fmt.Println()
}

//判断环形队列满了
func (q *CircleQueue) IsFull() bool {
	return (q.tail+1)%q.maxSize == q.head
}

//判断是否为空
func (q *CircleQueue) IsEmpty() bool {
	return q.tail == q.head
}

//取总元素
func (q *CircleQueue) GetSize() int {
	// 关键算法
	return (q.tail + q.maxSize - q.head) % q.maxSize
}
func main() {
	circleQueue := &CircleQueue{
		maxSize: 4,
		array:   [4]int{},
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.add 添加数据")
		fmt.Println("2 get 从队列获取数据")
		fmt.Println("3 show 显示数据")
		fmt.Println("4 exit 退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入int入列")
			fmt.Scanln(&val)
			err := circleQueue.Push(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := circleQueue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("val is ", val)
			}
		case "show":
			circleQueue.ListQueue()
		case "exit":
			os.Exit(0)

		}
	}
}
