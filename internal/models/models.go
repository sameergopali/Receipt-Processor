package models

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Response structure for the result of GetPoints.
type GetPointsResult struct {
	Points int `json:"points"`
}

// Response  structure for the result structure of ProcessReceipt.
type ProcessReceiptResult struct {
	Id string `json:"id"`
}
