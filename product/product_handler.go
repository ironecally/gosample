package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type jsonReturn struct {
	StatusCode   int         `json:"status_code"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.FormValue("product_id")
	if productIDStr == "" {
		// fmt.Fprintf(w, "err no product_id")
		JSONify(jsonReturn{
			StatusCode:   http.StatusBadRequest,
			Data:         nil,
			ErrorMessage: "err no product_id",
		}, w)
		return
	}

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		// fmt.Fprintf(w, "err found %s", err.Error())
		JSONify(jsonReturn{
			StatusCode:   http.StatusBadRequest,
			Data:         nil,
			ErrorMessage: err.Error(),
		}, w)
		return
	}

	productData, err := GetProduct(productID)
	if err != nil {
		// fmt.Fprintf(w, "err found %s", err.Error())
		JSONify(jsonReturn{
			StatusCode:   http.StatusInternalServerError,
			Data:         nil,
			ErrorMessage: err.Error(),
		}, w)
		return
	}

	// fmt.Fprintf(w, "%+v", productData)
	JSONify(jsonReturn{
		StatusCode: http.StatusOK,
		Data:       productData,
	}, w)
	return
}

func JSONify(jr jsonReturn, w http.ResponseWriter) {
	result, _ := json.Marshal(jr)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(jr.StatusCode)
	fmt.Fprint(w, string(result))
}
