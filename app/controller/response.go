package controller

import (
	"fmt"
	"math"
	"net/http"
)

const (
	/* code response app */
	CODE_PREFIX = "PCFG"

	CODE_SUCCESS             = "PCFG-200"
	CODE_BAD_REQUEST         = "PCFG-400"
	CODE_INTERNAL_SERVER     = "PCFG-500"
	CODE_UNAUTHORIZED        = "401"
	CODE_UNAUTHORIZED_ACCESS = "PCFG-403"

	CODE_SUCCESS_MSG             = "Success"
	CODE_BAD_REQUEST_MSG         = "Bad Request"
	CODE_INTERNAL_SERVER_MSG     = "Internal Server Error"
	CODE_UNAUTHORIZED_MSG        = "Unauthorized"
	CODE_UNAUTHORIZED_ACCESS_MSG = "Unauthorized Access"
)

type (
	Response struct {
		ResponseCode    string `json:"response_code"`
		ResponseMessage string `json:"response_message"`
		Meta            Meta   `json:"debug_param"`
		TraceID         string `json:"trace_id"`
		Pagination
		Data interface{} `json:"data,omitempty"`
	}

	Pagination struct {
		TotalPage     float64 `json:"total_page,omitempty"`
		TotalRecord   int64   `json:"total_record,omitempty"`
		RecordPerPage int     `json:"record_per_page,omitempty"`
		PageNum       int     `json:"page_num,omitempty"`
	}

	Meta struct {
		DebugParam string `json:"debug_param"`
	}

	IResponse interface {
		Set(statusCode int, data interface{}, err error)
		SetPagination(page, limit int, totalCount int64)
		SetTraceID(traceID string)

		setDebugParam(err error)
		mappingCode(statusCode int)
	}
)

func (response *Response) SetTraceID(traceID string) {
	response.TraceID = traceID
}

func (response *Response) Set(statusCode int, data interface{}, err error) {
	// mapping response code
	response.mappingCode(statusCode)
	response.Data = data

	// mapping to debug param
	if err != nil {
		response.setDebugParam(err)
	}
}

func (response *Response) SetPagination(page, limit int, totalCount int64) {
	if limit != 0 {
		response.Pagination.TotalPage = math.Ceil(float64(totalCount) / float64(limit))
	}
	response.Pagination.TotalRecord = totalCount
	response.Pagination.RecordPerPage = limit
	response.Pagination.PageNum = page
}

func (response *Response) setDebugParam(err error) {
	response.Meta = Meta{
		DebugParam: fmt.Sprintf("%v", err),
	}
}

func (response *Response) mappingCode(statusCode int) {
	switch statusCode {
	case http.StatusInternalServerError:
		response.ResponseCode = CODE_INTERNAL_SERVER
		response.ResponseMessage = CODE_INTERNAL_SERVER_MSG
	case http.StatusBadRequest:
		response.ResponseCode = CODE_BAD_REQUEST
		response.ResponseMessage = CODE_BAD_REQUEST_MSG
	case http.StatusUnauthorized:
		response.ResponseCode = CODE_UNAUTHORIZED
		response.ResponseMessage = CODE_UNAUTHORIZED_MSG
	case http.StatusOK:
		response.ResponseCode = CODE_SUCCESS
		response.ResponseMessage = CODE_SUCCESS_MSG
	default:
		response.ResponseCode = CODE_INTERNAL_SERVER
		response.ResponseMessage = CODE_INTERNAL_SERVER_MSG
	}
}
