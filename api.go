package freebox

import (
	"encoding/json"
	_ "github.com/Sirupsen/logrus"
)

type apiResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func (c *Client) GetResult(uri string, payload interface{}) error {
	response := apiResponse{Result: payload}
	body, err := c.GetResource(uri, true)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}
