package model

import (
	"fmt"
)

type Identifier struct {
	ID          string `yaml:"id,omitempty"`
	ImportID    string `yaml:"importId,omitempty"`
}

func (c Identifier) HasIdentifier() bool {
	return c.ID != "" || c.ImportID != ""
}

func (c Identifier) Ref(logicalName string) string {
	if c.ImportID != "" {
		return fmt.Sprintf(`{ "ImportValue" : %q }`, c.ImportID)
	} else if c.ID != "" {
		return fmt.Sprintf(`{ "Ref" : %q }`, c.ID)
	} else {
		return fmt.Sprintf(`{ "Ref" : %q }`, logicalName)
	}
}
