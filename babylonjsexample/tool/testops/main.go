package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

func main() {
	// Chromeを利用することを宣言
	agoutiDriver := agouti.ChromeDriver(agouti.Browser("chrome"))
	err := agoutiDriver.Start()
	if err != nil {
		panic(err)
	}
	defer agoutiDriver.Stop()
	page, err := agoutiDriver.NewPage()

	if err != nil {
		panic(err.Error())
	}

	// 集計対象のwebアプリを開く
	page.Navigate("https://www.babylonjs-playground.com/#4PP38P#5")
	time.Sleep(3 * time.Second)
	for i := 0; i < 7; i++ {
		time.Sleep(2 * time.Second)
		// ローディングが完了してから、各アニメーションごとに撮影をする
		page.Screenshot(fmt.Sprintf("Screenshot%d.png", i))

	}

	// ハンバーガーメニューをオープンする
	menuOpenButton := page.FindByClass("hamburger-button")

	menuOpenButton.Click()

	time.Sleep(2 * time.Second)

	// 画面撮影したので、パフォーマンス情報をダウンロードする
	selection := page.FindByClass("command-button-icon")

	selection.Click()

	page.Screenshot(fmt.Sprintf("screenShotResult.png"))

	elements, err := page.FindByClass("hambuger-menu").Elements()
	if err != nil {
		panic(err)
	}
	keyword := "INSPECTOR"

	for _, element := range elements {

		childs, _ := element.GetElements(api.Selector{"css selector", ".command-button"})
		for _, childObj := range childs {

			childText, _ := childObj.GetText()
			if strings.Contains(childText, keyword) == false {
				continue
			}

			childObj.Click()

			break

		}

		time.Sleep(2 * time.Second)

		break

	}

	time.Sleep(5 * time.Second)
	page.Screenshot(fmt.Sprintf("screenShotInspectorOpen.png"))

	// 物理エンジンを計算して、速度が落ちているかを確認する

	// 右側のパネルを開く
	sideTabDatabuttonLines, _ := page.FindByID("actionTabs").Elements()

	perforceViewerText := "Open Realtime Perf Viewer"

	for _, buttonLine := range sideTabDatabuttonLines {

		buttonLine.Click()

		time.Sleep(3 * time.Second)

		tabs, _ := buttonLine.GetElements(api.Selector{"css selector", ".tabs"})
		for _, tab := range tabs {

			tab.Click()

			time.Sleep(3 * time.Second)

			labels, _ := tab.GetElements(api.Selector{"css selector", ".labels"})

			for _, labelObj := range labels {
				labelObj.Click()
				time.Sleep(3 * time.Second)
			}

			panes, _ := tab.GetElements(api.Selector{"css selector", ".panes"})

			for _, paneObj := range panes {
				paneDatas, _ := paneObj.GetElements(api.Selector{"css selector", ".pane"})
				for _, paneData := range paneDatas {
					paneButtonLines, _ := paneData.GetElements(api.Selector{"css selector", ".buttonLine"})
					for _, panebuttonLine := range paneButtonLines {
						buttonText, _ := panebuttonLine.GetText()
						if strings.Contains(buttonText, perforceViewerText) {
							panebuttonLine.Click()
							time.Sleep(5 * time.Second)
							page.Screenshot("screenShotViewerOpen.png")
							break
						}
					}
				}
				break
			}

		}
	}

	// babylonjsが重そうだったら、後は修正してねをログに出力する

	//page.FindByClass(".command-button-icon").All().

	//<div class="command-button-icon"><img src="imgs/inspector.svg" class=""></div>

}
