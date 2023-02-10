package main

import (
	h "github.com/oddball707/trainingCalendar/handler"
	m "github.com/oddball707/trainingCalendar/model"
	s "github.com/oddball707/trainingCalendar/service"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"net/http"
)

func main() {
	// initialize service layer
	srv := s.NewService()
	hnd := h.NewHandler(srv)
	router := h.NewRouter()

	mux := http.NewServeMux()

	f := func(w http.ResponseWriter,  r *http.Request) {
		router.ServeHTTP(w, r)
	}

	mux.HandleFunc("/", f)

	lambda.Start(httpadapter.New(mux).ProxyWithConext)
}