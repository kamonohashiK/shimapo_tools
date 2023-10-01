package cmd

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

type ItemForSearch struct {
	Uid    string `json:"uid"`
	Label  string `json:"label"`
	Target string `json:"target"`
}

func GenItemsForSearchJsonFromCsv() {
	var combinedData []ItemForSearch

	// CSVファイルを開く
	file, err := os.Open("islands.csv")
	if err != nil {
		log.Fatal("ルートディレクトリにislands.csvが存在しません。")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		island, err := reader.Read()
		if err != nil {
			break
		}

		// 島名(都道府県名市区町村名)となる文字列を生成
		var label string = island[0] + "(" + island[6] + island[7] + ")"
		// 島名 かな アルファベット 都道府県名 市区町村となる文字列を生成
		var target string = island[0] + " " + island[1] + " " + island[2] + " " + island[6] + " " + island[7]

		r := ItemForSearch{
			Uid:    island[10],
			Label:  label,
			Target: target,
		}

		combinedData = append(combinedData, r)
	}

	// 結合されたJSONデータをJSONに変換
	combinedJSON, err := json.Marshal(combinedData)
	if err != nil {
		log.Fatal("JSONデータの変換に失敗しました。")
	}

	// JSONファイルを作成
	jsonFile, err := os.Create("islands_for_search.json")
	if err != nil {
		log.Fatal("JSONファイルの作成に失敗しました。")
	}
	defer jsonFile.Close()

	jsonFile.Write(combinedJSON)
}
