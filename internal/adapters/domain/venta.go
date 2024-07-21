package domain

type ResponseVenta struct {
	Idventa      string
	NombreVenta  string
	CostodeVenta int
}

type VentaRequest struct {
	Idventa      string `json:"Idventa"`
	NombreVenta  string `json:"NombreVenta"`
	CostodeVenta int    `json:"CostodeVenta"`
}
