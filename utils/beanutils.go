package utils

import (
	"github.com/jinzhu/copier"
)

func CopyProperties(target, source interface{}) {
	err := copier.CopyWithOption(target, source, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		panic(err)
	}
}
