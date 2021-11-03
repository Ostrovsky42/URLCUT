package localservices

type KeyGenerator struct {

}

func NewKeyGenerator() *KeyGenerator {
	return &KeyGenerator{}
}

func (c KeyGenerator)Generate()string{
	//genereter
	return "gen"
}