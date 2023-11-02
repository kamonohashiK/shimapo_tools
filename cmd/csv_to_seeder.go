package cmd

import (
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

	// Jsonファイルを生成
	util.JsonGenerator(dbData, "output/island_seeder_base.json")
}
