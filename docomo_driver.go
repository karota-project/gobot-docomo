package docomo

import (
	"github.com/hybridgroup/gobot"
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

func (d *DocomoDriver) adaptor() *DocomoAdaptor {
	return d.Driver.Adaptor().(*DocomoAdaptor)
}

func (d *DocomoDriver) Start() bool {
  return true
}

func (d *DocomoDriver) Halt() bool  {
  return true
}
