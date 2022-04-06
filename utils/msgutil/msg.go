package msgutil

import (
	"fmt"
)

type RestResp struct {
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewRestResp(message string, data interface{}) RestResp {
	if message == "" {
		message = "Saved successfully!"
	}
	return RestResp{
		Message: &message,
		Data:    data,
	}
}

func EntityUploadSuccessMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s uploaded successfully", entityName), nil)
}

func EntityGetEmptySuccessMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s has no data", entityName), nil)
}

func EntityUpdateSuccessMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s updated successfully", entityName), nil)
}

func EntityDeleteSuccessMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s deleted successfully", entityName), nil)
}

func EntityNotFoundMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s not found", entityName), nil)
}

func EntityDefaultSetMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s set as default", entityName), nil)
}

func EntityDisabledMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s disable successfully", entityName), nil)
}

func EntityEnableMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("%s enable successfully", entityName), nil)
}

func EntityCreationFailedMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("failed to create %s", entityName), nil)
}

func EntityStructToStructFailedMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("error occur when unmarshalling from struct to struct - %s", entityName), nil)
}

func EntityBindToStructFailedMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("error occur when bind from request to struct - %s", entityName), nil)
}

func EntityParamReadFailedMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("error occur on reading param - %s", entityName), nil)
}

func EntityGenericInvalidMsg(entityName string) RestResp {
	return NewRestResp(fmt.Sprintf("invalid %s", entityName), nil)
}
