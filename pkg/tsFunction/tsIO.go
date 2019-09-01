package tsIO

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//Teste si un fichier existe
func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else if !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

//Creation du dossier "dirName" s'il n'existe pas
func CreateDirIfNotExists(dirName string) error {
	var err error = nil
	if !FileExists(dirName) {
		err = os.Mkdir(dirName, 0777)
	}
	return err
}

//Creation du fichier "fileName" s'il n'existe pas
func CreateFileIfNotExists(fileName string) (bool, error) {
	var err error = nil
	var exist = true
	if !FileExists(fileName) {
		exist = false
		_, err = os.Create(fileName)
	}
	return exist, err
}

//Creation du fichier log s'il n'existe pas, ouverture et ecriture d'un enrg "date et heure"
func OpenLOG(ficLog string) (*os.File, error) {
	var err error = nil
	var fileLog *os.File
	_, err = CreateFileIfNotExists(ficLog)
	if err == nil {
		fileLog, err = os.OpenFile(ficLog, os.O_APPEND|os.O_WRONLY, 0777)
		if err == nil {
			fileLog.WriteString(fmt.Sprint(time.Now()) + "\n")
		}
	}
	return fileLog, err
}

func WriteJsonFile(filename string, data interface{}) error {
	fileContent, err := json.Marshal(data)
	if err == nil {
		err = ioutil.WriteFile(filename, fileContent, 0777)
	}
	return err
}

func ReadJsonFile(filename string, data interface{}) error {
	fileData, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(fileData, &data) //DECODAGE
	}
	return err
}

func PopLine(filename string, val string) error {
	input, err := ioutil.ReadFile(filename)
	if err == nil {
		lines := strings.Split(string(input), "\n")
		lines = lines[1:]
		lines = append(lines, val)
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filename, []byte(output), 0777)
	}
	return err
}
