package model

// DuitKuCallBack struct model
type DuitKuCallBack struct {
	MerchantCode    string `form:"merchantCode"`
	Amount          string `form:"amount"`
	MerchantOrderID string `form:"merchantOrderId"`
	ProductDetail   string `form:"productDetail"`
	AdditionalParam string `form:"additionalParam"`
	PaymentCode     string `form:"paymentCode"`
	ResultCode      string `form:"resultCode"`
	MerchantUserID  string `form:"merchantUserId"`
	Reference       string `form:"reference"`
	Signature       string `form:"signature"`
}

// DuitKuRedirectURL struct model
type DuitKuRedirectURL struct {
	MerchantOrderID string `form:"merchantOrderId"`
	ResultCode      string `form:"resultCode"`
	Reference       string `form:"reference"`
}

// DuitKuRequestTransaction struct model
type DuitKuRequestTransaction struct {
	MerchantCode     string             `json:"merchantCode"`
	PaymentAmount    uint               `json:"paymentAmount"`
	MerchantOrderID  string             `json:"merchantOrderId"`
	ProductDetails   string             `json:"productDetails"`
	Email            string             `json:"email"`
	AdditionalParam  string             `json:"additionalParam"`
	PaymentMethod    string             `json:"paymentMethod"`
	MerchantUserInfo string             `json:"merchantUserInfo"`
	CustomerVaName   string             `json:"customerVaName"`
	PhoneNumber      string             `json:"phoneNumber"`
	ItemDetails      []DuitKuItemDetail `json:"itemDetails"`
	ReturnURL        string             `json:"returnUrl"`
	CallbackURL      string             `json:"callbackUrl"`
	Signature        string             `json:"signature"`
	ExpiryPeriod     uint               `json:"expiryPeriod"`
}

// DuitKuResponseTransaction struct model
type DuitKuResponseTransaction struct {
	MerchantCode         string `json:"merchantCode"`
	Reference            string `json:"reference"`
	PaymentURL           string `json:"paymentUrl"`
	VirtualAccountNumber string `json:"vaNumber"`
	Amount               string `json:"amount"`
	StatusMessage        string `json:"statusMessage"`
	StatusCode           string `json:"statusCode"`
}

// DuitKuItemDetail struct models
type DuitKuItemDetail struct {
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
}

// DuitKuRequestCheckTransaction struct model
type DuitKuRequestCheckTransaction struct {
	MerchantCode    string `json:"merchantCode"`
	MerchantOrderID string `json:"merchantOrderId"`
	Signature       string `json:"signature"`
}

// DuitKuResponseCheckTransaction struct model
type DuitKuResponseCheckTransaction struct {
	MerchantCode  string `json:"merchantCode"`
	Amount        string `json:"amount"`
	Reference     string `json:"reference"`
	Fee           string `json:"fee"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}
