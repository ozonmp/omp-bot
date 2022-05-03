package path

import (
	"errors"
	"fmt"
	"strings"
)

type CallbackPath struct {
	Business     string
	Travel       string
	CallbackName string
	CallbackData string
}

var ErrUnknownCallback = errors.New("unknown callback")

func ParseCallback(callbackData string) (CallbackPath, error) {
	callbackParts := strings.SplitN(callbackData, "__", 4)
	if len(callbackParts) != 4 {
		return CallbackPath{}, ErrUnknownCallback
	}

	return CallbackPath{
		Business:     callbackParts[0],
		Travel:       callbackParts[1],
		CallbackName: callbackParts[2],
		CallbackData: callbackParts[3],
	}, nil
}

func (p CallbackPath) String() string {
	return fmt.Sprintf("%s__%s__%s__%s", p.Business, p.Travel, p.CallbackName, p.CallbackData)
}
