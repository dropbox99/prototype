package controller

import (
	"errors"
	"reflect"
	"testing"
)

func TestResponse_Set(t *testing.T) {
	type fields Response
	type args struct {
		statusCode int
		err        error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result *Response
	}{
		{
			name:   "success if set 200",
			fields: fields{},
			args: args{
				statusCode: 200,
				err:        nil,
			},
			result: &Response{
				ResponseCode:    CODE_SUCCESS,
				ResponseMessage: CODE_SUCCESS_MSG,
			},
		},
		{
			name:   "success if set 400",
			fields: fields{},
			args: args{
				statusCode: 400,
				err:        nil,
			},
			result: &Response{
				ResponseCode:    CODE_BAD_REQUEST,
				ResponseMessage: CODE_BAD_REQUEST_MSG,
			},
		},
		{
			name:   "success if set 500",
			fields: fields{},
			args: args{
				statusCode: 500,
				err:        nil,
			},
			result: &Response{
				ResponseCode:    CODE_INTERNAL_SERVER,
				ResponseMessage: CODE_INTERNAL_SERVER_MSG,
			},
		},
		{
			name:   "success if set 401",
			fields: fields{},
			args: args{
				statusCode: 401,
				err:        nil,
			},
			result: &Response{
				ResponseCode:    CODE_UNAUTHORIZED,
				ResponseMessage: CODE_UNAUTHORIZED_MSG,
			},
		},
		{
			name:   "success if set default",
			fields: fields{},
			args: args{
				statusCode: 20,
				err:        nil,
			},
			result: &Response{
				ResponseCode:    CODE_INTERNAL_SERVER,
				ResponseMessage: CODE_INTERNAL_SERVER_MSG,
			},
		},
		{
			name:   "error if there is an error 500",
			fields: fields{},
			args: args{
				statusCode: 500,
				err:        errors.New("there is error"),
			},
			result: &Response{
				ResponseCode:    CODE_INTERNAL_SERVER,
				ResponseMessage: CODE_INTERNAL_SERVER_MSG,
				Meta: Meta{
					DebugParam: "there is error",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := &Response{
				ResponseCode:    tt.fields.ResponseCode,
				ResponseMessage: tt.fields.ResponseMessage,
				Meta:            tt.fields.Meta,
				TraceID:         tt.fields.TraceID,
				Pagination:      tt.fields.Pagination,
			}
			response.Set(tt.args.statusCode, nil, tt.args.err)

			if !reflect.DeepEqual(response, tt.result) {
				t.Errorf("response.Set = %v, want %v", response, tt.result)
			}
		})
	}
}

func TestResponse_SetPagination(t *testing.T) {
	type fields struct {
		ResponseCode    string
		ResponseMessage string
		Meta            Meta
		TraceID         string
		Pagination      Pagination
	}
	type args struct {
		page       int
		limit      int
		totalCount int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result *Response
	}{
		{
			name:   "success",
			fields: fields{},
			args: args{
				page:       1,
				limit:      10,
				totalCount: 15,
			},
			result: &Response{
				Pagination: Pagination{
					TotalPage:     2,
					TotalRecord:   15,
					RecordPerPage: 10,
					PageNum:       1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := &Response{
				ResponseCode:    tt.fields.ResponseCode,
				ResponseMessage: tt.fields.ResponseMessage,
				Meta:            tt.fields.Meta,
				TraceID:         tt.fields.TraceID,
				Pagination:      tt.fields.Pagination,
			}
			response.SetPagination(tt.args.page, tt.args.limit, tt.args.totalCount)

			if !reflect.DeepEqual(response, tt.result) {
				t.Errorf("response.SetPagination = %v, want %v", response, tt.result)
			}
		})
	}
}

func TestResponse_setDebugParam(t *testing.T) {
	type fields struct {
		ResponseCode    string
		ResponseMessage string
		Meta            Meta
		TraceID         string
		Pagination      Pagination
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result *Response
	}{
		{
			name:   "success",
			fields: fields{},
			args: args{
				err: errors.New("there is an error"),
			},
			result: &Response{
				Meta: Meta{
					DebugParam: "there is an error",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := &Response{
				ResponseCode:    tt.fields.ResponseCode,
				ResponseMessage: tt.fields.ResponseMessage,
				Meta:            tt.fields.Meta,
				TraceID:         tt.fields.TraceID,
				Pagination:      tt.fields.Pagination,
			}
			response.setDebugParam(tt.args.err)

			if !reflect.DeepEqual(response, tt.result) {
				t.Errorf("response.setDebugParam = %v, want %v", response, tt.result)
			}
		})
	}
}

func TestResponse_mappingCode(t *testing.T) {
	type fields struct {
		ResponseCode    string
		ResponseMessage string
		Meta            Meta
		TraceID         string
		Pagination      Pagination
	}
	type args struct {
		statusCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result *Response
	}{
		{
			name:   "success if set 200",
			fields: fields{},
			args: args{
				statusCode: 200,
			},
			result: &Response{
				ResponseCode:    CODE_SUCCESS,
				ResponseMessage: CODE_SUCCESS_MSG,
			},
		},
		{
			name:   "success if set 400",
			fields: fields{},
			args: args{
				statusCode: 400,
			},
			result: &Response{
				ResponseCode:    CODE_BAD_REQUEST,
				ResponseMessage: CODE_BAD_REQUEST_MSG,
			},
		},
		{
			name:   "success if set 500",
			fields: fields{},
			args: args{
				statusCode: 500,
			},
			result: &Response{
				ResponseCode:    CODE_INTERNAL_SERVER,
				ResponseMessage: CODE_INTERNAL_SERVER_MSG,
			},
		},
		{
			name:   "success if set 401",
			fields: fields{},
			args: args{
				statusCode: 401,
			},
			result: &Response{
				ResponseCode:    CODE_UNAUTHORIZED,
				ResponseMessage: CODE_UNAUTHORIZED_MSG,
			},
		},
		{
			name:   "success if set default",
			fields: fields{},
			args: args{
				statusCode: 20,
			},
			result: &Response{
				ResponseCode:    CODE_INTERNAL_SERVER,
				ResponseMessage: CODE_INTERNAL_SERVER_MSG,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := &Response{
				ResponseCode:    tt.fields.ResponseCode,
				ResponseMessage: tt.fields.ResponseMessage,
				Meta:            tt.fields.Meta,
				TraceID:         tt.fields.TraceID,
				Pagination:      tt.fields.Pagination,
			}
			response.mappingCode(tt.args.statusCode)

			if !reflect.DeepEqual(response, tt.result) {
				t.Errorf("response.mappingCode = %v, want %v", response, tt.result)
			}
		})
	}
}
