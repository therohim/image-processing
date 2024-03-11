package utils

import (
	"fmt"
	"image"
	"strings"

	"gocv.io/x/gocv"
)

func ResizeImage(filePath, destinationPath string) *string {
	src := gocv.IMRead(filePath, gocv.IMReadColor)
	if src.Empty() {
		Logger.Error(fmt.Sprintf("error reading image from: %s", filePath))
		return nil
	}
	defer src.Close()

	fmt.Println(src.Size())

	dst := gocv.NewMat()
	defer dst.Close()

	gocv.Resize(src, &dst, image.Point{}, 0.5, 0.5, gocv.InterpolationDefault)

	fileArr := strings.Split(filePath, "/")
	newFile := fmt.Sprintf("%s/%s.jpeg", destinationPath, strings.Split(fileArr[len(fileArr)-1], ".")[0])
	if ok := gocv.IMWrite(newFile, dst); !ok {
		Logger.Error("write file is failed")
		return nil
	}

	return &newFile
}
