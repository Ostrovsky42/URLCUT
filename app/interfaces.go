package app

//go:generate mockgen -source=interfaces.go -destination=./mock.go -package=app

type UrlCutterServ interface {
	MakeKey(string) (string, error)
	GetURL(string) (string, error)
	GetUniqueKey() string
}

type UrlSaver interface {
	Save(string, string) error
	Get(string) (string, error)
}
