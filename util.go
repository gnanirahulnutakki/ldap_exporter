package main

import (
        "errors"
	"fmt"
	"strings"
)

func checkOverflow(m map[string]interface{}, ctx string) error {
	if len(m) > 0 {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		return errors.New(fmt.Sprintf("Unknown fields in '%s': %s", ctx, strings.Join(keys, ", ")))
	}
	return nil
}
