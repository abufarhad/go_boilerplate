package client

import (
	"core/infra/config"
	"core/infra/errors"
	"core/infra/logger"
	"core/serializers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func NewIdentity() *serviceClient {
	return initiateHttpClient()
}

type UserResp struct {
	Message string                    `json:"message"`
	Data    serializers.UserWithPerms `json:"data"`
}

type UserListResp struct {
	Message string                    `json:"message"`
	Data    []serializers.MinimalUser `json:"data"`
}

func (ec serviceClient) GetCurrentUser(token string) (*serializers.UserWithPerms, *errors.RestErr) {
	reqURL, _ := url.Parse(config.InternalService().GetCurrentUserUrl)
	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	req := prepareGetRequest(headers, "POST", reqURL)

	res, err := ec.Client.Do(&req)
	if err != nil {
		logger.ErrorAsJson("error response: ", err.Error())
		return nil, errors.NewInternalServerError(err)
	}

	if res.StatusCode == http.StatusOK {
		var respBody *UserResp
		err := json.NewDecoder(res.Body).Decode(&respBody)
		if err != nil {
			logger.ErrorAsJson("", err)
		}
		return &respBody.Data, nil
	} else {
		var respBody *errors.RestErr
		json.NewDecoder(res.Body).Decode(&respBody)

		return nil, respBody
	}
}

func (ec serviceClient) GetUserById(id uint) (*serializers.UserWithPerms, *errors.RestErr) {
	reqURL, _ := url.Parse(fmt.Sprintf("%s/%d", config.InternalService().GetUserByIdUrl, id))
	req := prepareGetRequest(nil, "GET", reqURL)

	res, err := ec.Client.Do(&req)
	if err != nil {
		logger.ErrorAsJson("error response: ", err.Error())
		return nil, errors.NewInternalServerError(err)
	}

	if res.StatusCode == http.StatusOK {
		var respBody *UserResp
		err := json.NewDecoder(res.Body).Decode(&respBody)
		if err != nil {
			logger.ErrorAsJson("", err)
		}
		return &respBody.Data, nil
	} else {
		var respBody *errors.RestErr
		json.NewDecoder(res.Body).Decode(&respBody)

		return nil, respBody
	}
}

func (ec serviceClient) GetAllUsers() ([]serializers.MinimalUser, *errors.RestErr) {
	var userList = make([]serializers.MinimalUser, 0)

	reqURL, _ := url.Parse(config.InternalService().GetUserByIdUrl)
	req := prepareGetRequest(nil, "GET", reqURL)

	res, err := ec.Client.Do(&req)
	if err != nil {
		logger.ErrorAsJson("error response: ", err.Error())
		return userList, errors.NewInternalServerError(err)
	}

	if res.StatusCode == http.StatusOK {
		var respBody *UserListResp
		err := json.NewDecoder(res.Body).Decode(&respBody)
		if err != nil {
			logger.ErrorAsJson("", err)
		}

		return respBody.Data, nil
	} else {
		var respBody *errors.RestErr
		json.NewDecoder(res.Body).Decode(&respBody)

		return userList, respBody
	}
}
