package main

import (
	"time"
)

var (
	taskList = map[string]*Task{}
	// drinkTask *time.Timer
)

type TaskFunc func()

type Task struct {
	Name     string
	Duration time.Duration
	Timer    *time.Timer
}

func NewTask(name string, d time.Duration, fn TaskFunc) {
	taskList[name] = &Task{
		Name:     name,
		Duration: d,
		Timer:    time.AfterFunc(d, fn),
	}
}

// 返回一个 bool 值
// 若调用停止了计时器，则返回 true
// 若计时器已经过期或停止，则返回 false
// 停止不关闭通道，以防止读取通道错误的成功
func (t Task) Stop() {
	t.Timer.Stop()
}

func (t Task) Reset() {
	t.Timer.Reset(t.Duration)
}

func (t Task) Change(d time.Duration) {
	t.Timer.Reset(d)
}

func (t Task) Notify() {

}

func (t Task) Message() {

}
