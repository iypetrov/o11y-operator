package controller

import (
	"testing"

	"github.com/iypetrov/o11y-operator/internal/testutils"
)

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
