package model

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	define, ok := d[key]
	if !ok {
		return "", errors.New("didn't find the color you were looking for")
	}
	return define, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
