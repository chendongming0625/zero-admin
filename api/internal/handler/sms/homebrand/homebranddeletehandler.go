package handler

import (
	"net/http"

	"go-zero-admin/api/internal/logic/sms/homebrand"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HomeBrandDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteHomeBrandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHomeBrandDeleteLogic(r.Context(), ctx)
		resp, err := l.HomeBrandDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
