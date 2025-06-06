package client

type Client interface {
	GetAge(name string) (int, error)
	GetGender(name string) (string, error)
	GetNationalities(name string) ([]string, error)
}
