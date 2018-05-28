package cerror

type JsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
