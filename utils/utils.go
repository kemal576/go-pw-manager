package utils

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetIntParam(name string, r *http.Request) (int, error) {
	var param int
	var err error
	params := httprouter.ParamsFromContext(r.Context())
	paramStr := params.ByName(name)

	/*if paramStr == "" {
		return param, err
	}*/

	param, err = strconv.Atoi(paramStr)
	if err != nil {
		return param, err
	}

	return param, nil
}

func GetStrParam(name string, r *http.Request) (string, error) {
	var param string
	//var err error
	params := httprouter.ParamsFromContext(r.Context())
	param = params.ByName(name)

	/*if paramStr == "" {
		return param, error
	}*/

	return param, nil
}
