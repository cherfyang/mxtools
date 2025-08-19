package fomat

type TextFormat struct {
	Text      string // 文本内容
	Font      string // 字体类型
	Color     string // 颜色
	Prefix    string // 前缀
	Suffix    string // 后缀
	Bold      bool   // 是否加粗
	Italic    bool   // 是否斜体
	Underline bool   // 是否下划线
}
