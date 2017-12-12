package models

type (
	// Transport is 托运记录
	Transport struct {
		BaseModel
		WareroomID int      `json:"wareroom_id"`
		ProductID  int      `json:"product_id"`
		Quantity   int      `json:"quantity"`
		Product    Product  `json:"product"`
		Wareroom   Wareroom `json:"wareroom"`
	}
)
