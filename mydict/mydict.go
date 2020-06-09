package mydict

import "errors"

// Dictionary - 별칭이 Dictionary인 map
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)

// Search - map에서 키에 해당하는 값을 return
func (d Dictionary) Search(word string) (string, error) {
	// map[key]는 key에 해당하는 값과 존재여부를 return
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add - map에 값이 존재하지 않으면 추가
func (d Dictionary) Add(word, def string) error {
	// map에 값이 있는지 확인하기 위해 만들어 놓은 Search() Method를 사용
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update - mpa에 값이 존재하면 수정
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}
