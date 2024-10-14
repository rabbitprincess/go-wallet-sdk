package runestone

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/stretchr/testify/require"
)

func GetMsgTx(payload string) (*wire.MsgTx, error) {
	var tx *wire.MsgTx = wire.NewMsgTx(wire.TxVersion)

	builder := txscript.NewScriptBuilder().
		AddOp(OP_RETURN).
		AddOp(OP_MAGIC_NUMBER).
		AddData([]byte(payload))

	pkScript, err := builder.Script()
	if err != nil {
		return nil, err
	}
	txOut := wire.NewTxOut(0, pkScript)
	tx.AddTxOut(txOut)
	return tx, nil
}

// func TestBuildOpReturnData_forRunesMain(t *testing.T) {
// 	data := `{"edicts":[{"block":"837557","id":"1234","amount":"100000000000000000100000000000000000","output":0}],"isDefaultOutput":true,"defaultOutput":1,"mint":true,"mintNum":1}`

//		res, err := BuildOpReturnDataJson([]byte(data))
//		require.NoError(t, err)
//		t.Log(res)
//		t.Log(hex.EncodeToString(res))
//		assert.Equal(t, "6a5d0914b58f3314d2091601", hex.EncodeToString(res))
//	}
func TestDecodeArtifact(t *testing.T) {
	for _, test := range []struct {
		artifact string
		payload  string
	}{
		{
			"{\"Cenotaph\":{\"Etching\":null,\"Flaw\":7,\"Mint\":null},\"Runestone\":null}",
			"16020089a5b101bd0580c2d72f00"}, // Cenotaph
	} {
		artifact := &Artifact{}
		err := json.Unmarshal([]byte(test.artifact), artifact)
		require.NoError(t, err)
		fmt.Println(artifact)

		tx, err := GetMsgTx(test.payload)
		require.NoError(t, err)

		rs := &Runestone{}
		artifact2, err := rs.Decipher(tx)
		require.NoError(t, err)

		data, err := json.Marshal(artifact2)
		require.NoError(t, err)

		fmt.Println(string(data))
	}

}

func TestA(t *testing.T) {
	dec, err := hex.DecodeString("16020089a5b101bd0580c2d72f00")
	require.NoError(t, err)
	fmt.Println(string(dec))
}
