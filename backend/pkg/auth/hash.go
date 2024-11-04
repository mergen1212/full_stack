package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash creates a SHA-256 hash of the input string
func Hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// CompareHash compares a plain text input with a hashed string
func CompareHash(input, hashedStr string) bool {
	return Hash(input) == hashedStr
}
