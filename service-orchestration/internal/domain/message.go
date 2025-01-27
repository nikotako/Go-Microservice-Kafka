package domain

// struct order request itu mereperesentasikan permintaan penmesaanan yng d perlukan untuk memproses order
type OrderRequest struct {
	OrderType     string `json:"orderType" binding:"required"`
	TransactionID string `json:"transactionId" binding:"required"`
	UserId        string `json:"userId" binding:"required"`
	PackageId     string `json:"packageId" binding:"required"`
}

// struct ini di gunakan untuk menangani pesan kafka yang masuk
type IncomingMessage struct {
	OrderType     string `json:"orderType"`
	OrderService  string `json:"orderService,omitempty"`
	TransactionId string `json:"transactionId"`
	UserId        string `json:"userId"`
	PackageId     string `json:"packageId"`
	RespCode      int    `json:"respCode,omitempty"`
	RespStatus    string `json:"respStatus,omitempty"`
	RespMessage   string `json:"respMessage,omitempty"`
}
