module main

require aph-go-service-master/transport v0.0.0

require aph-go-service-master/datastruct v0.0.0

require aph-go-service-master/logging v0.0.0

require aph-go-service-master/service v0.0.0

require (
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.10.3 // indirect
)

replace aph-go-service-master/datastruct => ./datastruct

replace aph-go-service-master/logging => ./logging

replace aph-go-service-master/service => ./service

replace aph-go-service-master/transport => ./transport

go 1.17
