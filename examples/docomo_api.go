package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-docomo"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	docomoDriver := docomo.NewDocomoDriver("docomo-d01")

	master.AddRobot(
		gobot.NewRobot(
			"dialogue",
			[]gobot.Connection{},
			[]gobot.Device{docomoDriver},
			func() {
				fmt.Println("work")

				/* API param detail -> https://dev.smt.docomo.ne.jp/?p=api_console.index */
				request := make(map[string]string)
				request["utt"] = "こちらルンバです"
				request["context"] = "53e816d98b3b3"
				request["nickname"] = "光"
				request["nickname_y"] = "ヒカリ"
				request["sex"] = "女"
				request["birthdateY"] = "1997"
				request["birthdateM"] = "5"
				request["birthdateD"] = "30"
				request["age"] = "16"
				request["constellations"] = "双子座"
				request["place"] = "東京"
				request["mode"] = "dialog"

				docomoDriver.POST(request)

			}))

	master.Start()
}
