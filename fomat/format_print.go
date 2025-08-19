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
	s := PrintColor("JsonFormat", "laanse", "----------------------------------------------------------------------------")
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
	wordcolor := "\033[37m"
	switch color {
	case "black", "heise", "黑", "黑色", "Black":
		wordcolor = "\033[30m"
	case "red", "hongse", "红", "红色", "Red":
		wordcolor = "\033[31m"
	case "green", "lvse", "绿", "绿色", "Green":
		wordcolor = "\033[32m"
	case "yellow", "huangse", "黄", "黄色", "Yellow":
		wordcolor = "\033[33m"
	case "blue", "lanse", "蓝", "蓝色", "Blue":
		wordcolor = "\033[34m"
	case "purple", "zise", "紫", "紫色", "Purple", "magenta":
		wordcolor = "\033[35m"
	case "cyan", "qingse", "青", "青色", "Cyan":
		wordcolor = "\033[36m"
	case "white", "baise", "白", "白色", "White":
		wordcolor = "\033[37m"
	// 亮色
	case "bright_black", "liangheise", "亮黑", "亮黑色", "Gray":
		wordcolor = "\033[90m"
	case "bright_red", "lianghongse", "亮红", "亮红色", "BrightRed":
		wordcolor = "\033[91m"
	case "bright_green", "lianglvse", "亮绿", "亮绿色", "BrightGreen":
		wordcolor = "\033[92m"
	case "bright_yellow", "lianghuangse", "亮黄", "亮黄色", "BrightYellow":
		wordcolor = "\033[93m"
	case "bright_blue", "lianglanse", "亮蓝", "亮蓝色", "BrightBlue":
		wordcolor = "\033[94m"
	case "bright_purple", "liangzise", "亮紫", "亮紫色", "BrightPurple", "BrightMagenta":
		wordcolor = "\033[95m"
	case "bright_cyan", "liangqingse", "亮青", "亮青色", "BrightCyan":
		wordcolor = "\033[96m"
	case "bright_white", "liangbaise", "亮白", "亮白色", "BrightWhite":
		wordcolor = "\033[97m"
	default:
		wordcolor = "\033[0m"
	}
	if fmt.Sprintf(" %T", a) == "string" {
	}
	Separator = "\033[95m" + Separator + "\033[0m"
	// 返回带颜色的文本
	return fmt.Sprint(Separator + wordcolor + fmt.Sprint(a) + "\033[0m" + Separator + "\n")
}
