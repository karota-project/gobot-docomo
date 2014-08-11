package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-docomo"
	"github.com/karota-project/gobot-docomo/sentenceunderstanding"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	docomoAdaptor := docomo.NewDocomoAdaptor("docomo-a01", "your-api-key")
	sentenceUnderstandingDriver := sentenceunderstanding.NewSentenceUnderstandingDriver(docomoAdaptor, "sentence-understanding-d01")

	master.AddRobot(
		gobot.NewRobot(
			"sentenceunderstanding",
			[]gobot.Connection{docomoAdaptor},
			[]gobot.Device{sentenceUnderstandingDriver},
			func() {
				fmt.Println("work")

				// 発話理解
				sentenceUnderstandingResult, err := sentenceUnderstandingDriver.Get(sentenceunderstanding.RequestBody{
					ProjectKey: "OSU",
					AppInfo: sentenceunderstanding.AppInfo{
						AppName: "hoge_name",
						AppKey:  "hoge_key",
					},
					ClientVer:  "1.0.0",
					DialogMode: "off",
					Language:   "ja",
					UserId:     "userId",
					Location: sentenceunderstanding.Location{
						Lat: 0,
						Lon: 0,
					},
					UserUtterance: sentenceunderstanding.UserUtterance{
						UtteranceText: "山田さんに電話",
						UtteranceWord: []string{""},
						UtteranceKana: []string{""},
					},
					ProjectSpecific: sentenceunderstanding.ProjectSpecific{},
				})

				if err == nil {
					fmt.Println(sentenceUnderstandingResult)
				} else {
					fmt.Println(err)
				}
			}))

	master.Start()
}
