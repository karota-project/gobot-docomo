package sentenceunderstanding

import (
	"github.com/hybridgroup/gobot"
	"github.com/karota-project/gobot-docomo"
	"testing"
)

func initTestSentenceUnderstandingDriver() *SentenceUnderstandingDriver {
	return NewSentenceUnderstandingDriver(docomo.NewDocomoAdaptor("myAdaptor", "docomo api key"), "myDriver")
}

func TestSentenceUnderstandingDriverStart(t *testing.T) {
	d := initTestSentenceUnderstandingDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestSentenceUnderstandingDriverHalt(t *testing.T) {
	d := initTestSentenceUnderstandingDriver()
	gobot.Expect(t, d.Halt(), true)
}
