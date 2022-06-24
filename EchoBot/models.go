package main

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"cat_id"`
	Text   string `json:"text"`
}

/*
func main() {
	t := mustToken()

	// token = flags.Get(token)

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}
func mustToken () string {
		token := flag.String (
			name: "token-bot-token",
			value: "",
			usage: "token for acces to tekegram bot"
		)
		 flag.Parse()

		 if *token ==""{
		 log.Fatal(v...:"token is not specified")

		}
	return *token
	}

*/
