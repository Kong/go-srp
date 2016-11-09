package srp

import (
	"crypto/rand"
	"io"
	"math/big"
)

func GenKey(numBytes int) []byte {
	bytes := make([]byte, numBytes)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		panic("Random source is broken!")
	}

	return bytes
}

func getK(params *SRPParams, S []byte) []byte {
	hashK := params.Hash.New()
	hashK.Write(S)
	return hashToBytes(hashK)
}

func getu(params *SRPParams, A, B *big.Int) *big.Int {
	hashU := params.Hash.New()
	hashU.Write(A.Bytes())
	hashU.Write(B.Bytes())

	return hashToInt(hashU)
}

func getM1(params *SRPParams, A, B, S []byte) []byte {
	hashM1 := params.Hash.New()
	hashM1.Write(A)
	hashM1.Write(B)
	hashM1.Write(S)
	return hashToBytes(hashM1)
}

func getM2(params *SRPParams, A, M, K []byte) []byte {
	hashM1 := params.Hash.New()
	hashM1.Write(A)
	hashM1.Write(M)
	hashM1.Write(K)
	return hashToBytes(hashM1)
}

func getMultiplier(params *SRPParams) *big.Int {
	hashK := params.Hash.New()
	hashK.Write(padToN(params.N, params))
	hashK.Write(padToN(params.G, params))

	return hashToInt(hashK)
}
