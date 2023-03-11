package images

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
)

func GetResizeParams(req *http.Request) (multipart.File, int, int, int, error) {
	// Parse form
	err := req.ParseForm()
	if err != nil {
		return nil, 0, 0, 0, fmt.Errorf(fmt.Sprintf("failed to get file 'attachment': %s", err.Error()))
	}

	// Get file from multipart
	file, _, err := req.FormFile("attachment")
	if err != nil {
		return nil, 0, 0, 0, fmt.Errorf(fmt.Sprintf("failed to get file 'attachment': %s", err.Error()))
	}

	// Get params
	widthParam := req.Form.Get("width")
	heightParam := req.Form.Get("height")
	qualityParam := req.Form.Get("quality")

	// Set default params in case user didnt send them
	if widthParam == "" {
		widthParam = "0"
	}
	if heightParam == "" {
		heightParam = "0"
	}
	if qualityParam == "" {
		qualityParam = "100"
	}

	// Make sure they are int
	width, err := strconv.Atoi(widthParam)
	if err != nil {
		return nil, 0, 0, 0, fmt.Errorf("width needs to be an int %s", err.Error())
	}

	height, err := strconv.Atoi(heightParam)
	if err != nil {
		return nil, 0, 0, 0, fmt.Errorf("height needs to be an int")
	}

	quality, err := strconv.Atoi(qualityParam)
	if err != nil {
		return nil, 0, 0, 0, fmt.Errorf("quality needs to be an int")
	}

	// Check size limits
	if width > 3000 || height > 3000 {
		return nil, 0, 0, 0, fmt.Errorf("max width or height is 3000")
	}

	// Make sure you send either width or height
	if width == 0 && height == 0 {
		return nil, 0, 0, 0, fmt.Errorf("width or height should be bigger than 0")
	}

	return file, width, height, quality, nil
}
