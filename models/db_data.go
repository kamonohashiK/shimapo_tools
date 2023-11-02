package models

/*
*
islandsテーブルのデータを格納する構造体
*/
type IslandDbData struct {
	Id          int     `json:"id"`
	FirestoreId string  `json:"firestore_id"`
	Name        string  `json:"name"`
	Kana        string  `json:"kana"`
	EnName      string  `json:"en_name"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
}
