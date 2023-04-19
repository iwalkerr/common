package respw

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type responseSuccessBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func success(data any) *responseSuccessBean {
	return &responseSuccessBean{uint32(OK), mapErrMsg(OK), data}
}

type responseErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func errorx(e *CodeError) *responseErrorBean {
	return &responseErrorBean{uint32(e.getCode()), e.getMsg()}
}

func Json(r *http.Request, w http.ResponseWriter, resp any, err error) {
	if err == nil {
		httpx.WriteJson(w, http.StatusOK, success(resp))
		return
	}

	logx.WithContext(r.Context()).Errorf("【API-ERR】=> %+v ", err)
	causeErr := errors.Cause(err)
	if e, ok := causeErr.(*CodeError); ok { //自定义错误类型
		httpx.WriteJson(w, http.StatusBadRequest, errorx(e))
		return
	}

	httpx.WriteJson(w, http.StatusBadRequest, errorx(NewCode(SERVER_COMMON_ERROR)))
}
