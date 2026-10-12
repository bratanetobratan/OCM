package telegram

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"messsage"`
}

type IncomingMessage struct {
	Chat Chat   `json:"chat"`
	From From   `json:"from"`
	Text string `json:"text"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type From struct {
	Username string `json:"username"`
}
