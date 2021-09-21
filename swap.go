package common

import (
	"encoding/json"
)

func SwapTo(form, to interface{}) error {
	marshal, err := json.Marshal(form)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, to)
}
