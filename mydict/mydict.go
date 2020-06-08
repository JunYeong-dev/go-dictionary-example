package mydict

import "errors"

// Dictionary - 별칭이 Dictionary인 map
type Dictionary map[string]string

var errNotFound = "Not Found"

// Search - map에서 키에 해당하는 값을 return
func (d Dictionary) Search(word string) (string, error) {
	// map[key]는 key에 해당하는 값과 존재여부를 return
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errors.New(errNotFound)
}
