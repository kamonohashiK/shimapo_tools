package cmd

import (
	"log"

	"github.com/kamonohashiK/shimapo_tools/util"
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
	// CSV -> 構造体
	csvData := util.CsvToStruct()
	var combinedData []IslandSummary

	for _, data := range csvData {

		r := IslandSummary{
			Uid:        data.FirebaseId,
			Name:       data.Name,
			Kana:       data.Kana,
			EnName:     data.EnName,
			Lat:        data.Latitude,
			Lng:        data.Longitude,
			Prefecture: data.Prefecture,
			City:       data.Cities,
		}

		combinedData = append(combinedData, r)
	}

	util.JsonGenerator(combinedData, "output/islands.json")
	log.Print("出力完了")
}
