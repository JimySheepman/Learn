package main

type iShoe interface {
	setLogo(logo string)
	getLogo() string
	setSize(size int)
	getSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}
