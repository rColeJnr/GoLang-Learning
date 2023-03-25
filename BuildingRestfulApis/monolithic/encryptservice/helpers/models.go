package helpers

// EncryptRequest coming from client
type EncryptRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

// EncryptResponse going to the client
type EncryptResponse struct {
	Message string `json:"message"`
	Err     string `json:"err"`
}

// DecryptRequest coming from client
type DecryptRequest struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

// DecryptResponse going to the client
type DecryptResponse struct {
	Text string `json:"text"`
	Err  string `json:"err"`
}
