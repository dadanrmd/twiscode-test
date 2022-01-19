package errorMessages

import "errors"

var (

	//Error about Application Data
	ErrCustAgeNotMet                       = errors.New("[APP-001] customer age should be between 21 to 55 during the active loan period")
	ErrMerchantTypeNotAllowedForKURProduct = errors.New("[APP-002] merchant type is not allowed for this product")
	ErrMerchantTypeNotFound                = errors.New("[APP-003] merchant type not found")
	ErrApplicationNotFound                 = errors.New("[APP-004] Application not found")

	//Error about NIK : Prefix : NIK
	ErrInvalidNIKLength = errors.New("[NIK-001] invalid NIK length")
	ErrInvalidNIKFormat = errors.New("[NIK-002] invalid NIK format")

	//Error about Whitelisting data
	ErrNoWhitelistingData = errors.New("[WHTLST-001] please insert atleast 1 data of whitelist account")

	//error about phone number
	ErrInvalidPhoneLength          = errors.New("[PHONE-001] invalid phone number length, phone number length must be between 10 - 14")
	ErrInvalidPhoneNumberINDPrefix = errors.New("[PHONE-002] invalid phone number prefix, must be started by +62")
)
