package testing

import "errors"

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can't be empty")
	}

	return "Hello " + name, nil
}

func SayHi(name string) error {
	return nil
}
