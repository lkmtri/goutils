package goutils

import (
	"encoding/json"
	"io/ioutil"
)

type File struct {
	Path string
}

func (f *File) ReadJson(out interface{}) error {
	bytes, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, out)
}
