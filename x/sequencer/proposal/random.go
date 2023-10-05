package proposal

import "golang.org/x/crypto/sha3"

func generateRandomL1(sortKey string) []byte {
	bytes := []byte(sortKey)
	hash := sha3.Sum256(bytes[:])
	return hash[:]
}

func generateRandomL2(sequencerBlockHash []byte, sortKey string) []byte {
	bytes := append(sequencerBlockHash, []byte(sortKey)...)
	hash := sha3.Sum256(bytes[:])
	return hash[:]
}
