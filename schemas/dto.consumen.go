package schemas

type CreateConsumenRequest struct {
	Name        string  `json:"name"  binding:"required"`
	NIK         string  `json:"nik"  binding:"required"`
	FullName    string  `json:"full_name"  binding:"required"`
	LegalName   string  `json:"legal_name"  binding:"required"`
	BirthPlace  string  `json:"birth_place"  binding:"required"`
	BirthDate   string  `json:"birth_date"  binding:"required"`
	Salary      float64 `json:"salary"  binding:"required"`
	KTPPhoto    string  `json:"ktp_photo"  binding:"required"`
	SelfiePhoto string  `json:"selfie_photo"  binding:"required"`
}

type CreateTransactionRequest struct {
	ConsumerID  uint    `json:"consumer_id"  binding:"required"`
	ContractNo  string  `json:"contract_no"  binding:"required"`
	OTR         float64 `json:"otr"  binding:"required"`
	AdminFee    float64 `json:"admin_fee"  binding:"required"`
	Installment float64 `json:"installment"  binding:"required"`
	Interest    float64 `json:"interest"  binding:"required"`
	AssetName   string  `json:"asset_name"  binding:"required"`
}
