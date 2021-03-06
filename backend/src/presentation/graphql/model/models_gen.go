// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ServerEnvType string

const (
	ServerEnvTypeLocal      ServerEnvType = "local"
	ServerEnvTypeDev        ServerEnvType = "dev"
	ServerEnvTypeStaging    ServerEnvType = "staging"
	ServerEnvTypeProduction ServerEnvType = "production"
)

var AllServerEnvType = []ServerEnvType{
	ServerEnvTypeLocal,
	ServerEnvTypeDev,
	ServerEnvTypeStaging,
	ServerEnvTypeProduction,
}

func (e ServerEnvType) IsValid() bool {
	switch e {
	case ServerEnvTypeLocal, ServerEnvTypeDev, ServerEnvTypeStaging, ServerEnvTypeProduction:
		return true
	}
	return false
}

func (e ServerEnvType) String() string {
	return string(e)
}

func (e *ServerEnvType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServerEnvType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServerEnvType", str)
	}
	return nil
}

func (e ServerEnvType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
