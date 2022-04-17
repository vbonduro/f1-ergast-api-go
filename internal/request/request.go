package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Get performs an HTTP Get request at the given uri and unmarshals the
// resulting JSON into the provided data structure.
func Get(uri string, response interface{}) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}
