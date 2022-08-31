package file

import (
	"fmt"
	"os"
)

//@GetFileNameFromDir
//parm : in file name
//parm : out file data , error
func ReadAllDataFromFile(fileName string) (string, error) {

	var data []byte
	var err error
	data, err = os.ReadFile(fileName)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}

}

func WriteDataToFile(fileName string, data string) error {
	var err error
	/******************* 使用 io.WriteString 写入文件 **********************/
	f1, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open write file fail ,", fileName, err)
		return err
	}
	defer f1.Close()
	_, err = f1.WriteString(data) //写入文件(字符串)
	if err != nil {
		//fmt.Println("write file fail ,", fileName, err)
		return err
	}
	//fmt.Printf("写入 %v %v 个字节\n", fileName, n)
	return nil
}
