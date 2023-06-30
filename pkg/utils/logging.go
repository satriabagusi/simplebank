/*
Author: Satria Bagus(satria.bagus18@gmail.com)
logging.go (c) 2023
Desc: description
Created:  2023-06-27T23:58:48.659Z
Modified: !date!
*/

package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/satriabagusi/simplebank/internal/entity"
)

func PutToLog(logData entity.Log) (bool, error) {
	log := logData

	putLogFile, err := json.Marshal(log)
	if err != nil {
		return false, err
	}

	f, err := os.OpenFile("data/log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false, err
	}
	defer f.Close()

	n, err := f.Write(putLogFile)
	if err != nil {
		return false, err
	}

	if n, err = f.WriteString(",\n"); err != nil {
		return false, errors.New(string(n))
	}

	return true, nil
}
