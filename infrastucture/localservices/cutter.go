package localservices

import (
	"crypto/rand"
	"fmt"
)

type KeyGenerator struct {
}

func NewKeyGenerator() *KeyGenerator {
	return &KeyGenerator{}
}

func (c KeyGenerator) Generate() string {
	key := make([]byte, 4)
	rand.Read(key)
	return fmt.Sprintf("%x", key)
}
