package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

func convertFlag(inputFile, outputFile string) error {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	csvData := strings.Replace(string(file), "\n", "", -1)
	csvFile := strings.Split(csvData, ",")
	if len(csvFile) > 512 {
		return errors.New("Only the first 512 sprites can have sprite tiles")
	}

	conv := [512]byte{}
	for i, bstr := range csvFile {
		b, err := stringToByte(bstr)
		if err != nil {
			return err
		}
		conv[i] = b
	}

	content := "package main\n\nvar spriteFlags = [0x200]byte{\n"
	for _, b := range conv {
		content += printByte(b) + ","
	}
	content += "\n}"
	err = ioutil.WriteFile(outputFile, []byte(content), 0666)
	return err
}

func stringToByte(bString string) (byte, error) {
	if len(bString) != 8 {
		return 0, errors.New("string \"" + bString + "\" is not 8 characters long")
	}
	ret := byte(0)
	for i, c := range bString {
		if c != '0' && c != '1' {
			return 0, errors.New("string \"" + bString + "\" contains a character other than 0 and 1")
		}
		if c == '1' {
			ret++
		}
		if i < 7 {
			ret <<= 1
		}
	}
	return ret, nil
}
