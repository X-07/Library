package tsUtils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/X-07/Library/gotk3"
	"github.com/gotk3/gotk3/glib"
)

// TraceLog :
var TraceLog *bool //trace dans le fichier log
// TraceConsole :
var TraceConsole *bool //trace sur la console
// FicLog :
var FicLog *string //fichier log
var fileLog *os.File

// FileExists : Teste si un fichier existe
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

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

// CreateDirIfNotExists : Creation du dossier "dirName" s'il n'existe pas
func CreateDirIfNotExists(dirName string) error {
	var err error = nil
	if !FileExists(dirName) {
		err = os.Mkdir(dirName, 0777)
	}
	return err
}

// CreateDirAllIfNotExists : Creation du dossier "dir", ainsi que ses parents, s'il(s) n'existe(nt) pas
func CreateDirAllIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// CreateFileIfNotExists : Creation du fichier "fileName" s'il n'existe pas
func CreateFileIfNotExists(fileName string) (bool, error) {
	var err error = nil
	var exist = true
	if !FileExists(fileName) {
		exist = false
		_, err = os.Create(fileName)
	}
	return exist, err
}

func CreateFile(fic string) *os.File {
	var err error
	if Exists(fic) {
		err = os.Remove(fic)
		if err != nil {
			gotk3.ErrorCheckIHM("failed to delete file: '"+fic+"' ", err)
		}
	}

	file, err := os.Create(fic)
	if err != nil {
		gotk3.ErrorCheckIHM("failed to create file: '"+fic+"' ", err)
	}

	return file
}

// OpenLOG : Creation du fichier log s'il n'existe pas, ouverture et écriture d'un enregistrement "date et heure"
func OpenLOG(ficLog string) *os.File {
	if *TraceLog {
		var err error = nil
		_, err = CreateFileIfNotExists(ficLog)
		if err == nil {
			fileLog, err = os.OpenFile(ficLog, os.O_APPEND|os.O_WRONLY, 0777)
			if err == nil {
				fileLog.WriteString(fmt.Sprint(time.Now()) + "\n")
			} else {
				*TraceLog = false
				fileLog = nil
			}
		} else {
			*TraceLog = false
			fileLog = nil
		}
	}
	return fileLog
}

// CloseLog : Fermeture du fichier log s'il a été ouvert
func CloseLog(fileLog *os.File) {
	if *TraceLog {
		fileLog.Close()
	}
}

// OpenFileRW : Creation du fichier 'fic' s'il n'existe pas, ouverture
func OpenFileRW(fic string) (*os.File, error) {
	var err error = nil
	_, err = CreateFileIfNotExists(fic)
	if err == nil {
		fileLog, err = os.OpenFile(fic, os.O_APPEND|os.O_RDWR, 0777)
	}
	return fileLog, err
}

// OpenNewFileRW : Creation du fichier 'fic' s'il n'existe pas, ouverture
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

// WriteJSONFile : Écriture dans un fichier de données "struct" au format json
func WriteJSONFile(filename string, data interface{}) error {
	fileContent, err := json.Marshal(data)
	if err == nil {
		err = os.WriteFile(filename, fileContent, 0777)
	}
	return err
}

// ReadJSONFile : Lecture dans un fichier au format json et restitution en "struct"
func ReadJSONFile(filename string, data interface{}) error {
	fileData, err := os.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(fileData, &data) //DÉCODAGE
	}
	return err
}

// PopLine : Ajoute une ligne en fin de fichier tout en supprimant la première
func PopLine(filename string, val string) error {
	input, err := os.ReadFile(filename)
	if err == nil {
		lines := strings.Split(string(input), "\n")
		lines = lines[1:]
		lines = append(lines, val)
		output := strings.Join(lines, "\n")
		err = os.WriteFile(filename, []byte(output), 0777)
	}
	return err
}

// GetAppPathDev : Récupère le répertoire de l'application pour une execution en DEV/DEBUG
func GetAppPathDev(myRep string) (string, error) {
	//	myRep := "/home/" + username + "/.Script_GO"
	appPath, err := GetAppPath()
	if err == nil {
		parts := strings.Split(appPath, string(os.PathSeparator))
		Trace(parts)
		Trace(filepath.Join(parts...))
		if parts[1] == "tmp" && strings.Contains(parts[2], "go-build") {
			Trace("Session de DEV.")
			appPath = myRep
		}
		if parts[len(parts)-1] == "__debug_bin" {
			Trace("Session de DEBUG")
			appPath = myRep
		}
	}
	return appPath, err
}

// GetAppPath : Récupère le répertoire de l'application
func GetAppPath() (string, error) {
	appPath, err := os.Executable()
	if err == nil {
		appPath = filepath.Dir(appPath)
		Trace("appPath - os.Executable()           : " + appPath)

		appPath, err = filepath.EvalSymlinks(appPath)
		if err == nil {
			Trace("appPath - filepath.EvalSymlinks(...): " + appPath)
		}
	}
	return appPath, err
}

// ReadFileForValue : Parcourt un fichier 'fileName' ligne à ligne à la recherche du mot clé 'cleVal' et retourne l'enregistrement le contenant
func ReadFileForValue(fileName string, cleVal string) (string, error) {
	var enr string
	input, err := os.ReadFile(fileName)
	if err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			mots := strings.Fields(line)
			for _, mot := range mots {
				if mot == cleVal {
					enr = line
					break
				}
			}
			if enr != "" {
				break
			}
		}
	}
	return enr, err
}

// ReadFileInTab : lire un fichier 'fileName' dans un tableau
func ReadFileInTab(fileName string, cleVal string) ([]string, error) {
	var enrgs []string
	input, err := os.ReadFile(fileName)
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

// PrintLog : Écriture sur le fichier log
func PrintLog(message string) {
	if TraceLog != nil && *TraceLog {
		fileLog.WriteString(message)
	}
}

// PrintConsole : Écriture sur la console
func PrintConsole(message ...interface{}) {
	if TraceConsole != nil && *TraceConsole {
		fmt.Println(message...)
	}
}

// Trace : Trace l'info
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

// FmtConsole : Écriture sur la console
func FmtConsole(a ...interface{}) {
	if TraceConsole != nil && *TraceConsole {
		fmt.Println(a...)
	}
}

func CopyDirectory(scrDir, dest string) error {
	if err := CreateDirAllIfNotExists(dest, 0755); err != nil {
		return err
	}
	entries, err := os.ReadDir(scrDir)
	if err != nil {
		return err
	}

	glib.IdleAdd(func() {
		if gotk3.ProgressBar != nil {
			gotk3.ProgressBar.SetText("Copie des photos de la collection ...")
		}
	})
	nbTotal := len(entries)
	for idx, entry := range entries {
		glib.IdleAdd(func() {
			if gotk3.ProgressBar != nil {
				gotk3.ProgressBar.SetText("Copie des photos de la collection [N° " + ItoA(idx) + "]")
				gotk3.ProgressBar.SetFraction(float64(idx) / float64(nbTotal))
			}
		})

		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateDirAllIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := copySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		fInfo, err := entry.Info()
		if err != nil {
			return err
		}

		isSymlink := fInfo.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, fInfo.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func copySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
