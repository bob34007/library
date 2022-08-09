package dir

import (
	"ioutil"
)

// @GetFileNameFromDir
// parm  in : path , dir path
// parm	 out : slice , error

func GetFileNameFromDir(path string) ([]string, error) {
	ss := make([]string, 0)
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		ss = append(ss, f.Name())
	}
	return ss, nil
}
