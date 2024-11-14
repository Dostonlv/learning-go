package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func Logic(ctx context.Context, Info string) (string, error) {
	return "", nil
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}

	data := r.PostFormValue("data")
	res, err := Logic(ctx, data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return

	}
	_, err = w.Write([]byte(res))
	if err != nil {
		return
	}
}

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/"+data, nil)
	if err != nil {
		return "", err
	}
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	id, err := processResponse(resp.Body)
	return id, err
}

func processResponse(body io.ReadCloser) (string, error) {
	return "", nil
}
func main() {
	ctx := context.Background()
	result, err := Logic(ctx, "a string")
	fmt.Println(result, err)
}
