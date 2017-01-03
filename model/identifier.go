package model

import (
	"fmt"
)

type Identifier struct {
	ID      string `yaml:"id,omitempty"`
	StackID string `yaml:"stackId,omitempty"`
}

func (c Identifier) HasIdentifier() bool {
	return c.ID != "" || c.StackID != ""
}

func (c Identifier) Ref(logicalName string) string {
	if c.StackID != "" {
		return fmt.Sprintf(`{ "ImportValue" : %q }`, c.StackID)
	} else if c.ID != "" {
		return fmt.Sprintf(`{ "Ref" : %q }`, c.ID)
	} else {
		return fmt.Sprintf(`{ "Ref" : %q }`, logicalName)
	}
}
