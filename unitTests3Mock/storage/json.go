package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"unitTests3Mock/entity"
)

const filename = "data.json"

func GetPersons() []entity.Person {
	out, err := loadStream()
	if err != nil {
		panic(fmt.Errorf("failed to load people: %v", err))
	}

	return out
}

func loadStream() ([]entity.Person, error) {
	var data []entity.Person

	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return nil, err
	}

	everyRecord, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(everyRecord, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Store(data []entity.Person) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(data)
}
