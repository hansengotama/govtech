package handler

import (
	"github.com/hansengotama/govtech/internal/lib/errcustom"
	"github.com/hansengotama/govtech/internal/service/product"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var queryParam product.GetProductsQueryParam

	products, err := product.GetProducts(r.Context(), queryParam)
	if err != nil {
		httpRes := constructErrResponse(errcustom.GetHTTPErrCode(err), err.Error())
		writeResponse(w, httpRes)
		return
	}

	httpRes := constructSuccessResp(products, http.StatusOK)
	writeResponse(w, httpRes)
}
