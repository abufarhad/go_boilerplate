package client

import (
	"core/infra/config"
	"core/infra/errors"
	"core/infra/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func NewMail() *serviceClient {
	return initiateHttpClient()
}

func (ec serviceClient) SendContactMail() error {
	byteData, _ := json.Marshal(map[string]interface{}{})

	logger.InfoAsJson("request body: ", string(byteData))
	reqURL, _ := url.Parse(config.InternalService().ContactEmailUrl)
	req := prepareRequest(byteData, "POST", reqURL)

	res, err := ec.Client.Do(&req)
	if err != nil {
		logger.ErrorAsJson("error response: ", err.Error())
		return err
	}

	if res.StatusCode != http.StatusOK {
		var respBody map[string]string
		json.NewDecoder(res.Body).Decode(&respBody)

		logger.ErrorAsJson("resp on ERROR: ", respBody)
		err := respBody["detail"]
		return errors.NewBadRequestError(err, nil)
	}

	logger.InfoAsJson("email sent", fmt.Sprintf("response %v with status: %v ", res, res.StatusCode))

	return nil
}

//func (ec serviceClient) InviteUsersToEventMail(eventID uint, emails serializers.InviteUsersEmail) error {
//	byteData, _ := json.Marshal(map[string]interface{}{
//		"emails":   emails.Emails,
//		"event_id": eventID,
//	})
//
//	logger.InfoAsJson("request body: ", string(byteData))
//	reqURL, _ := url.Parse(config.InternalService().InviteUsersToEvent)
//	req := prepareRequest(byteData, "POST", reqURL)
//
//	res, err := ec.Client.Do(&req)
//	if err != nil {
//		logger.ErrorAsJson("error response: ", err.Error())
//		return err
//	}
//
//	if res.StatusCode != http.StatusOK {
//		var respBody map[string]string
//		json.NewDecoder(res.Body).Decode(&respBody)
//
//		logger.ErrorAsJson("resp on ERROR: ", respBody)
//		err := respBody["detail"]
//		return errors.NewBadRequestError(err, nil)
//	}
//
//	logger.InfoAsJson("email sent", fmt.Sprintf("response %v with status: %v ", res, res.StatusCode))
//
//	return nil
//}
