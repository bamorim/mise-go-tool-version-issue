package main

//go:generate moq -skip-ensure -stub -pkg mocks -out mocks/greeter_mock.go . Greeter:GreeterMock
type Greeter interface {
	Greet(string) string
}

func main() {
}
