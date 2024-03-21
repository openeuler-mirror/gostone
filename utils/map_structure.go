package utils

import (
	"github.com/mitchellh/mapstructure"
)

func Map2Structure(m interface{}, raw interface{}) error {
	if err := mapstructure.Decode(m, raw); err != nil {
		return err
	}
	return nil
}
