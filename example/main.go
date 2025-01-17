package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/geenkel/systray"
	"github.com/geenkel/systray/example/icon"
	"github.com/labstack/gommon/log"
)

func main() {
	onExit := func() {
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}
	log.Errorf("Starting systray")
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Pretty awesome棒棒嗒")

		items := []*systray.MenuItem{
			{
				Title:   "Item 1",
				Tooltip: "Item 1 tooltip",
			},
			{
				Title:   "Item 2",
				Tooltip: "Item 1 tooltip",
			},
			{
				Title:   "Item 3",
				Tooltip: "Item 1 tooltip",
			},
		}
		systray.SetMenu(items, nil)

		fmt.Println("Adding submenu-----")

		err := items[1].SetMenu(
			[]*systray.MenuItem{
				{
					Title:   "Item 1.1",
					Tooltip: "Item 1.1 tooltip",
				},
				{
					Title:   "Item 1.2",
					Tooltip: "Item 1.2 tooltip",
				},
			},
		)
		if err != nil {
			fmt.Println(err)
		}

		err = items[1].SetMenu(
			[]*systray.MenuItem{
				{
					Title:   "Item 1.1.1",
					Tooltip: "Item 1.1 tooltip",
				},
				{
					Title:   "Item 1.2.2",
					Tooltip: "Item 1.2 tooltip",
				},
			},
		)
		if err != nil {
			fmt.Println(err)
		}

		// systray.AddMenuItem("1", "Ignored")
		// systray.AddMenuItem("2", "Ignored")
		// systray.AddMenuItem("3", "Ignored")
		// item4 := systray.AddMenuItem("4", "Ignored")
		// item4.AddSubMenuItem("4.1", "Ignored")
		// item42 := item4.AddSubMenuItem("4.2", "Ignored")
		// item42.AddSubMenuItem("4.2.1", "Ignored")
		// item42.AddSubMenuItem("4.2.2", "Ignored")
		// item42.AddSubMenuItem("4.2.3", "Ignored")
		// item42.AddSubMenuItem("4.2.3", "Ignored")
		// for i := 0; i < 5; i++ {
		// 	systray.AddMenuItem(fmt.Sprintf("Do something #%d", i), "Does awesome things")
		// }

		// time.Sleep(5 * time.Second)

		// fmt.Println("Adding 100000 menu items")
		// for i := 0; i < 10000000; i++ {
		// 	item := systray.AddMenuItem("test", "test")
		// 	systray.RemoveMenuItem(item)
		// }
		// fmt.Println("Finished to add 1million menu items")

		// 	mChange := systray.AddMenuItem("Change Me", "Change Me")
		// 	mChecked := systray.AddMenuItemCheckbox("Unchecked", "Check Me", true)
		// 	mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		// 	// Sets the icon of a menu item. Only available on Mac.
		// 	mEnabled.SetTemplateIcon(icon.Data, icon.Data)

		// 	systray.AddMenuItem("Ignored", "Ignored")

		// 	subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
		// 	subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
		// 	subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
		// 	subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")

		// 	mUrl := systray.AddMenuItem("Open UI", "my home")
		// 	mQuit := systray.AddMenuItem("退出", "Quit the whole app")

		// 	// Sets the icon of a menu item. Only available on Mac.
		// 	mQuit.SetIcon(icon.Data)

		// 	systray.AddSeparator()
		// 	mToggle := systray.AddMenuItem("Toggle", "Toggle the Quit button")
		// 	shown := true
		// 	toggle := func() {
		// 		if shown {
		// 			subMenuBottom.Check()
		// 			subMenuBottom2.Hide()
		// 			mQuitOrig.Hide()
		// 			mEnabled.Hide()
		// 			shown = false
		// 		} else {
		// 			subMenuBottom.Uncheck()
		// 			subMenuBottom2.Show()
		// 			mQuitOrig.Show()
		// 			mEnabled.Show()
		// 			shown = true
		// 		}
		// 	}

		// 	for {
		// 		select {
		// 		case <-mChange.ClickedCh:
		// 			mChange.SetTitle("I've Changed")
		// 		case <-mChecked.ClickedCh:
		// 			if mChecked.Checked() {
		// 				mChecked.Uncheck()
		// 				mChecked.SetTitle("Unchecked")
		// 			} else {
		// 				mChecked.Check()
		// 				mChecked.SetTitle("Checked")
		// 			}
		// 		case <-mEnabled.ClickedCh:
		// 			mEnabled.SetTitle("Disabled")
		// 			mEnabled.Disable()
		// 		case <-mUrl.ClickedCh:
		// 			open.Run("https://www.getlantern.org")
		// 		case <-subMenuBottom2.ClickedCh:
		// 			panic("panic button pressed")
		// 		case <-subMenuBottom.ClickedCh:
		// 			toggle()
		// 		case <-mToggle.ClickedCh:
		// 			toggle()
		// 		case <-mQuit.ClickedCh:
		// 			systray.Quit()
		// 			fmt.Println("Quit2 now...")
		// 			return
		// 		}
		// 	}
	}()
}
