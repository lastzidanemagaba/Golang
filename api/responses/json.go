package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type meta_resp struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

type Response2 struct {
	Meta meta_resp `json:"meta"`
	//Data []string  `json:"data"`
	Data interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, metaCode int, metaMsg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	res_meta := meta_resp{Code: metaCode, Msg: metaMsg}
	resp := Response2{Meta: res_meta, Data: data}
	//err := json.NewEncoder(w).Encode(data)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		JSON(w, statusCode, 1, "Error", struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, 1, "Error", nil)
}
