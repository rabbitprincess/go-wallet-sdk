package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/okx/go-wallet-sdk/util"
)

type SignParams struct {
	Type                 int    `json:"type"`
	ChainId              string `json:"chainId"`
	Nonce                string `json:"nonce"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	GasLimit             string `json:"gasLimit"`
	To                   string `json:"to"`
	Value                string `json:"value"`
	Data                 string `json:"data"`
	IsToken              bool   `json:"isToken"`
}

func SignEip1559Transaction(chainId *big.Int, tx *types.Transaction, prvKey *ecdsa.PrivateKey) ([]byte, string, error) {
	signer := types.NewLondonSigner(chainId)
	signedTx, err := types.SignTx(tx, signer, prvKey)
	if err != nil {
		return nil, "", err
	}
	rawTx, err := signedTx.MarshalBinary()
	if err != nil {
		return nil, "", err
	}
	return rawTx, signedTx.Hash().Hex(), nil
}

func toJson(r interface{}) string {
	res, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(res)
}

type SignedTx struct {
	Hash string `json:"hash"`
	Hex  string `json:"hex"`
}

func SignTransaction(txJson, prvHex string) (string, error) {
	if len(txJson) == 0 {
		return "", errors.New("invalid txJson")
	}
	if len(prvHex) == 0 {
		return "", errors.New("invalid prvHex")
	}
	var err error
	var s SignParams
	if err := json.Unmarshal([]byte(txJson), &s); err != nil {
		return "", err
	}
	chainId := util.ConvertToBigInt(s.ChainId)
	var to *common.Address
	if len(s.To) != 0 {
		addr := common.HexToAddress(s.To)
		to = &addr
	}
	var data []byte
	if len(s.Data) != 0 {
		if data, err = util.DecodeHexString(s.Data); err != nil {
			return "", err
		}
	}
	prvBytes, err := hex.DecodeString(prvHex)
	if err != nil {
		return "", errors.New("invalid prvHex")
	}
	var jsonTx Eip1559Token
	if err := json.Unmarshal([]byte(txJson), &jsonTx); err != nil {
		return "", err
	}
	if jsonTx.TxType == types.DynamicFeeTxType { // EIP1559 sign
		prv, _ := btcec.PrivKeyFromBytes(prvBytes)
		tx := NewEip1559Transaction(
			chainId,
			util.ConvertToBigInt(jsonTx.Nonce).Uint64(),
			util.ConvertToBigInt(jsonTx.MaxPriorityFeePerGas),
			util.ConvertToBigInt(jsonTx.MaxFeePerGas),
			util.ConvertToBigInt(jsonTx.GasLimit).Uint64(),
			to,
			util.ConvertToBigInt(jsonTx.Value),
			data,
		)

		res, hash, err := SignEip1559Transaction(chainId, tx, prv.ToECDSA())
		if err != nil {
			return "", err
		}
		return toJson(SignedTx{Hash: hash, Hex: util.EncodeHexWith0x(res)}), nil
	} else {
		prv, _ := btcec.PrivKeyFromBytes(prvBytes)
		// Token processing
		var tx *EthTransaction
		if s.IsToken {
			tx = NewEthTransaction(util.ConvertToBigInt(jsonTx.Nonce), util.ConvertToBigInt(jsonTx.GasLimit), util.ConvertToBigInt(jsonTx.GasPrice), big.NewInt(0), jsonTx.ContractAddress, util.EncodeHexWith0x(data))
		} else {
			tx = NewEthTransaction(util.ConvertToBigInt(jsonTx.Nonce), util.ConvertToBigInt(jsonTx.GasLimit), util.ConvertToBigInt(jsonTx.GasPrice), util.ConvertToBigInt(jsonTx.Value), jsonTx.To, util.EncodeHexWith0x(data))
		}
		res, err := tx.SignTransaction(chainId, prv)
		if err != nil {
			return "", err
		}
		return toJson(SignedTx{Hash: CalTxHash(res), Hex: res}), nil
	}
}
