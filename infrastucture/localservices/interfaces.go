package localservices

//go:generate mockgen -source=interfaces.go -destination=./mock.go -package=localservices
type GenerateInterface interface {
	GenerateKey() string
}

type MyRandInterface interface {
	Read([]byte) (int, error)
}
