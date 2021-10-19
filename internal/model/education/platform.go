package education

import (
	"encoding/json"
	"fmt"
)

type Platform struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SiteUrl     string `json:"site_url"`
	Enabled     bool   `json:"enabled"`
}

func (p Platform) String() string {
	encoded, err := json.Marshal(p)
	if err != nil {
		encoded = []byte("")
	}

	return fmt.Sprintf("Platform: %s", encoded)
}
