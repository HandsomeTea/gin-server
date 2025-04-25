package services

import (
	"gin-server/server/globals"
	"net/http"
	"strings"
	"sync"
)

type requestServer interface {
	Get(path string) (any, *globals.HttpException)
	// 添加其他必要的方法
}

type httpClient struct {
	client  *http.Client
	baseURL string
}

func (c *httpClient) Get(path string) (any, *globals.HttpException) {
	print("调用到了封装的get方法")
	return "123", nil
}

func createHTTPClient(baseURL string) (*httpClient, *globals.HttpException) {
	if baseURL == "" {
		return nil, globals.NewException("create http client: baseURL cannot be empty")
	}
	return &httpClient{
		client:  &http.Client{},
		baseURL: strings.TrimSuffix(baseURL, "/"),
	}, nil
}

func createUserManagerServer() (*httpClient, *globals.HttpException) {
	baseURL := "http://localhost:8080"
	client, err := createHTTPClient(baseURL)

	if err != nil {
		return nil, err
	}
	return client, nil
}

type httpService struct {
	userManagerServer requestServer
}

func (c *httpService) CheckUserPermission() (any, *globals.HttpException) {
	if c == nil || c.userManagerServer == nil {
		return nil, globals.NewException("user manager request server is not initialized")
	}
	c.userManagerServer.Get("/users/1")

	return nil, nil
}

var (
	HTTP *httpService
	once sync.Once
)

func init() {
	once.Do(func() {
		server, err := createUserManagerServer()

		if err != nil {
			return
		}
		HTTP = &httpService{userManagerServer: server}
	})
}
