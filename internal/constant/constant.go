package constant

// Constant Payment Method Code
const (
	PaymentMethodCreditCard                = "VC"
	PaymentMethodKlikPayBCA                = "BK"
	PaymentMethodVirtualAccountMandiri     = "M1"
	PaymentMethodVirtualAccountPermataBank = "BT"
	PaymentMethodVirtualAccountCIMBNiaga   = "B1"
	PaymentMethodVirtualAccountBNI         = "I1"
	PaymentMethodVirtualMayBank            = "VA"
	PaymentMethodATMBersama                = "A1"
	PaymentMethodRitel                     = "FT"
	PaymentMethodOVO                       = "OV"
)

// Constant CallBack Code
const (
	CallBackSuccess = "01"
	CallBackFailed  = "00"
)

// Constant Redirect Code
const (
	RedirectSuccess  = "00"
	RedirectPending  = "01"
	RedirectCanceled = "02"
)
