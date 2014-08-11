package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-docomo"
	"github.com/karota-project/gobot-docomo/dialogue"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	docomoAdaptor := docomo.NewDocomoAdaptor("docomo-a01", "your-api-key")
	dialogueDriver := dialogue.NewDialogueDriver(docomoAdaptor, "dialogue-d01")

	master.AddRobot(
		gobot.NewRobot(
			"dialogue",
			[]gobot.Connection{docomoAdaptor},
			[]gobot.Device{dialogueDriver},
			func() {
				fmt.Println("work")

				// 雑談対話
				dialogueResult, err := dialogueDriver.Get(dialogue.RequestBody{
					Utt:            "こちらルンバです",
					Context:        "53e816d98b3b3",
					Nickname:       "光",
					NicknameY:      "ヒカリ",
					Sex:            "女",
					Bloodtype:      "A",
					BirthdateY:     "1997",
					BirthdateM:     "5",
					BirthdateD:     "30",
					Age:            "16",
					Constellations: "双子座",
					Place:          "東京",
					Mode:           "dialog",
				})

				if err == nil {
					fmt.Println(dialogueResult)
				} else {
					fmt.Println(err)
				}
			}))

	master.Start()
}
