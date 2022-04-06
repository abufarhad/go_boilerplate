package methodutil

import (
	"core/infra/errors"
	"core/infra/logger"
	"encoding/json"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func IsInvalid(value string) bool {
	return value == ""
}

func InArray(needle interface{}, haystack interface{}) bool {
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

func IsEmpty(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func MapToStruct(input map[string]interface{}, output interface{}) error {
	if b, err := json.Marshal(input); err == nil {
		return json.Unmarshal(b, &output)
	} else {
		return err
	}
}

func StructToMap(input interface{}, output *map[string]interface{}) error {
	if b, err := json.Marshal(input); err == nil {
		return json.Unmarshal(b, output)
	} else {
		return err
	}
}

func StructToStruct(input interface{}, output interface{}) error {
	if b, err := json.Marshal(input); err == nil {
		return json.Unmarshal(b, &output)
	} else {
		return err
	}
}

func ParseJwtToken(token, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidJwtSigningMethod
		}
		return []byte(secret), nil
	})
}

func StringToIntArray(stringArray []string) []int {
	var res []int

	for _, v := range stringArray {
		if i, err := strconv.Atoi(v); err == nil {
			res = append(res, i)
		}
	}

	return res
}

func GenerateRandomStringOfLength(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	if length == 0 {
		length = 8
	}

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func RecoverPanic() {
	if r := recover(); r != nil {
		logger.ErrorAsJson("error on panic recover: ", r)
	}
}

func ContainsUint(s []uint, item uint) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

func ContainsInt(s []int, item int) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

func ContainsString(s []string, item string) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}
