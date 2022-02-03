package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 数组=》 模拟队列
	front   int    //标识指向队列首部(不含队首)
	rear    int    // 指向队列尾部(含队尾)
}

// 添加数据到队列
func (q *Queue) AddQueue(val int) error {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 {
		return errors.New("queue full")
	}
	q.rear++ //后移
	q.array[q.rear] = val
	return nil
}

// 显示队列,找到队首，然后遍历到位尾
func (q *Queue) ShowQueue() {
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d \t", i, q.array[i])
	}
	fmt.Println()
}

//从队列取数据
func (q *Queue) GetQueue() (int, error) {
	//判断队列是否为空
	if q.rear == q.front {
		return 0, errors.New("队列为空")
	}
	q.front++
	return q.array[q.front], nil
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1. add 添加数据")
		fmt.Println("2 get 从队列获取数据")
		fmt.Println("3 show 显示数据")
		fmt.Println("4 exit 退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入int入列")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("val is ", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)

		}
	}
}
