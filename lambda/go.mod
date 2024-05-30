module trainingCalendar/lambda

go 1.22

require (
	github.com/aws/aws-lambda-go v1.37.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.14.0
	github.com/oddball707/trainingCalendar v1.0.0
)

require (
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/jordic/goics v0.0.0-20210404174824-5a0337b716a0 // indirect
)

replace (
	github.com/oddball707/trainingCalendar/handler => ../handler
	github.com/oddball707/trainingCalendar/server => ../server
)
