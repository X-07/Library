package tsFunction

import (
	"fmt"
	"os"
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
func CreateFileIfNotExists(fileName string) error {
	var err error = nil
	if !FileExists(fileName) {
		_, err = os.Create(fileName)
	}
	return err
}

//Creation du fichier log s'il n'existe pas, ouverture et ecriture d'un enrg "date et heure"
func OpenLOG(ficLog string) (*os.File, error) {
	var err error = nil
	var fileLog *os.File
	err = CreateFileIfNotExists(ficLog)
	if err == nil {
		fileLog, err = os.OpenFile(ficLog, os.O_APPEND|os.O_WRONLY, 0777)
		if err == nil {
			fileLog.WriteString(fmt.Sprint(time.Now()) + "\n")
		}
	}
	return fileLog, err
}
