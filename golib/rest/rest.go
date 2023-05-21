package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rest struct {
	URL      string
	Username string
	Password string
}

// Send a GET request
func (r *Rest) Get() ([]byte, error) {
	req, err := http.NewRequest("GET", r.URL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.Username, r.Password)
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Send a POST request
func (r *Rest) Post(payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", r.URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.Username, r.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Send a PUT request
func (r *Rest) Put(payload []byte) ([]byte, error) {
	req, err := http.NewRequest("PUT", r.URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.Username, r.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Send a DELETE request
func (r *Rest) Delete() error {
	req, err := http.NewRequest("DELETE", r.URL, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(r.Username, r.Password)

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Failed to delete the resource. Response Status Code: %d", response.StatusCode)
	}

	return nil
}
