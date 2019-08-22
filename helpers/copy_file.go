package helpers

import (
	"io/ioutil"
)

func CopyFile(source string, dest string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		return err
	}

	return nil
}
