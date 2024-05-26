package api

import (

	"errors"

	"github.com/twilio/twilio-go"
	twilioAPi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient  = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUTID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {

	params := &twilioAPi.CreateVerificationParams{}

	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	resp, err:= client.VerifyV2.CreateVerification( envSERVICEID() ,params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func (app *Config) twilioVerfifyOTP(phoneNumber string, code string) error {

	params := &twilioAPi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err:= client.VerifyV2.CreateVerificationCheck( envSERVICEID() ,params)
	if err != nil {
		return err
	}
	if *resp.Status!= "approved" {
		return errors.New("Not a valid code")
	}
	return nil
}