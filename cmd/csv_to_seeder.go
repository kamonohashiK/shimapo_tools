package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kamonohashiK/shimapo_tools/models"
	"github.com/kamonohashiK/shimapo_tools/util"
)

func IslandSeederGenerator() {
	// CSV -> 構造体
	csvData := util.CsvToStruct()

	// islandsテーブル用にデータを変換
	var dbData []models.IslandDbData
	for _, data := range csvData {
		d := models.IslandDbData{
			Id:          data.Id,
			FirestoreId: data.FirebaseId,
			Name:        data.Name,
			Kana:        data.Kana,
			EnName:      data.EnName,
			Latitude:    data.Latitude,
			Longitude:   data.Longitude,
		}
		dbData = append(dbData, d)
	}

	// 結合されたJSONデータをJSONに変換
	combinedJSON, err := json.Marshal(dbData)
	if err != nil {
		log.Fatal("JSONデータの変換に失敗しました。")
	}

	// JSONファイルを作成
	jsonFile, err := os.Create("island_seeder_base.json")
	if err != nil {
		log.Fatal("JSONファイルの作成に失敗しました。")
	}
	defer jsonFile.Close()

	jsonFile.Write(combinedJSON)
}
