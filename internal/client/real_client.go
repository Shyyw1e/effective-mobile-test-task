package client

type RealClient struct{}

func (r *RealClient) GetAge(name string) (int, error) {
	return GetAge(name)
}

func (r *RealClient) GetGender(name string) (string, error) {
	return GetGender(name)
}

func (r *RealClient) GetNationalities(name string) ([]string, error) {
	return GetNationalities(name)
}
