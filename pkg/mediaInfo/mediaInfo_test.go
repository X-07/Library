package mediaInfo

import (
	"fmt"
	"testing"
)

func TestGetMediaInfo(t *testing.T) {
	mediaInfo := GetMediaInfo("/home/jluc1404x/Téléchargements/Movies/Films_Divers/303.Squadron.mkv")
	fmt.Printf("%+v", mediaInfo)
}
