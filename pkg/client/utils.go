package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"github.com/spf13/viper"
)

var (
	ErrForbidden           = errors.New("forbidden")
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
)

func ensureStatusOK(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusCreated:
		fallthrough
	case http.StatusNoContent:
		fallthrough
	case http.StatusResetContent:
		fallthrough
	case http.StatusOK:
		return nil
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusInternalServerError:
		return ErrInternalServerError
	default:
		var jerr types.ResponseError
		if err := json.NewDecoder(resp.Body).Decode(&jerr); err != nil {
			return err
		}
		return errors.New(jerr.Error)
	}
}

func substringReplace(str string, dict map[string]string) string {
	for pattern, value := range dict {
		str = strings.Replace(str, pattern, value, -1)
	}
	return str
}

func CreateKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\",", key, value)
	}
	return b.String()
}

func trimURLSlash(url string) string {
	return strings.TrimPrefix(url, "/")
}

func persistClientConfigFile() error {
	return viper.WriteConfigAs(viper.ConfigFileUsed())
}
