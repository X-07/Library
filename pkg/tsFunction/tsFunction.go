package tsFunction

import "os"

//Teste si un fichier existe
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
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
