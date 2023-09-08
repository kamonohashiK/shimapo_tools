package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type IslandSummary struct {
	Name   string  `json:"name"`
	Kana   string  `json:"kana"`
	EnName string  `json:"en_name"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}

func main() {
	var combinedData []IslandSummary

	// CSVファイルを開く
	file, err := os.Open("islands.csv")
	if err != nil {
		panic(err)
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
			Name:   island[0],
			Kana:   island[1],
			EnName: island[2],
			Lat:    lat,
			Lng:    lng,
		}

		combinedData = append(combinedData, r)
	}

	// 結合されたJSONデータをJSONに変換
	combinedJSON, err := json.Marshal(combinedData)
	if err != nil {
		panic(err)
	}

	// 結合されたJSONデータをファイルに出力
	outputFile, err := os.Create("islands.json")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(combinedJSON)
	if err != nil {
		panic(err)
	}

	fmt.Println("jsonファイルへの出力完了")
}
