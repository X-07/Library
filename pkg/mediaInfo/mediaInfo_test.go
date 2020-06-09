package mediainfo

// A lancer par la commande "go test -v *.go"		-v pour écrire sur la console avec t.Log

import (
	"fmt"
	"testing"
)

func TestGetMediaInfo(t *testing.T) {
	mediaInfo := GetMediaInfo("/home/jluc1804x/Téléchargements/Movies/Films/10.Minutes.Gone.mkv")
	t.Log(fmt.Printf("%+v", mediaInfo))
}
