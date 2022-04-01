package sicgolib

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadFromRequestBody(r io.Reader, d interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &d)
	if err != nil {
		return err
	}

	return nil
}