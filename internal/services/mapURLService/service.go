package mapurlservice

type Service interface {
}

type service struct {
}

func ProvideMapURLService() Service {
	return &service{}
}
