package models

/**
 * islands.csvのデータを格納する構造体
 */
type IslandCsvData struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Kana       string  `json:"kana"`
	EnName     string  `json:"en_name"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lng"`
	Prefecture string  `json:"prefecture"`
	Cities     string  `json:"cities"`
	PrefCode   string  `json:"pref_code"`
	CityCode   string  `json:"city_code"`
	FirebaseId string  `json:"firebase_id"`
}
