package docomo

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestDocomoDriver() *DocomoDriver {
	return NewDocomoDriver(NewDocomoAdaptor("myAdaptor"), "myDriver")
}

func TestDocomoDriverStart(t *testing.T) {
	d := initTestDocomoDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestDocomoDriverHalt(t *testing.T) {
	d := initTestDocomoDriver()
	gobot.Expect(t, d.Halt(), true)
}
