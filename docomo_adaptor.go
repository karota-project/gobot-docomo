package docomo

import (
  "github.com/hybridgroup/gobot"
)

type DocomoAdaptor struct {
  gobot.Adaptor
}

func NewDocomoAdaptor(name string) *DocomoAdaptor {
  return &DocomoAdaptor{
    Adaptor: *gobot.NewAdaptor(
      name,
      "docomo.DocomoAdaptor",
    ),
  }
}

func (d *DocomoAdaptor) Connect() bool {
  return true
}

func (d *DocomoAdaptor) Finalize() bool {
  return true
}
