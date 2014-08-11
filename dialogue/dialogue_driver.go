package dialogue

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hybridgroup/gobot"
	"github.com/karota-project/gobot-docomo"
)

const (
	ENDPOINT = "https://api.apigw.smt.docomo.ne.jp/dialogue/v1/dialogue"
)

type DialogueDriver struct {
	gobot.Driver
}

type DialogueInterface interface {
}

type RequestBody struct {
	Utt            string `json:"utt"` //require
	Context        string `json:"context"`
	Nickname       string `json:"nickname"`
	NicknameY      string `json:"nickname_y"`
	Sex            string `json:"sex"`
	Bloodtype      string `json:"bloodtype"`
	BirthdateY     string `json:"birthdateY"`
	BirthdateM     string `json:"birthdateM"`
	BirthdateD     string `json:"birthdateD"`
	Age            string `json:"age"`
	Constellations string `json:"constellations"`
	Place          string `json:"place"`
	Mode           string `json:"mode"`
}

type ResponseBody struct {
	Utt     string `json:"utt"`
	Yomi    string `json:"yomi"`
	Mode    string `json:"mode"`
	Da      string `json:"da"`
	Context string `json:"context"`
}

func NewDialogueDriver(a *docomo.DocomoAdaptor, name string) *DialogueDriver {
	d := &DialogueDriver{
		Driver: *gobot.NewDriver(
			name,
			"docomo.DialogueDriver",
			a,
		),
	}

	d.AddCommand("Get", func(params map[string]interface{}) interface{} {
		requestBody := RequestBody{}
		requestBody.Utt = params["utt"].(string) //require

		if value, ok := params["context"].(string); ok {
			requestBody.Context = value
		}

		if value, ok := params["nickname"].(string); ok {
			requestBody.Nickname = value
		}

		if value, ok := params["nicknameY"].(string); ok {
			requestBody.NicknameY = value
		}

		if value, ok := params["sex"].(string); ok {
			requestBody.Sex = value
		}

		if value, ok := params["bloodtype"].(string); ok {
			requestBody.Bloodtype = value
		}

		if value, ok := params["birthdateY"].(string); ok {
			requestBody.BirthdateY = value
		}

		if value, ok := params["birthdateM"].(string); ok {
			requestBody.BirthdateM = value
		}

		if value, ok := params["birthdateD"].(string); ok {
			requestBody.BirthdateD = value
		}

		if value, ok := params["age"].(string); ok {
			requestBody.Age = value
		}

		if value, ok := params["constellations"].(string); ok {
			requestBody.Constellations = value
		}

		if value, ok := params["place"].(string); ok {
			requestBody.Place = value
		}

		if value, ok := params["mode"].(string); ok {
			requestBody.Mode = value
		}

		return resultApi(d.Get(requestBody))
	})

	return d
}

func (d *DialogueDriver) adaptor() *docomo.DocomoAdaptor {
	return d.Driver.Adaptor().(*docomo.DocomoAdaptor)
}

func (d *DialogueDriver) Start() bool {
	return true
}

func (d *DialogueDriver) Halt() bool {
	return true
}

func (d *DialogueDriver) Get(requestBody RequestBody) (ResponseBody, error) {
	v, err := json.Marshal(requestBody)

	if err != nil {
		return ResponseBody{}, err
	}

	httpClient := http.Client{}
	resp, err := httpClient.Post(ENDPOINT+"?APIKEY="+d.adaptor().ApiKey, "application/json", bytes.NewReader(v))
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ResponseBody{}, err
	}

	responseBody := ResponseBody{}
	err = json.Unmarshal(contents, &responseBody)

	if err == nil {
		return responseBody, err
	} else {
		return ResponseBody{}, err
	}
}

func resultApi(v interface{}, err error) interface{} {
	if err == nil {
		return struct {
			Result interface{} `json:"result"`
		}{v}
	} else {
		return struct {
			Result error `json:"result"`
		}{err}
	}
}
