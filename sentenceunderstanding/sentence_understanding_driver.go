package sentenceunderstanding

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hybridgroup/gobot"
	"github.com/karota-project/gobot-docomo"
)

const (
	ENDPOINT = "https://api.apigw.smt.docomo.ne.jp/sentenceUnderstanding/v1/task"
)

type SentenceUnderstandingDriver struct {
	gobot.Driver
}

type SentenceUnderstandingInterface interface {
}

type AppInfo struct {
	AppName string `json:"appName"`
	AppKey  string `json:"appKey"` //require
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lat"`
}

type UserUtterance struct {
	UtteranceText    string   `json:"utteranceText"` //require
	UtteranceWord    []string `json:"utteranceWord"` //require response
	UtteranceKana    []string `json:"utteranceKana"`
	UtteranceRevised string   `json:"utteranceRevised"` //require response
}

type ProjectSpecific struct {
	// project specific
}

type RequestBody struct {
	ProjectKey      string          `json:"projectKey"` //fix value "OSU"
	AppInfo         AppInfo         `json:"appInfo"`    //require
	ClientVer       string          `json:"clientVer"`  //fix value "1.0.0"
	DialogMode      string          `json:"dialogMode"` //fix value "off"
	Language        string          `json:"language"`   //require
	UserId          string          `json:"userId"`
	Location        Location        `json:"location"`
	UserUtterance   UserUtterance   `json:"userUtterance"` //require
	ProjectSpecific ProjectSpecific `json:"projectSpecific"`
}

type Command struct {
	CommandId   string `json:"commandId"`   //require
	CommandName string `json:"commandName"` //require
}

type SlotStatus struct {
	SlotName  string `json:"slotName"`
	SlotValue string `json:"slotValue"`
	ValueType string `json:"valueType"`
}

type DialogStatus struct {
	Command    Command      `json:"command"`
	SlotStatus []SlotStatus `json:"slotStatus"`
}

type Content struct {
	ContentSource string `json:"contentSource"` //require
	ContentType   string `json:contentType"`
	ContentEncode string `json:contentEncode"`
	ContentValue  string `json:contentValue"` //require
}

type ExtractedWords struct {
	WordsValue string   `json:"wordsValue"`
	WordsType  []string `json:"wordsType"`
}

type ResponseBody struct {
	ProjectKey      string          `json:"projectKey"` //fix value "OSU"
	AppInfo         AppInfo         `json:"appInfo"`    //require
	ClientVer       string          `json:"clientVer"`  //fix value "1.0.0"
	DialogMode      string          `json:"dialogMode"` //fix value "off"
	Language        string          `json:"language"`   //require
	UserId          string          `json:"userId"`
	DialogStatus    DialogStatus    `json:"dialogStatus"` //require
	Content         Content         `json:"content"`
	ProjectSpecific ProjectSpecific `json:"projectSpecific"`
	UserUtterance   UserUtterance   `json:"userUtterance"`  //require
	TaskIdList      []string        `json:"taskIdList"`     //require
	ServerSendTime  string          `json:"serverSendTime"` //require
}

func NewSentenceUnderstandingDriver(a *docomo.DocomoAdaptor, name string) *SentenceUnderstandingDriver {
	return &SentenceUnderstandingDriver{
		Driver: *gobot.NewDriver(
			name,
			"docomo.SentenceUnderstandingDriver",
			a,
		),
	}
}

func (s *SentenceUnderstandingDriver) adaptor() *docomo.DocomoAdaptor {
	return s.Driver.Adaptor().(*docomo.DocomoAdaptor)
}

func (s *SentenceUnderstandingDriver) Start() bool {
	return true
}

func (s *SentenceUnderstandingDriver) Halt() bool {
	return true
}

func (s *SentenceUnderstandingDriver) Get(requestBody RequestBody) (ResponseBody, error) {
	v, err := json.Marshal(requestBody)

	if err != nil {
		return ResponseBody{}, err
	}

	httpClient := http.Client{}
	resp, err := httpClient.Post(ENDPOINT+"?APIKEY="+s.adaptor().ApiKey, "application/x-www-form-urlencoded", bytes.NewReader(append([]byte("json="), v...)))
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
