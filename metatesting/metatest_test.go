//go:build !sim

package metatesting_test

import (
	"testing"

	"github.com/glycerine/gosim/metatesting"
)

func TestMetatest(t *testing.T) {
	runner := metatesting.ForOtherPackage(t, "github.com/glycerine/gosim/internal/tests/behavior")
	tests, err := runner.ListTests()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tests)

	run, err := runner.Run(t, &metatesting.RunConfig{
		Test:     "TestHello",
		Seed:     1,
		ExtraEnv: []string{"HELLO=goodbye"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(run.LogOutput))
}
