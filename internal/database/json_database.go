package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/google/uuid"
)

type JsonDatabase struct {
	Filepath string
	data     map[string]interface{}
}

func readDataFromFile(filepath string) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)

	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func saveDataToFile(filepath string, data map[string]interface{}) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, out, 0755); err != nil {
		return err
	}

	return nil
}

func (jd *JsonDatabase) getData() (map[string]interface{}, error) {
	if jd.data == nil {
		if fileData, err := readDataFromFile(jd.Filepath); err == nil {
			jd.data = fileData
		} else {
			return nil, err
		}
	}

	return jd.data, nil
}

func (jd *JsonDatabase) Save(m interface{}) {
	r := reflect.ValueOf(m)
	id := reflect.Indirect(r).FieldByName("Id")

	if id.String() == "" {
		id.SetString(uuid.NewString())
	}

	data, err := jd.getData()
	if err != nil {
		log.Fatal(err)
	}

	data[id.String()] = m

	if err := saveDataToFile(jd.Filepath, data); err != nil {
		log.Fatal(err)
	}
}
