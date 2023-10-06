package main

import (
	"log"
	"os"

	"github.com/kamonohashiK/shimapo_tools/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "shimapo_tools"
	app.Usage = "離島情報共有アプリ 開発者用ツール"
	app.Version = "0.1.1"

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
			Name:  "search_list_json",
			Usage: "CSVファイルからフロントアプリの検索時に使用するJSONファイルを生成",
			Action: func(*cli.Context) error {
				cmd.GenItemsForSearchJsonFromCsv()
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
