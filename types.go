package router

import (
	"fmt"
	"io"
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request)

type urlHandle struct {
	Path   string
	Handle Handle
}

func NewUrlHandle(path string, handle Handle) *urlHandle {
	return &urlHandle{
		Path:   path,
		Handle: handle,
	}
}

type YAMux struct {
	DefaultHandle Handle
	PanicHandle   Handle
	getHandles    []urlHandle
	postHandles   []urlHandle
	putHandles    []urlHandle
	deleteHandles []urlHandle
}

func NewYAMux() *YAMux {
	return &YAMux{
		DefaultHandle: defaultHandle,
		PanicHandle:   panicHandle,
	}
}

func defaultHandle(writer http.ResponseWriter, request *http.Request) {
	_, err := io.WriteString(writer, fmt.Sprintf("Response Status Code: %d, for method %s, at %s",
		http.StatusBadRequest, request.Method, request.URL.Path))
	if err != nil {
		panic("Unable to write response...")
	}
}

func panicHandle(writer http.ResponseWriter, request *http.Request) {
	_, err := io.WriteString(writer, fmt.Sprintf("Response Status Code: %d, for method %s, at %s",
		http.StatusBadRequest, request.Method, request.URL.Path))
	if err != nil {
		panic("Unable to write response...")
	}
}
