package main

import (
	"example.com/mod/fomat"
	"fmt"
	"github.com/disintegration/imaging"
	"time"
)

func main() {
	fomat.Println("Hello, World!", "green")

}

// Path := "/Users/developer/Downloads/"
// originalPath := Path + ""
// savePath := Path + ""
// ImageCutSquare(60, originalPath, savePath)
func apiSpeedLimiter(count float64) {
	speed := 1 / count

	fmt.Println("speed:")
	fmt.Println(speed)
	time.Sleep(time.Duration(speed) * time.Second)
}
func imageCutSquare(sidelength int, originalPath, savePath string) {
	// 1. 打开图片
	src, err := imaging.Open(originalPath)
	if err != nil {
		fmt.Println("打开图片失败")
	}

	// 2. 调整大小（固定 75x75）
	dst := imaging.Resize(src, sidelength, sidelength, imaging.Lanczos)

	// 3. 保存图片
	err = imaging.Save(dst, savePath)
	if err != nil {
		fmt.Println("保存图片失败")
	}

}
