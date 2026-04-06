package structs

// BaseResponse adalah format standar JSON API kita
type BaseResponse struct {
	RC   string      `json:"rc"`   // Response Code
	Msg  string      `json:"msg"`  // Message
	Data interface{} `json:"data"` // Payload (Bisa struct, slice, atau null)
}
