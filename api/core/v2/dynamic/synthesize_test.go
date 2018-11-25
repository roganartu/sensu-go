package dynamic_test

import (
	"reflect"
	"testing"

	types "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/api/core/v2/dynamic"
)

func TestSynthesizeEvent(t *testing.T) {
	event := types.FixtureEvent("foo", "bar")
	synth := dynamic.Synthesize(event).(map[string]interface{})
	if !reflect.DeepEqual(event.HasCheck(), synth["has_check"]) {
		t.Fatal("bad synthesis")
	}
}
