package trace

import (
	"io"
	"t_task/client-api/config"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func InitGlobalTracer(env config.Config) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Reporter: &jaegercfg.ReporterConfig{
			LocalAgentHostPort: env.JaegerAgentHost + ":" + env.JaegerAgentPort,
		},
		Sampler: &jaegercfg.SamplerConfig{
			Type:  env.JaegerSamplerType,
			Param: env.JaegerSamplerParam,
		},
	}

	jLogger := jaegerlog.StdLogger
	// jMetricsFactory := jprom.New(jprom.WithRegisterer(prometheus.DefaultRegisterer))

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		"truesight-api",
		jaegercfg.Logger(jLogger),
		// jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}
	return closer, nil
}
