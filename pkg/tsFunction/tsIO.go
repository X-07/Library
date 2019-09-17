package tsIO

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const myRep = "/home/jluc1404x/.go/TableauDeBord"

var TraceLog *bool     //trace dans le fichier log
var TraceConsole *bool //trace sur la console
var FicLog *string     //fichier log
var fileLog *os.File

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
	_, err = CreateFileIfNotExists(ficLog)
	if err == nil {
		fileLog, err = os.OpenFile(ficLog, os.O_APPEND|os.O_WRONLY, 0777)
		if err == nil {
			fileLog.WriteString(fmt.Sprint(time.Now()) + "\n")
		}
	}
	return fileLog, err
}

//Ecriture dans un fichier de données "struct" au format json
func WriteJsonFile(filename string, data interface{}) error {
	fileContent, err := json.Marshal(data)
	if err == nil {
		err = ioutil.WriteFile(filename, fileContent, 0777)
	}
	return err
}

//Lecture dans un fichier au format json et restitution en "struct"
func ReadJsonFile(filename string, data interface{}) error {
	fileData, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(fileData, &data) //DECODAGE
	}
	return err
}

//Ajoute une ligne en fin de fichier tout en supprimant la première
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

// récupère le répertoire de l'application
func GetAppPath() (string, error) {
	pid := os.Getpid()
	lnk := "/proc/" + strconv.Itoa(pid) + "/exe"
	appRep, err := os.Readlink(lnk)
	if err == nil {
		parts := strings.Split(appRep, "/")
		if parts[1] == "tmp" && strings.Contains(parts[2], "go-build") {
			appRep = myRep
		} else {
			appRep = ""
			for idx := 1; idx < len(parts)-1; idx++ {
				appRep += "/" + parts[idx]
			}

		}
	}
	return appRep, err
}

// ReadFileForValue(...) parcourt un fichier 'fileName' ligne à ligne à la recherche du mot clé 'cleVal' et retourne l'enregistrement le contenant
func ReadFileForValue(fileName string, cleVal string) (string, error) {
	var enrg string
	input, err := ioutil.ReadFile(fileName)
	if err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			mots := strings.Fields(line)
			for _, mot := range mots {
				if mot == cleVal {
					enrg = line
					break
				}
			}
			if enrg != "" {
				break
			}
		}
	}
	return enrg, err
}

// ReadFileInTab(...) lire un fichier 'fileName' dans un tableau
func ReadFileInTab(fileName string, cleVal string) ([]string, error) {
	var enrgs []string
	input, err := ioutil.ReadFile(fileName)
	if err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			if strings.Contains(line, cleVal) {
				enrgs = append(enrgs, line)
			}
		}
	}
	return enrgs, err
}

// Ecriture sur le fichier log
func PrintLog(message string) {
	if *TraceLog {
		fileLog.WriteString(message)
	}
}

// Ecriture sur la console
func PrintConsole(message ...interface{}) {
	if *TraceConsole {
		fmt.Println(message...)
	}
}

// Trace l'info
func Trace(message ...interface{}) {
	PrintLog(fmt.Sprintln(message...))
	PrintConsole(message...)
}

// Ecriture sur la console
//func PrintConsole(message string) {
//	if *TraceConsole {
//		fmt.Println(message)
//	}
//}

// Trace l'info
//func Trace(message string) {
//	PrintLog(message + "\n")
//	PrintConsole(message)
//}

// Ecriture sur la console
//func FmtConsole(a ...interface{}) {
//	if *TraceConsole {
//		fmt.Println(a...)
//	}
//}
