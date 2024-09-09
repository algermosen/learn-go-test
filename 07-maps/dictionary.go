package maps

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("the definition being added already exists")
	ErrWordDoesNotExist = DictionaryErr("the key does not exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(target string) (r string, err error) {
	r, ok := d[target]
	if !ok {
		err = ErrNotFound
	}
	return
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if err == nil {
		return ErrWordExists
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
