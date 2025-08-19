package fomat

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Format interface {
	ColorPrint()
}
type Print struct{}

func (p Print) ColorPrint() string {
	return "Print"
}

type Color struct{}

func (c *Color) ColorPrint() string {
	return "ColorPrint"
}

func JsonFormat(jsonStr any) {

	// 使用 MarshalIndent 格式化 JSON 输出
	prettyJSON, err := json.MarshalIndent(jsonStr, "", "    ")
	if err != nil {
		log.Fatalf("格式化JSON失败: %v", err)
	}
	// 输出格式化后的 JSON
	lines := strings.Split(string(prettyJSON), "\n")
	//分割线
	s := PrintColor("JsonFormat", "lanse", "----------------------------------------------------------------------------")
	jsonString := s

	for _, line := range lines {
		jsonString += PrintColor(line, "BrightYellow", "")
	}
	jsonString += s
	fmt.Println(jsonString)
	//PrintColor(string(prettyJSON), "红色")

}

// 用Separator包起来的color颜色返回a,例如(xxx, "red", "==")返回"==\033[31mxxx\033[0m==\n"
func PrintColor(a, color, Separator string) string {
	wordcolor := GetColorCode(color)
	if fmt.Sprintf(" %T", a) == "string" {
	}
	Separator = "\033[95m" + Separator + "\033[0m"
	// 返回带颜色的文本
	return fmt.Sprint(Separator + wordcolor + fmt.Sprint(a) + "\033[0m" + Separator + "\n")
}

// 封装输出不同颜色文字的函数
func PrintColoredLines(title, titleColor, separator string, lines []string, lineColor string) {
	s := PrintColor(title, titleColor, separator)
	jsonString := s

	for _, line := range lines {
		jsonString += PrintColor(line, lineColor, "")
	}
	jsonString += s
	fmt.Println(jsonString)
}
func FormatText(tf TextFormat) string {
	// ANSI 转义码
	style := GetColorCode(tf.Color)
	// 加粗
	if tf.Bold {
		style += "\033[1m"
	}
	// 斜体
	if tf.Italic {
		style += "\033[3m"
	}
	// 下划线
	if tf.Underline {
		style += "\033[4m"
	}

	// 拼接前缀、后缀
	result := tf.Prefix + style + tf.Text + "\033[0m" + tf.Suffix

	return result
}
func GetColorCode(color string) string {
	wordcolor := "\033[37m"
	switch color {
	case "black", "heise", "黑", "黑色", "Black", "\033[30m":
		wordcolor = "\033[30m"
	case "red", "hongse", "红", "红色", "Red", "\033[31m":
		wordcolor = "\033[31m"
	case "green", "lvse", "绿", "绿色", "Green", "\033[32m":
		wordcolor = "\033[32m"
	case "yellow", "huangse", "黄", "黄色", "Yellow", "\033[33m":
		wordcolor = "\033[33m"
	case "blue", "lanse", "蓝", "蓝色", "Blue", "\033[34m":
		wordcolor = "\033[34m"
	case "purple", "zise", "紫", "紫色", "Purple", "magenta", "\033[35m":
		wordcolor = "\033[35m"
	case "cyan", "qingse", "青", "青色", "Cyan", "\033[36m":
		wordcolor = "\033[36m"
	case "white", "baise", "白", "白色", "White", "\033[37m":
		wordcolor = "\033[37m"
	// 亮色
	case "bright_black", "liangheise", "亮黑", "亮黑色", "Gray", "\033[90m":
		wordcolor = "\033[90m"
	case "bright_red", "lianghongse", "亮红", "亮红色", "BrightRed", "\033[91m":
		wordcolor = "\033[91m"
	case "bright_green", "lianglvse", "亮绿", "亮绿色", "BrightGreen", "\033[92m":
		wordcolor = "\033[92m"
	case "bright_yellow", "lianghuangse", "亮黄", "亮黄色", "BrightYellow", "\033[93m":
		wordcolor = "\033[93m"
	case "bright_blue", "lianglanse", "亮蓝", "亮蓝色", "BrightBlue", "\033[94m":
		wordcolor = "\033[94m"
	case "bright_purple", "liangzise", "亮紫", "亮紫色", "BrightPurple", "BrightMagenta", "\033[95m":
		wordcolor = "\033[95m"
	case "bright_cyan", "liangqingse", "亮青", "亮青色", "BrightCyan", "\033[96m":
		wordcolor = "\033[96m"
	case "bright_white", "liangbaise", "亮白", "亮白色", "BrightWhite", "\033[97m":
		wordcolor = "\033[97m"
	default:
		wordcolor = "\033[0m"
	}
	return wordcolor
}
func Println(text string, parameter ...string) {
	var b, i, u bool
	color := ""
	if len(parameter) > 0 {
		for _, param := range parameter {
			switch param {
			case "b", "B":
				b = true
			case "i", "I":
				i = true
			case "u", "U":
				u = true
			default:
				// 如果参数不是 b, i, u，查看是否是color
				color1 := GetColorCode(param)
				if color1 != "\033[0m" {
					color = color1
				}
			}

		}
	}

	tf := TextFormat{
		Text:      text,
		Font:      "",
		Color:     color,
		Prefix:    "",
		Suffix:    "",
		Bold:      b,
		Italic:    i,
		Underline: u,
	}
	fmt.Println(FormatText(tf))
}
