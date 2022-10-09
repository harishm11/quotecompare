package contracts

type RateResponse struct {
	Amount      float32
	VehDetails  []VehData
	DiscDetails []DiscData
}
type CvgData struct {
	CoverageCode string
	Amount       float32
}

type VehData struct {
	Vehid      uint
	Amount     float32
	CvgDetails []CvgData
}

type DiscData struct {
	DiscountCode string
}
