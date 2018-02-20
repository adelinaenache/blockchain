package main 

import (
	"bytes"
	"math/big"
	"math"
	"crypto/sha256"
	"strconv"
)

const maxNounce = math.MaxInt64
const targetBytes = 3

// ProofOfWork : proof of work 
type ProofOfWork struct {
	block *Block
	target *big.Int
}

// NewProofOfWork : creates a new proof of work 
func NewProofOfWork (b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBytes))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data, 
			intToHex(int64(pow.block.Timestamp)),
			intToHex(int64(targetBytes)),
			intToHex(int64(nonce)),
		}, 
		[]byte{},
	)

	return data
}

// Run : run the proof of work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNounce {
		data := pow.prepareData(nonce)
		hash := sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		if (hashInt.Cmp(pow.target) == -1) {
			return nonce, hash[:]
		} 
		nonce++
	}

	return nonce, hash[:]
}

// Validate a proof of work
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func intToHex(n int64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}