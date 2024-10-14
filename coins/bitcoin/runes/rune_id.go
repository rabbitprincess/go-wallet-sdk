package runes

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"lukechampine.com/uint128"
)

func NewRuneIdFromString(s string) (RuneId, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return RuneId{}, ErrSeparator
	}
	block, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return RuneId{}, ErrBlock(parts[0])
	}
	tx, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return RuneId{}, ErrTransaction(parts[1])
	}
	return NewRuneId(block, uint32(tx))
}

func NewRuneId(block uint64, tx uint32) (RuneId, error) {
	if block == 0 && tx > 0 {
		return RuneId{}, errors.New("block=0 but tx>0")
	}
	return RuneId{Block: block, Tx: tx}, nil
}

type RuneId struct {
	Block uint64
	Tx    uint32
}

func (r RuneId) Delta(next RuneId) (uint64, uint32, error) {
	if next.Block < r.Block {
		return 0, 0, fmt.Errorf("next block is less than current block")
	}
	block := next.Block - r.Block
	var tx uint32
	if block == 0 {
		if next.Tx < r.Tx {
			return 0, 0, fmt.Errorf("next tx is less than current tx")
		}
		tx = next.Tx - r.Tx
	} else {
		tx = next.Tx
	}
	return block, tx, nil
}

func (r RuneId) Cmp(other RuneId) int {
	return uint128.New(uint64(r.Tx), r.Block).Cmp(uint128.New(uint64(other.Tx), other.Block))
}

func (r RuneId) Next(block uint128.Uint128, tx uint128.Uint128) (RuneId, error) {
	//check block overflow
	if block.Hi > 0 {
		return RuneId{}, fmt.Errorf("block overflow")
	}
	if tx.Hi > 0 || tx.Lo > math.MaxUint32 {
		return RuneId{}, fmt.Errorf("tx overflow")
	}
	newBlock := r.Block + block.Lo
	//check for overflow
	if newBlock < r.Block {
		return RuneId{}, fmt.Errorf("block overflow")

	}
	var newTx uint32
	if block.IsZero() {
		newTx = r.Tx + uint32(tx.Lo)
		//check for overflow
		if newTx < r.Tx {
			return RuneId{}, fmt.Errorf("tx overflow")
		}
	} else {
		newTx = uint32(tx.Lo)
	}
	return NewRuneId(newBlock, newTx)
}

func (r RuneId) String() string {
	return fmt.Sprintf("%d:%d", r.Block, r.Tx)
}
