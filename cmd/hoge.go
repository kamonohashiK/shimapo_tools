package cmd

import (
	"fmt"

	"github.com/kamonohashiK/shimapo_tools/util"
)

func Hoge() {
	hoge := util.CsvToStruct()
	fmt.Println(hoge)
}
