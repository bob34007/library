package file

import (
	"os"
)

//@GetFileNameFromDir
//parm : in file name
//parm : out file data , error
func ReadAllDataFromFile(filename string) (string, error) {

	var data []byte
	var err error
	data, err = os.ReadFile(filename)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}

}

func WriteDataToFile(filename string, data []byte) error {
	var err error

	return err
}
