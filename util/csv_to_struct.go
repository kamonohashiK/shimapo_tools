package util

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/kamonohashiK/shimapo_tools/models"
)

/**
 * islands.csvのデータを構造体に変換する
 */
func CsvToStruct() []models.IslandCsvData {
	var combinedData []models.IslandCsvData

	//CSVから島名とIDを取得
	file, err := os.Open("islands.csv")
	if err != nil {
		log.Fatal("ルートディレクトリにislands.csvが存在しません。")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	index := 1

	for {
		island, err := reader.Read()
		if err != nil {
			break
		}

		latitude, err := strconv.ParseFloat(island[4], 64)
		longitude, err := strconv.ParseFloat(island[5], 64)
		if err != nil {
			log.Fatalf("Failed to parse latitude for island %s: %v", island[0], err)
		}

		d := models.IslandCsvData{
			Id:         index,
			Name:       island[0],
			Kana:       island[1],
			EnName:     island[2],
			Latitude:   latitude,
			Longitude:  longitude,
			Prefecture: island[6],
			Cities:     island[7],
			PrefCode:   island[8],
			CityCode:   island[9],
			FirebaseId: island[10],
		}

		combinedData = append(combinedData, d)
		index++
	}
	return combinedData
}
