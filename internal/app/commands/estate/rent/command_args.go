package rent

import "encoding/json"

type commandListArgs struct {
	Page uint64 `json:"page"`
}

func (c *commandListArgs) Unmarshal(text string) error {
	return json.Unmarshal([]byte(text), c)
}

func (c *commandListArgs) ToJsonString() (string, error) {
	data, err := json.Marshal(c)
	return string(data), err
}

type commandDeleteArgs struct {
	ID uint64 `json:"id"`
}

func (c *commandDeleteArgs) Unmarshal(text string) error {
	return json.Unmarshal([]byte(text), c)
}

func (c *commandDeleteArgs) ToJsonString() (string, error) {
	data, err := json.Marshal(c)
	return string(data), err
}
