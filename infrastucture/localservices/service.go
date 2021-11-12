package localservices

import (
	"crypto/rand"
	"fmt"
)

type Service struct {
	KeyGenerator
}

type KeyGenerator struct {
	//MyRand GenerateInterface
}

func NewKeyGenerator() *KeyGenerator {
	//myRand :=new(myRand)
	//return &KeyGenerator{MyRand: myRand}
	return &KeyGenerator{}
}

func (c KeyGenerator) GenerateKey() string {
	key := make([]byte, 4)
	_, _ = rand.Read(key)
	return fmt.Sprintf("%x", key)
}

//type myRand struct {}
//
//func (mr *myRand)Read(b []byte)(n int,err error)  {
//	return rand.Read(b)
//}

//func (mr *myRand) GenerateKey() string{
//	return "key"
//}
//
