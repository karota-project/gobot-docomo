package dialogue

import (
	"github.com/hybridgroup/gobot"
	"github.com/karota-project/gobot-docomo"
	"testing"
)

func initTestDialogueDriver() *DialogueDriver {
	return NewDialogueDriver(docomo.NewDocomoAdaptor("myAdaptor", "docomo api key"), "myDriver")
}

func TestDialogueDriverStart(t *testing.T) {
	d := initTestDialogueDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestDialogueDriverHalt(t *testing.T) {
	d := initTestDialogueDriver()
	gobot.Expect(t, d.Halt(), true)
}
