//go:generate rsrc -manifest assets/Reminder.manifest -ico assets/favicon.ico -o assets/Reminder.syso

package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

var (
	//go:embed assets/favicon.ico
	Icon []byte

	mReset     *systray.MenuItem
	subMenu    *systray.MenuItem
	subMenu3s  *systray.MenuItem
	subMenu15m *systray.MenuItem
	subMenu30m *systray.MenuItem
	subMenu1h  *systray.MenuItem

	mNotify     *systray.MenuItem
	mMessageBox *systray.MenuItem
	mChecked    *systray.MenuItem
	mQuit       *systray.MenuItem
)

func onReady() {
	systray.SetIcon(Icon)
	systray.SetTitle("Reminder")
	systray.SetTooltip("服务已最小化右下角, 右键点击打开菜单！")
	mReset = systray.AddMenuItem("重新开始", "重新开始计时")
	subMenu = systray.AddMenuItem("选择时间", "选择时间，默认 30m")
	subMenu3s = subMenu.AddSubMenuItemCheckbox("3s", "3s", false)
	subMenu15m = subMenu.AddSubMenuItemCheckbox("15m", "15m", false)
	subMenu30m = subMenu.AddSubMenuItemCheckbox("30m", "30m", true)
	subMenu1h = subMenu.AddSubMenuItemCheckbox("1h", "1h", false)
	systray.AddSeparator()
	mNotify = systray.AddMenuItemCheckbox("使用通知", "使用通知", true)
	mMessageBox = systray.AddMenuItemCheckbox("使用弹窗", "使用弹窗", false)
	systray.AddSeparator()
	mChecked = systray.AddMenuItemCheckbox("打开提醒", "Check Me", true)
	systray.AddSeparator()
	mQuit = systray.AddMenuItem("退出", "退出程序")

	NewDrinkTask(30 * time.Minute)

	go func() {
		for {
			select {
			case <-mReset.ClickedCh:
				fmt.Println("Reset")
				ResetDrinkTask()

			case <-subMenu3s.ClickedCh:
				subMenu.SetTitle("3s")
				systray.SetTooltip("当前计时器: 3s")
				subMenu3s.Check()
				subMenu15m.Uncheck()
				subMenu30m.Uncheck()
				subMenu1h.Uncheck()
				ChangeDrinkTask(3 * time.Second)
			case <-subMenu15m.ClickedCh:
				subMenu.SetTitle("15m")
				systray.SetTooltip("当前计时器: 15m")
				subMenu3s.Uncheck()
				subMenu15m.Check()
				subMenu30m.Uncheck()
				subMenu1h.Uncheck()
				ChangeDrinkTask(15 * time.Minute)
			case <-subMenu30m.ClickedCh:
				subMenu.SetTitle("30m")
				systray.SetTooltip("当前计时器: 30m")
				subMenu3s.Uncheck()
				subMenu15m.Uncheck()
				subMenu30m.Check()
				subMenu1h.Uncheck()
				ChangeDrinkTask(30 * time.Minute)
			case <-subMenu1h.ClickedCh:
				subMenu.SetTitle("1h")
				systray.SetTooltip("当前计时器: 1h")
				subMenu3s.Uncheck()
				subMenu15m.Uncheck()
				subMenu30m.Uncheck()
				subMenu1h.Check()
				ChangeDrinkTask(time.Hour)

			case <-mNotify.ClickedCh:
				mNotify.Check()
				mMessageBox.Uncheck()
			case <-mMessageBox.ClickedCh:
				mNotify.Uncheck()
				mMessageBox.Check()

			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					StopDrinkTask()
				} else {
					mChecked.Check()
					ResetDrinkTask()
				}

			case <-mQuit.ClickedCh:
				fmt.Println("Quit")
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	// clean up here
}

func main() {
	// 托盘程序逻辑
	systray.Run(onReady, onExit)
}
