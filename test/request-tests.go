package gorch

import (
	"gorch/internal"
	"testing"
)

func TestSetQueryString(t *testing.T) {
	gorchRequest = GorchRequest.create("http://url", "method")
}
