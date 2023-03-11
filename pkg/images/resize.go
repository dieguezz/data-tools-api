package images

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"strings"

	"github.com/disintegration/gift"
)

func ResizeImage(file multipart.File, format string, width int, height int, quality int) (*bytes.Buffer, error) {
	// Determine the image format

	file.Seek(0, 0) // Reset the file offset
	var err error
	// Create an image decoder based on the format
	var img image.Image
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	case "svg":
		img, _, err = image.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %s", err)
	}

	// Resize the image
	g := gift.New(
		gift.Resize(width, height, gift.LanczosResampling),
	)
	resized := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(resized, img)

	// Create an image encoder based on the format
	var buf bytes.Buffer
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, resized, &jpeg.Options{Quality: quality})
	case "png":
		err = png.Encode(&buf, resized)
	case "gif":
		err = gif.Encode(&buf, resized, nil)
	case "svg":
		return nil, fmt.Errorf("SVG format is not supported for resizing")
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to encode image: %s", err)
	}
	return &buf, nil
}
