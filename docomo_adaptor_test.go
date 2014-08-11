package docomo

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestDocomoAdaptor() *DocomoAdaptor {
	return NewDocomoAdaptor("myAdaptor", "api key")
}

func TestDocomoAdaptorConnect(t *testing.T) {
	a := initTestDocomoAdaptor()
	gobot.Expect(t, a.Connect(), true)
}

func TestDocomoAdaptorFinalize(t *testing.T) {
	a := initTestDocomoAdaptor()
	gobot.Expect(t, a.Finalize(), true)
}
