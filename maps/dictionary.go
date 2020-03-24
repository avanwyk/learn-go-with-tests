package maps

type Dictionary map[string]string

const (
	DefinitionNotFound = DictionaryErr("no definition exists for word")
	DefinitionExists   = DictionaryErr("word already exists")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", DefinitionNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, found := d[word]
	if found {
		return DefinitionExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	d[word] = definition
	return nil
}
