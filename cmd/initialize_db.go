package cmd

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type CsvData struct {
	IslandName string
	IslandId   string
	Location   string
}

var initialQuestions = []string{
	"この島へのアクセス方法について教えてください。",
	"この島の観光スポットについて教えてください。",
	"この島の食事スポットについて教えてください。",
	"この島のプレイスポットについて教えてください。",
	"この島の宿泊施設について教えてください。",
}

func InitializeDb() {
	var combinedData []CsvData

	//CSVから島名とIDを取得
	file, err := os.Open("islands.csv")
	if err != nil {
		log.Fatal("ルートディレクトリにislands.csvが存在しません。")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		island, err := reader.Read()
		if err != nil {
			break
		}

		d := CsvData{
			IslandName: island[0],
			IslandId:   island[10],
			Location:   island[6] + island[7],
		}

		combinedData = append(combinedData, d)
	}

	// Firebaseの初期化
	ctx := context.Background()
	sa := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		fmt.Println(err)
	}

	// Firestoreの初期化
	firestore, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer firestore.Close()

	// Firestoreにデータを登録
	for _, data := range combinedData {
		//islandsコレクションにデータを登録
		islandRef := firestore.Collection("islands").Doc(data.IslandId)
		_, err := islandRef.Set(ctx, map[string]interface{}{
			"name":           data.IslandName,
			"main_image_url": "",
			"location":       data.Location,
		})

		if err != nil {
			log.Fatalf("Failed adding island: %v", err)
		}

		// islandsのサブコレクションにquestionsコレクションを作成
		questionsRef := islandRef.Collection("questions")
		for _, question := range initialQuestions {
			postedByRef := firestore.Doc("user_profiles/ADMIN")
			_, _, err = questionsRef.Add(ctx, map[string]interface{}{
				"question":     question,
				"answer_count": 0,
				"posted_by":    postedByRef,
				"posted_at":    time.Now(),
				"is_default":   true,
			})
			if err != nil {
				log.Fatalf("Failed adding question: %v", err)
			}
		}

		fmt.Fprintf(os.Stdout, "島名: %s を登録しました。\n", data.IslandName)
		// 1秒待機
		time.Sleep(1 * time.Second)
	}
}
