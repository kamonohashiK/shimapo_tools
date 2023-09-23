package cmd

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type IslandSummary struct {
	Uid        string  `json:"uid"`
	Name       string  `json:"name"`
	Kana       string  `json:"kana"`
	EnName     string  `json:"en_name"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Prefecture string  `json:"prefecture"`
	City       string  `json:"city"`
}

// CSVファイルを読み込み、JSONファイルに変換する
func GenIslandSummaryJsonFromCsv() {
	var combinedData []IslandSummary

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

		lat, _ := strconv.ParseFloat(island[4], 64)
		lng, _ := strconv.ParseFloat(island[5], 64)
		r := IslandSummary{
			Uid:        island[10],
			Name:       island[0],
			Kana:       island[1],
			EnName:     island[2],
			Lat:        lat,
			Lng:        lng,
			Prefecture: island[6],
			City:       island[7],
		}

		combinedData = append(combinedData, r)
	}

	// 結合されたJSONデータをJSONに変換
	combinedJSON, err := json.Marshal(combinedData)
	if err != nil {
		log.Fatal("JSONデータの変換に失敗しました。")
	}

	// 結合されたJSONデータをファイルに出力
	outputFile, err := os.Create("islands.json")
	if err != nil {
		log.Fatal("JSONファイルの出力に失敗しました。")
	}
	defer outputFile.Close()

	_, err = outputFile.Write(combinedJSON)
	if err != nil {
		log.Fatal("JSONファイルの保存に失敗しました。")
	}

	log.Print("出力完了")
}
