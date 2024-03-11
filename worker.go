package main

import (
	"fmt"
	"time"
)

type Worker struct {
	Name string
}

func NewWorker(name string) *Worker {
	return &Worker{Name: name}
}

// Work 處理肉品
func (w Worker) Work(meat Meat) {
	fmt.Printf("%s 在 %s 取得 %s\n", w.Name, time.Now().Format(time.DateTime), meat.Name())
	meat.Process()
	fmt.Printf("%s 在 %s 處理完 %s\n", w.Name, time.Now().Format(time.DateTime), meat.Name())
}
