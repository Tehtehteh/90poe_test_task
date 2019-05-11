module t_task/pds

go 1.12

require (
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.3.0
	t_task/proto v0.0.0-00010101000000-000000000000
)

replace t_task/proto => ../proto
