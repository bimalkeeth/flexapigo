// Code generated by goa v3.0.8, DO NOT EDIT.
//
// flexapi HTTP client CLI support package
//
// Command:
// $ goa gen flexapi/design

package client

import (
	flexapi "flexapi/gen/flexapi"
	"fmt"
	"strconv"
)

// BuildAddPayload builds the payload for the flexapi add endpoint from CLI
// flags.
func BuildAddPayload(flexapiAddA string, flexapiAddB string) (*flexapi.AddPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(flexapiAddA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(flexapiAddB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	payload := &flexapi.AddPayload{
		A: a,
		B: b,
	}
	return payload, nil
}
