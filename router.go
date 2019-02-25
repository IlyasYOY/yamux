package router

import (
	"fmt"
	"net/http"
	"strings"
)

func (mux *YAMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer mux.recover(writer, request)

	path := request.URL.Path

	switch request.Method {
	case http.MethodGet:
		for _, handle := range mux.getHandles {
			if strings.HasPrefix(path, handle.Path) {
				handle.Handle(writer, request)
				return
			}
		}
	case http.MethodPost:
		for _, handle := range mux.postHandles {
			if strings.HasPrefix(path, handle.Path) {
				handle.Handle(writer, request)
				return
			}
		}
	case http.MethodPut:
		for _, handle := range mux.putHandles {
			if strings.HasPrefix(path, handle.Path) {
				handle.Handle(writer, request)
				return
			}
		}
	case http.MethodDelete:
		for _, handle := range mux.deleteHandles {
			if strings.HasPrefix(path, handle.Path) {
				handle.Handle(writer, request)
				return
			}
		}
	case http.MethodHead:
		return
	}
	mux.DefaultHandle(writer, request)
}

func (mux *YAMux) Get(path string, handle Handle) *YAMux {
	checkIfExists(&mux.getHandles, path, http.MethodGet)
	addHandle(&mux.getHandles, path, handle)
	return mux
}

func (mux *YAMux) Post(path string, handle Handle) *YAMux {
	checkIfExists(&mux.postHandles, path, http.MethodPost)
	addHandle(&mux.postHandles, path, handle)
	return mux
}

func (mux *YAMux) Put(path string, handle Handle) *YAMux {
	checkIfExists(&mux.putHandles, path, http.MethodPut)
	addHandle(&mux.putHandles, path, handle)
	return mux
}

func (mux *YAMux) Delete(path string, handle Handle) *YAMux {
	checkIfExists(&mux.deleteHandles, path, http.MethodDelete)
	addHandle(&mux.deleteHandles, path, handle)
	return mux
}

func addHandle(handles *[]urlHandle, path string, handle Handle) {
	urlHandles := append(*handles, *NewUrlHandle(path, handle))
	handles = &urlHandles
}

func checkIfExists(handles *[]urlHandle, path string, method string) {
	for _, h := range *handles {
		if h.Path == path {
			panic(fmt.Sprintf("You registered handle with the same method (%s) and path (%s).", method, path))
		}
	}
}

func (mux *YAMux) recover(writer http.ResponseWriter, request *http.Request) {
	if rcv := recover(); mux.PanicHandle != nil && rcv != nil {
		mux.PanicHandle(writer, request)
	}
}
