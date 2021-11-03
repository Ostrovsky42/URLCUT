package app

type KeyGenerator interface {
	MakeKey(string) (string, error)
	GetURL(string) (string, error)
}

type UrlSaver interface {
	Save(string, string) error
	Get(string) (string, error)
	GetKeys() []string
}
