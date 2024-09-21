package vtupass_go

const SandboxBaseURL = "https://sandbox.vtpass.com/api/"
const LiveEnviromentURL = "https://vtpass.com/api/"

// response codes
const BILLER_CONFIRMED = 020
const INVALID_ARGUMENTS = 011
const PRODUCT_DOES_NOT_EXIST = 012
const BILLER_NOT_REACHABLE_AT_THIS_POINT = 030

const (
	IdentifierAirtime         = "airtime"
	IdentifierData            = "data"
	IdentifierTVSubscription  = "tv-subscription"
	IdentifierElectricityBill = "electricity-bill"
	IdentifierEducation       = "education"
	IdentifierFunds           = "funds"
	IdentifierEvents          = "events"
	IdentifierOtherServices   = "other-services"
	IdentifierInsurance       = "insurance"
)
