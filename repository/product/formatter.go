package product

type GetResponse struct {
	NameUser    string `json:"nameUser"`
	NameProduct string `json:"nameProduct"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Stock       string `json:"stock"`
	Image       string `json:"image"`
}

type GetResponses struct {
	Responses []GetResponse `json:"responses"`
}
