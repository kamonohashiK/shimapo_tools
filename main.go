package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kamonohashiK/shimapo_tools/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "shimapo_tools"
	app.Usage = "離島情報共有アプリ 開発者用ツール"
	app.Version = "0.1.0"

	app.Commands = []*cli.Command{
		{
			Name:  "summary_json",
			Usage: "CSVファイルからフロントアプリで使用するJSONファイルを生成",
			Action: func(*cli.Context) error {
				cmd.GenIslandSummaryJsonFromCsv()
				return nil
			},
		},
		{
			Name:  "sample",
			Usage: "コマンドを追加する際のサンプル",
			Action: func(*cli.Context) error {
				fmt.Println("another func")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
