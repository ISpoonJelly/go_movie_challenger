go install ./models ./controllers
govendor update +l
go clean
go build