package util

import (
	"encoding/json"
	"log"
	"os"
)

// Jsonファイルを生成する
func JsonGenerator(data interface{}, fileName string) {
	// 結合されたJSONデータをJSONに変換
	combinedJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal("JSONデータの変換に失敗しました。")
	}

	// JSONファイルを作成
	jsonFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal("JSONファイルの作成に失敗しました。")
	}
	defer jsonFile.Close()

	jsonFile.Write(combinedJSON)
}
