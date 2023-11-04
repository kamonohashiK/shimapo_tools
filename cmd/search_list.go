package cmd

import (
	"log"

	"github.com/kamonohashiK/shimapo_tools/util"
)

type ItemForSearch struct {
	Uid    string `json:"uid"`
	Label  string `json:"label"`
	Target string `json:"target"`
}

func GenItemsForSearchJsonFromCsv() {
	// CSV -> 構造体
	csvData := util.CsvToStruct()
	var combinedData []ItemForSearch

	for _, data := range csvData {
		// 島名(都道府県名市区町村名)となる文字列を生成
		var label string = data.Name + "(" + data.Prefecture + data.Cities + ")"
		// 島名 かな アルファベット 都道府県名 市区町村となる文字列を生成
		var target string = data.Name + " " + data.Kana + " " + data.EnName + " " + data.Prefecture + " " + data.Cities

		r := ItemForSearch{
			Uid:    data.FirebaseId,
			Label:  label,
			Target: target,
		}

		combinedData = append(combinedData, r)
	}

	util.JsonGenerator(combinedData, "output/islands_for_search.json")
	log.Print("出力完了")
}
