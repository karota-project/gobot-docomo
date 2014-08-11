package docomo

import (
	"crypto/tls"
	"fmt"
	"github.com/hybridgroup/gobot"
	"io/ioutil"
	"net/http"
	"strings"
)

type DocomoDriver struct {
	gobot.Driver
}

type DocomoInterface interface {
}

func NewDocomoDriver(name string) *DocomoDriver {
	return &DocomoDriver{
		Driver: *gobot.NewDriver(
			name,
			"docomo.DocomoDriver",
		),
	}
}

// エラー回避　(Adaptorを実装していないため)
// func (d *DocomoDriver) adaptor() *DocomoAdaptor {
// 	return d.Driver.Adaptor().(*DocomoAdaptor)
// }

func (d *DocomoDriver) Start() bool {
	fmt.Println("Start")
	return true
}

func (d *DocomoDriver) Halt() bool {
	fmt.Println("Halt")
	return true
}

func (d *DocomoDriver) POST(req map[string]string) {

	// use TLS
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	s := []string{"{\"utt\": \"", req["utt"],
		"\", \"context\": \"", req["context"],
		"\", \"nickname\": \"", req["nickname"],
		"\", \"nickname_y\": \"", req["nickname_y"],
		"\", \"sex\": \"", req["sex"],
		"\", \"bloodtype\": \"", req["bloodtype"],
		"\", \"birthdateY\": \"", req["birthdateY"],
		"\", \"birthdateM\": \"", req["birthdateM"],
		"\", \"birthdateD\": \"", req["birthdateD"],
		"\", \"age\": \"", req["age"],
		"\", \"constellations\": \"", req["constellations"],
		"\", \"place\": \"", req["place"],
		"\", \"mode\": \"", req["mode"], "\"}"}

	json := strings.Join(s, "")

	b := strings.NewReader(json)

	res, err := client.Post("https://api.apigw.smt.docomo.ne.jp/dialogue/v1/dialogue?APIKEY=75487141684762594c784a515653694e334d53356f5376624e64753332497679734a37526c6d686d43662f",
		"application/json", b)

	if err != nil {
		fmt.Println(fmt.Println(err))
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(fmt.Println(err))
	}

	ss := byteToString(body[:])
	fmt.Println("http_response_body : ", ss)
}

func byteToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}
