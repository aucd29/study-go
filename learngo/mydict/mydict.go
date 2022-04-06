package mydict

import "errors"

type Dictionary map[string]string
type Money int

var errNotFoundMap = errors.New("Not Found")
var errExists = errors.New("exists")

func (p Dictionary) Find(key string) (string, error) {
	value, exists := p[key]

	if exists {
		return value, nil
	}

	return "", errNotFoundMap
}

func (p Dictionary) Add(key, value string) error {
	_, err := p.Find(key)

	if err != nil {
		p[key] = value
		return nil
	} else {
		return errExists
	}
}
