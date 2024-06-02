package tsMediaInfo

// A lancer par la commande "go test -v *.go"		-v pour écrire sur la console avec t.Log

import (
	"fmt"
	"testing"
)

func TestGetMediaInfo(t *testing.T) {
	mediaInfo := GetMediaInfo("/home/jluc20mx/Téléchargements/Movies/Films/PB/Témoin.muet.mkv")
	t.Log(fmt.Printf("%+v", mediaInfo))
}
