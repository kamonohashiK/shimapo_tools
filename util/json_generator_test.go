package util

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonGenerator(t *testing.T) {
	// テストデータをセット
	data := []TestData{
		{Name: "John", Age: 30},
		{Name: "Alice", Age: 25},
	}

	// 関数の呼び出し
	filename := "test.json"
	JsonGenerator(data, filename)

	// 生成されたファイルを読み込み
	bytes, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// json.Unmarshalで構造体に変換
	var result []TestData
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// 復元したものが元のデータと一致するか確認
	if !reflect.DeepEqual(result, data) {
		t.Errorf("Expected %v, got %v", data, result)
	}

	// 生成したファイルを削除
	os.Remove(filename)
}
