package main

import (
	"fmt"
	"github.com/disintegration/imaging"
)

// ImageCutSquare
// 处理图片,把originalPath的图片裁剪成边长side_length大小,保存到savePath
func ImageCutSquare(sidelength int, originalPath, savePath string) {
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
