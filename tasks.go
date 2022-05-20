package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/go-toast/toast"
	"golang.org/x/sys/windows"
)

func NewDrinkTask(d time.Duration) {
	NewTask("drink", d, func() {
		if !mChecked.Checked() {
			return
		}
		message := fmt.Sprintf("该喝水了\n%+v", time.Now().Format("3:04:05 PM"))
		if mNotify.Checked() {
			notification := toast.Notification{
				AppID:   "Petrichor.Go.Reminder",
				Title:   "提醒",
				Message: message,
				Icon:    GetIconPath(),
				Audio:   toast.Reminder,
				Actions: []toast.Action{
					{Type: "protocol", Label: "继续", Arguments: ""},
					{Type: "protocol", Label: "关闭提醒", Arguments: ""},
				},
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
				abort("go-toast notify", err)
			}
		} else {
			var hwnd windows.HWND = 0
			ret, err := windows.MessageBox(
				hwnd,
				syscall.StringToUTF16Ptr(message),
				syscall.StringToUTF16Ptr("提醒"),
				windows.MB_OKCANCEL|windows.MB_ICONINFORMATION,
			)
			if err != nil {
				abort("windows.MessageBox", err)
			}
			switch ret {
			case 1:
				taskList["drink"].Reset()
				log.Println("[drink] ok, continue")
			case 2:
				log.Println("[drink] cancel")
			}
		}
	})
}

func StopDrinkTask() {
	taskList["drink"].Stop()
}

func ResetDrinkTask() {
	taskList["drink"].Reset()
}

func ChangeDrinkTask(d time.Duration) {
	taskList["drink"].Change(d)
}
