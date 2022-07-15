package model

type Product struct {
	ProductId   string `json:"productId" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	IsStatus    bool   `json:"isStatus"`
	UrlPath     string `json:"urlPath"` //product/images/namaImage.jpg
	ImgPath     string `json:"imgPath,omitempty"`
	//ini nanti buat form time multipart
	// tapi kita nyimpen pathnya
}
