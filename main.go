package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
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
func genIslandSummaryJsonFromCsv() {
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

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			genIslandSummaryJsonFromCsv()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
