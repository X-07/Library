package tsUtils

import (
	"bytes"
	"image/jpeg"
	"io"

	"golang.org/x/image/webp"
)

func ConvertWebpToJpgBytes(photo io.Reader) ([]byte, error) {
	img, err := webp.Decode(photo)
	if err == nil {
		jpgBuffer := new(bytes.Buffer)
		err = jpeg.Encode(jpgBuffer, img, &jpeg.Options{Quality: 80})
		if err == nil {
			body := jpgBuffer.Bytes()
			return body, nil
		}
	}
	return nil, err
}
