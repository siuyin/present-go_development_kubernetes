// 10 OMIT
package time

import (
	"fmt"
	"testing"
	gt "time"
)

type fixedClock struct{}

func (t fixedClock) Now() gt.Time {
	return gt.Date(2020, 12, 25, 1, 2, 3, 4, gt.UTC)
}

func TestNow(t *testing.T) {
	fc := fixedClock{}
	if s := Now(fc); s != "01:02:03" {
		t.Errorf("unexpected time: %s", s)
	}
}

// 20 OMIT
func ExampleNow() {
	fc := fixedClock{}
	fmt.Println(Now(fc))
	// Output:
	// 01:02:03
}

// 30 OMIT
