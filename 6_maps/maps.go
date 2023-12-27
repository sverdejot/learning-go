package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrKeyNotFound 		= DictionaryErr("the key was not found in the map")
	ErrKeyAlreadyExists 	= DictionaryErr("they key already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	val, found := d[key]

	if found {
		return val, nil
	} else {
		return "", ErrKeyNotFound
	}
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
		case ErrKeyNotFound:
			d[key] = value
		case nil:
			return ErrKeyAlreadyExists
		default:
			return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
		case nil:
			d[key] = value
		default:
			return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
