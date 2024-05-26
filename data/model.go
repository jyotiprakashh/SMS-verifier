package data

type OTPData struct{
	PhoneNumber string `json:"phone_number, omitempty" validate:"required"`
}

type VerifyData struct{
	User *OTPData `json:"user, omitempty" validate:"required"`
	Code string `json:"code, omitempty" validate:"required"`
}

//curl -H "Content-Type: application/json" -X POST -d '{"phone_number": "+919999999999"}' http://localhost:8000/otp
//curl -H "Content-Type: application/json" -X POST -d '{"phone_number": "+919999999999", "code": "123456"}' http://localhost:8000/verify