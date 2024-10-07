package aptos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAptosWallet_GetRandomPrivateKey(t *testing.T) {
	w := &AptosWallet{}
	p, err := w.GetRandomPrivateKey()
	assert.Nil(t, err)
	t.Log("private key : ", p)
}
