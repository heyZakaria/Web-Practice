package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			// TODO: Add test cases.
			name: "",
			want: "Hello Bro!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost:3456", nil)
			recorder := httptest.NewRecorder()
			HomeHandler(recorder, request)

			fmt.Println("Ana", request)
			fmt.Println("Ana 3awtani", recorder)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodystring := string(body)

			if !reflect.DeepEqual(bodystring, tt.want) {
				t.Error(response, bodystring, tt.want)
			}
		})
	}
}

func TestQueryHome(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		arg  args
		want string
	}{
		{
			// TODO: Add test cases.
			name: "success return response with name",
			arg: args{
				name: "Zakaria",
			},
			want: "Hello Zakaria",
		},
		{
			// TODO: Add test cases.
			name: "success return response without name",
			arg: args{
				name: "",
			},
			want: "Hello\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?name=%s", tt.arg.name), nil)
			recorder := httptest.NewRecorder()
			QueryHome(recorder,request)

			response := recorder.Result()
			fmt.Println("hhhh",response)
			body, _ := io.ReadAll(response.Body)
			bodystring := string(body)
			
			if !reflect.DeepEqual(bodystring, tt.want){
				t.Error(response,"3adna hada", bodystring, "Bghina hada",tt.want)
			}

		})
	}
}
