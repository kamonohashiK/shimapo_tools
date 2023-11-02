package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

// citiesテーブル用のJsonファイルを生成
func CitySeederGenerator() {
	// CSV -> 構造体
	csvData := util.CsvToStruct()

	var dbData []models.CityDbData
	var id int = 1

	for _, data := range csvData {
		// CSVデータからCitiesを抽出
		prefecture := data.Prefecture
		str := data.Cities

		// NOTE: ここまでに作成したファイルから都道府県のIDを取得する　→今回は手作業で対応

		//・区切りで分割
		tmpSlice := strings.Split(str, "・")

		for _, city := range tmpSlice {
			d := models.CityDbData{
				Id:           id,
				PrefectureId: 1, // TODO: ここに都道府県IDを入れる
				Name:         city,
				EnName:       "", //対応するデータがないので空文字にしておく
				Code:         "", //対応するデータがないので空文字にしておく
			}

			// 空の場合は追加
			if dbData == nil {
				fmt.Println(prefecture, city)
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
					fmt.Println(prefecture, city)
					dbData = append(dbData, d)
					id++
				}
			}
		}
	}

	// Jsonファイルを生成
	util.JsonGenerator(dbData, "output/city_seeder_base.json")
}

// city_islandsテーブル用のJsonファイルを生成
func CityIslandSeederGenerator() {
	// jsonからcities用のデータを呼び出し
	cityFile, err := os.Open("output/city_seeder_base.json")
	if err != nil {
		log.Fatal("output/city_seeder_base.jsonが存在しません。")
	}
	defer cityFile.Close()

	bytes, err := io.ReadAll(cityFile)
	if err != nil {
		log.Fatal(err)
	}

	var cityData []models.CityDbData
	err = json.Unmarshal(bytes, &cityData)
	if err != nil {
		log.Fatal(err)
	}

	// jsonからislands用のデータを呼び出し
	islandFile, err := os.Open("output/island_seeder_base.json")
	if err != nil {
		log.Fatal("output/island_seeder_base.jsonが存在しません。")
	}

	bytes, err = io.ReadAll(islandFile)
	if err != nil {
		log.Fatal(err)
	}

	var islandData []models.IslandDbData
	err = json.Unmarshal(bytes, &islandData)
	if err != nil {
		log.Fatal(err)
	}

	// CSVデータをベースにcity_islands用のデータを生成
	csvData := util.CsvToStruct()

	var id int = 1
	var dbData []models.CityIslandDbData
	for _, data := range csvData {
		//CSVデータからCitiesを抽出
		cities := data.Cities
		tmpSlice := strings.Split(cities, "・")
		firebaseId := data.FirebaseId

		for _, city := range tmpSlice {
			// cityDataからcityのIDを取得
			var cityId int
			var islandId int
			for _, c := range cityData {
				if c.Name == city {
					cityId = c.Id
					break
				}
			}

			for _, i := range islandData {
				if i.FirestoreId == firebaseId {
					islandId = i.Id
					break
				}
			}

			data := models.CityIslandDbData{
				Id:       id,
				CityId:   cityId,
				IslandId: islandId,
			}
			dbData = append(dbData, data)
			id++
		}
	}

	// Jsonファイルを生成
	util.JsonGenerator(dbData, "output/city_island_seeder_base.json")
}
