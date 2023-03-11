package images

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"

	"github.com/disintegration/gift"
)

func ResizeImage(file multipart.File, width int, height int, quality int) (*bytes.Buffer, error) {
	img, err := jpeg.Decode(file)

	if err != nil {
		return nil, fmt.Errorf("failed decoding file: %s", err)
	}

	// Resize the image
	g := gift.New(
		gift.Resize(width, height, gift.LanczosResampling),
	)
	resized := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(resized, img)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resized, &jpeg.Options{Quality: quality})

	if err != nil {
		return nil, fmt.Errorf("failed encoding image: %s", err)
	}
	return buf, nil
}
