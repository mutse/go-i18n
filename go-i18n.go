package main

import (
	"fmt"
	. "github.com/gosexy/gettext"
)

func main() {
	BindTextdomain("go-i18n", "./")
	Textdomain("go-i18n")
	SetLocale(LcAll, "zh_CN.utf8")
	fmt.Println(Gettext("Hello world!"))
}
