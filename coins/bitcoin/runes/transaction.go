package runes

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/okx/go-wallet-sdk/coins/bitcoin"
)

func BuildingEdictTx(runeId RuneId, utxo []*bitcoin.TxInput, fromAddress, toAddress string, toAmount int64, feeRate int64, revealValue int64, net *chaincfg.Params) (*wire.MsgTx, error) {
	rs, err := NewRuneStone().AddEdict(runeId, uint64(toAmount), 2).Encipher()
	if err != nil {
		return nil, err
	}

	builder := bitcoin.NewTxBuild(2, net)

	for _, input := range utxo {
		builder.AddInput(input.TxId, input.VOut, input.PrivateKey, "", input.Address, input.Amount)
	}

	builder.AddOutput2("", hex.EncodeToString(rs), 0)           // reveal the runestone
	builder.AddOutput2(fromAddress, "", RUNE_POSTAGE)           // return the change to the sender
	builder.AddOutput2(toAddress, "", RUNE_POSTAGE)             // send the rune to the receiver
	tx, err := builder.FundRawTransaction(fromAddress, feeRate) // fund the transaction
	if err != nil {
		return nil, err
	}
	return tx, nil
}
