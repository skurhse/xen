package sng

import "fmt"

func NewStackName(tokens ...string) *string {
	name := tokens[0]

	for _, token := range tokens[1:] {
		name = fmt.Sprintf("%s_%s", name, token)
	}

	return &name
}
