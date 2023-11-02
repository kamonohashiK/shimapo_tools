package cmd

import (
	"github.com/kamonohashiK/shimapo_tools/models"
	"github.com/kamonohashiK/shimapo_tools/util"
)

// islandsテーブルのベースになるJsonファイルを生成
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

// prefecturesテーブルのベースになるJsonファイルを生成
func PrefectureSeederGenerator() {
	// CSV -> 構造体
	csvData := util.CsvToStruct()

	var dbData []models.PrefectureDbData
	var id int = 1
	for _, data := range csvData {
		d := models.PrefectureDbData{
			Id:     id,
			Name:   data.Prefecture,
			EnName: "", //対応するデータがないので空文字にしておく
			Code:   data.PrefCode,
		}

		// 空の場合は追加
		if dbData == nil {
			dbData = append(dbData, d)
			id++
			continue
		} else {
			// 同じ名前の都道府県がなければ追加
			var isExist bool = false
			for _, db := range dbData {
				if db.Name == d.Name {
					isExist = true
					break
				}
			}

			if !isExist {
				dbData = append(dbData, d)
				id++
			}
		}
	}

	// Jsonファイルを生成
	util.JsonGenerator(dbData, "output/prefecture_seeder_base.json")
}
