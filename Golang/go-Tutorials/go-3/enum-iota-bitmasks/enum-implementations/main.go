package main

type HTTPMethod int

func (h HTTPMethod) IsValid() bool {
	panic("implement me")
}

func (h HTTPMethod) String() string {
	panic("implement me")
}

func (h HTTPMethod) UnmarshalJSON(bytes []byte) error {
	panic("implement me")
}

func (h HTTPMethod) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
