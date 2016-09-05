package confloader

type Conf struct {
	GoogleClientId string
}

type ConfLoader interface {
	Load() Conf;
}
