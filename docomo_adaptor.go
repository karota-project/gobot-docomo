package docomo

import (
	"github.com/hybridgroup/gobot"
)

type DocomoAdaptor struct {
	gobot.Adaptor
	ApiKey string
}

func NewDocomoAdaptor(name string, apiKey string) *DocomoAdaptor {
	return &DocomoAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"docomo.DocomoAdaptor",
		),
		ApiKey: apiKey,
	}
}

func (d *DocomoAdaptor) Connect() bool {
	return true
}

func (d *DocomoAdaptor) Finalize() bool {
	return true
}
