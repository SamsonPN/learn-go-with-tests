package main

// map[key-type]value-type
// the key-type can only be a *comparable* type b/c we need to determine if 2 keys are equal
type Dictionary map[string]string

// after refactoring, turned these errors into constants
// these are immutable!
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryErr("word does not exist in dictionary")
)

// since we convert errors into constants
// we need to create our own DictionaryErr
// which implements the error interface
type DictionaryErr string

// any type with an Error() string method fulfills the error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// map lookup can return TWO values
	// the second value is a boolean which determines if the key is found successfully
	// allows us to differentiate between word that has no definition
	// or word that doesn't exist in the dictionary!
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// why did we not need to pass in a pointer to the map for this???
// A MAP VALUE IS A POINTER TO A runtime.hmap structure!!!

// so you are actually copying the dictionary but you are copying the pointer to the dictionary

// maps can actually be a nil value and behaves like an empty map when reading
// so nilMap[word] will just return something empty
// but if you write to a nilmap, so nilMap[word] = definition, it will cause a runtime panic
// SO NEVER DO var m map[string]string
// just make an empty map like so:
// var dictionary = map[string]string{}
// or
// var dictionary = make(map[string]string)
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	// use of the switch here to match on err
	// provides a safety net in case Search returns an error
	// other than ErrNotFound
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
