package httputil

import (
	"core/infra/errors"
	"core/serializers"
	"core/utils/consts"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseIdParam(c echo.Context, paramName string) (uint, *errors.RestErr) {
	param, err := strconv.Atoi(c.Param(paramName))
	if err != nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("%v: No parameter found or parameter is not a number", paramName), err)
	}

	return uint(param), nil
}

func ParseParam(c echo.Context, paramName string) (string, *errors.RestErr) {
	param := c.Param(paramName)
	if param == "" {
		return "", errors.NewBadRequestError(fmt.Sprintf("%v: No parameter found", paramName), nil)
	}

	return param, nil
}

func ParseBody(c echo.Context, payload interface{}, payloadName string) *errors.RestErr {
	if err := c.Bind(&payload); err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("%v: Failed to bind request body parameters", payloadName), err)
	}
	return nil
}

func ParseBearerToken(c echo.Context) (string, *errors.RestErr) {
	bearerToken := c.Request().Header.Get(consts.AuthorizationTokenHeader)

	if bearerToken == "" {
		return "", errors.NewBadRequestError("Missing token!", nil)
	}

	splittedToken := strings.Split(strings.TrimSpace(bearerToken), " ")

	if len(splittedToken) != 2 {
		return "", errors.NewBadRequestError("Invalid bearer token!", nil)
	}

	token := splittedToken[1]
	return token, nil
}

func GetLoggedInUser(c echo.Context) (serializers.LoggedInUser, *errors.RestErr) {
	var currentUser serializers.LoggedInUser

	userIDString := c.Request().Header.Get(consts.UserIDHeader)
	userEmail := c.Request().Header.Get(consts.UserEmailHeader)
	roles := c.Request().Header.Get(consts.RolesHeader)
	permissions := c.Request().Header.Get(consts.PermissionsHeader)

	userID, _ := strconv.Atoi(userIDString)

	if userID == 0 {
		return currentUser, errors.NewBadRequestError(errors.ErrNoLoggedInUser, nil)
	}

	currentUser = serializers.LoggedInUser{
		ID:          userID,
		Email:       userEmail,
		Role:        roles,
		Permissions: permissions,
	}

	return currentUser, nil
}

func BaseUrl(c echo.Context) string {
	baseUrl := c.Scheme() + "://" + c.Request().Host
	return baseUrl
}
