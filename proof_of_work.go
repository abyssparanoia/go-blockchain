package main

import (
	"bytes"
	"math/big"
)

const targetBits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nouce int) []byte {

	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nouce)),
		},
		[]byte{},
	)

	return data

}

func NewProofOfWork(b *Block) *ProofOfWork {

	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}
