package tsUtils

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
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

type JpegQuality int64

const (
	NEAREST_NEIGHBOR JpegQuality = iota
	APPROX_BI_LINEAR
	BI_LINEAR
	CATMULL_ROM
)

func ResizeImg(srcPath, resizePath string, defaultImg []byte, newSizeWidth, newSizeHeight int, jpegQuality JpegQuality) {
	input, _ := os.Open(srcPath)
	defer input.Close()
	output, _ := os.Create(resizePath)
	defer output.Close()
	ext := strings.ToLower(filepath.Ext(srcPath))
	var src image.Image
	switch ext {
	case ".png":
		// Decode the image (from PNG to image.Image):
		src, _ = png.Decode(input)
	case ".jpg", ".jpeg":
		// Decode the image (from JPG to image.Image):
		src, _ = jpeg.Decode(input)
	default:
		if err := os.WriteFile(resizePath, defaultImg, 0644); err != nil {
			return
		}
		return
	}
	dst := image.NewRGBA(image.Rect(0, 0, newSizeWidth, newSizeHeight))
	// Resize:
	switch jpegQuality {
	case NEAREST_NEIGHBOR:
		draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	case APPROX_BI_LINEAR:
		draw.ApproxBiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	case BI_LINEAR:
		draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	case CATMULL_ROM:
		draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	}

	switch ext {
	case ".png":
		// Encode to `output`:
		png.Encode(output, dst)
	case ".jpg", ".jpeg", "":
		// Encode to `output`:
		jpeg.Encode(output, dst, nil)
	}
}
