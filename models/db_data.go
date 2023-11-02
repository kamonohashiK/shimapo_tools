package models

// islandsテーブルのデータを格納する構造体
type IslandDbData struct {
	Id          int     `json:"id"`
	FirestoreId string  `json:"firestore_id"`
	Name        string  `json:"name"`
	Kana        string  `json:"kana"`
	EnName      string  `json:"en_name"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
}

// prefecturesテーブルのデータを格納する構造体
type PrefectureDbData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	EnName string `json:"en_name"`
	Code   string `json:"code"`
}

// citiesテーブルのデータを格納する構造体
type CityDbData struct {
	Id           int    `json:"id"`
	PrefectureId int    `json:"prefecture_id"`
	Name         string `json:"name"`
	EnName       string `json:"en_name"`
	Code         string `json:"code"`
}
