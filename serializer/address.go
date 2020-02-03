package serializer

type Address struct {
	Consignee       string `json:"consignee"`
	Tel             string `json:"tel"`
	Province        string `json:"province"`
	City            string `json:"city"`
	County          string `json:"county"`
	Street          string `json:"street"`
	DetailedAddress string `json:"detailed_address"`
}
