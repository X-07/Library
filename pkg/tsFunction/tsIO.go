package tsfunction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//TraceLog :
var TraceLog *bool //trace dans le fichier log
//TraceConsole :
var TraceConsole *bool //trace sur la console
//FicLog :
var FicLog *string //fichier log
var fileLog *os.File

//FileExists : Teste si un fichier existe
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

//CreateDirIfNotExists : Creation du dossier "dirName" s'il n'existe pas
func CreateDirIfNotExists(dirName string) error {
	var err error = nil
	if !FileExists(dirName) {
		err = os.Mkdir(dirName, 0777)
	}
	return err
}

//CreateFileIfNotExists : Creation du fichier "fileName" s'il n'existe pas
func CreateFileIfNotExists(fileName string) (bool, error) {
	var err error = nil
	var exist = true
	if !FileExists(fileName) {
		exist = false
		_, err = os.Create(fileName)
	}
	return exist, err
}

//OpenLOG : Creation du fichier log s'il n'existe pas, ouverture et ecriture d'un enrg "date et heure"
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

//OpenFileRW : Creation du fichier 'fic' s'il n'existe pas, ouverture
func OpenFileRW(fic string) (*os.File, error) {
	var err error = nil
	_, err = CreateFileIfNotExists(fic)
	if err == nil {
		fileLog, err = os.OpenFile(fic, os.O_APPEND|os.O_RDWR, 0777)
	}
	return fileLog, err
}

//OpenNewFileRW : Creation du fichier 'fic' s'il n'existe pas, ouverture
func OpenNewFileRW(fic string) (*os.File, error) {
	var err error = nil
	if FileExists(fic) {
		err = os.Remove(fic)
	}
	if err == nil {
		_, err = CreateFileIfNotExists(fic)
		if err == nil {
			fileLog, err = os.OpenFile(fic, os.O_RDWR, 0777)
		}
	}
	return fileLog, err
}

//WriteJSONFile : Écriture dans un fichier de données "struct" au format json
func WriteJSONFile(filename string, data interface{}) error {
	fileContent, err := json.Marshal(data)
	if err == nil {
		err = ioutil.WriteFile(filename, fileContent, 0777)
	}
	return err
}

//ReadJSONFile : Lecture dans un fichier au format json et restitution en "struct"
func ReadJSONFile(filename string, data interface{}) error {
	fileData, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(fileData, &data) //DÉCODAGE
	}
	return err
}

//PopLine : Ajoute une ligne en fin de fichier tout en supprimant la première
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

//GetAppPath : Récupère le répertoire de l'application
func GetAppPath() (string, error) {
	appPath, err := os.Executable()
	if err == nil {
		username, err := GetCurrentUser()
		if err == nil {
			myRep := "/home/" + username + "/.Script_GO"

			appPath = filepath.Dir(appPath)
			PrintConsole("appPath - os.Executable()           : " + appPath)

			appPath, err = filepath.EvalSymlinks(appPath)
			if err == nil {
				PrintConsole("appPath - filepath.EvalSymlinks(...): " + appPath)

				parts := strings.Split(appPath, string(os.PathSeparator))
				PrintConsole(parts)
				PrintConsole(filepath.Join(parts...))
				if parts[1] == "tmp" && strings.Contains(parts[2], "go-build") {
					PrintConsole("Session de DEV.")
					appPath = myRep
				}
				if parts[len(parts)-1] == "__debug_bin" {
					PrintConsole("Session de DEBUG")
					appPath = myRep
				}
			}
		}
	}
	return appPath, err
}

//ReadFileForValue : Parcourt un fichier 'fileName' ligne à ligne à la recherche du mot clé 'cleVal' et retourne l'enregistrement le contenant
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

//ReadFileInTab : lire un fichier 'fileName' dans un tableau
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

//PrintLog : Écriture sur le fichier log
func PrintLog(message string) {
	if *TraceLog {
		fileLog.WriteString(message)
	}
}

//PrintConsole : Écriture sur la console
func PrintConsole(message ...interface{}) {
	if *TraceConsole {
		fmt.Println(message...)
	}
}

//Trace : Trace l'info
func Trace(message ...interface{}) {
	PrintLog(fmt.Sprintln(message...))
	PrintConsole(message...)
}

//PrintConsole : Écriture sur la console
//func PrintConsole(message string) {
//	if *TraceConsole {
//		fmt.Println(message)
//	}
//}

//Trace : Trace l'info
//func Trace(message string) {
//	PrintLog(message + "\n")
//	PrintConsole(message)
//}

//FmtConsole : Écriture sur la console
func FmtConsole(a ...interface{}) {
	if *TraceConsole {
		fmt.Println(a...)
	}
}
