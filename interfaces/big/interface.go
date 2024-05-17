package big

import (
	"io"
	"net/http"
)

// bigInterface is a private interface that defines the CRUD operations in this package.
type bigInterface interface {
	Create() (string, error)
	CreateAsync() (string, error)
	Read() (string, error)
	Update() (string, error)
	Delete() (string, error)
}

// BigStruct implements the bigInterface
type BigStruct struct{}

func (b *BigStruct) Create() (string, error) {
	return "Create", nil
}

func (b *BigStruct) Read() (string, error) {
	return "Read", nil
}

func (b *BigStruct) Update() (string, error) {
	return "Update", nil
}

func (b *BigStruct) Delete() (string, error) {
	return "Delete", nil
}

func (b *BigStruct) CreateAsync() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	return "DoSomething", nil
}
