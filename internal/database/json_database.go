package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/google/uuid"
)

type JsonDatabase struct {
	Filepath string
	data     map[string]interface{}
}

func NewJsonDatabase(filepath string) JsonDatabase {
	return JsonDatabase{
		Filepath: filepath,
	}
}

type Entity interface {
	GetId() string
	SetId(string)
}

var ErrNotFound = errors.New("Entity not found")

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
	out, err := json.MarshalIndent(data, "", "  ")
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

func (jd *JsonDatabase) Save(e Entity) error {
	if e.GetId() == "" {
		e.SetId(uuid.NewString())
	}

	data, err := jd.getData()
	if err != nil {
		return err
	}

	data[e.GetId()] = e

	if err := saveDataToFile(jd.Filepath, data); err != nil {
		return err
	}

	return nil
}

func (jd *JsonDatabase) Get(id string, out interface{}) error {
	data, err := jd.getData()
	if err != nil {
		return err
	}

	if data[id] == nil {
		return ErrNotFound
	}

	// TODO find a better way to convert map[string]interface{} -> struct
	v, _ := json.Marshal(data[id])
	return json.Unmarshal(v, out)
}

func (jd *JsonDatabase) FindFirstBy(property, value string, out interface{}) error {
	data, err := jd.getData()
	if err != nil {
		return nil
	}

	// TODO find a safer way to find via property
	for _, v := range data {
		if entity, ok := v.(map[string]interface{}); ok {
			if entity[property] == value {

				bytes, _ := json.Marshal(entity)
				return json.Unmarshal(bytes, out)
			}
		}
	}

	return ErrNotFound
}
