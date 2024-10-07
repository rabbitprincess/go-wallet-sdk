package bip32

import (
	"encoding/hex"
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

func GenerateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	return mnemonic, err
}

func GetDerivedPath(index int) string {
	return fmt.Sprintf(`m/44'/60'/0'/0/%d`, index)
}

func GetDerivedPrivateKey(mnemonic string, hdPath string) (string, error) {
	seed := bip39.NewSeed(mnemonic, "")
	rp, err := NewMasterKey(seed)
	if err != nil {
		return "", err
	}
	c, err := rp.NewChildKeyByPathString(hdPath)
	if err != nil {
		return "", err
	}
	childPrivateKey := hex.EncodeToString(c.Key.Key)
	return childPrivateKey, nil
}
