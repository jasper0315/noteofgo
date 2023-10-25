# Start line bot

## contents

## Initialize Line bot　& ngrok
Google it!!!


## Develop a parroting line bot
1. write the code below（以下のcodeを記述）
 ```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("we can't access to dotenv.")
	}

	// ハンドラの登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込みが出来ませんでした: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!!!!"
	fmt.Fprint(w, msg)
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	loadEnv()
	channelSecret := os.Getenv("channel_secret")
	channelAccessToken := os.Getenv("channel_accesstoken")

	// BOTを初期化
	bot, err := linebot.New(
		channelSecret,
		channelAccessToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	// リクエストからBOTのイベントを取得
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
			}
			// 他にもスタンプや画像、位置情報など色々受信可能
		}
	}
}
```
2. save the file and run the code on your terminal(ファイルを保存し、ターミナルで実行)
3. checking the motion on your localhost(https://localhost:8080 で実行できるか確認)
4. open ngrok and copy URL of Forwarding
5. paste the URL into WebhookURL of MessagingAPI in your Line Developers Page
6. <strong>write "/callback" in the end of URL </strong>
