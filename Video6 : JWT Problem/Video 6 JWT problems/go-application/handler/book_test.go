package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetBooks(t *testing.T) {
	type args struct {
		in *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetBooks(tt.args.in)
		})
	}
}

func TestHandler_SaveBook(t *testing.T) {
	type args struct {
		in *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.SaveBook(tt.args.in)
		})
	}
}

func TestHandler_GetBookByID(t *testing.T) {
	type args struct {
		in *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetBookByID(tt.args.in)
		})
	}
}

func TestHandler_DeleteBookByID(t *testing.T) {
	type args struct {
		in *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.DeleteBookByID(tt.args.in)
		})
	}
}

func TestHandler_UpdateBookByID(t *testing.T) {
	type args struct {
		in *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.UpdateBookByID(tt.args.in)
		})
	}
}
