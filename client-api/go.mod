module t_task/client-api

go 1.12

require (
	github.com/gorilla/mux v1.7.1
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	google.golang.org/grpc v1.20.1
	t_task/proto v0.0.0-00010101000000-000000000000
)

replace t_task/proto => ../proto
