package main

import (
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/dhuruvah-apps/BuildSMEs-Api/config"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/server"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/db/redis"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/db/sqlite"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/logger"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/utils"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// @title BuildSMEs REST API
// @version 1.0
// @description Golang REST API for BuildSMEs - A construction management platform for SMEs
// @contact.name Dhuruvah Apps
// @contact.url https://github.com/dhuruvah-apps/BuildSMEs-Api
// @contact.email jaganathan.eswaran@gmail.com
// @BasePath /api/v1
func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	sqliteDB, err := sqlite.NewSqliteDB(cfg)
	if err != nil {
		appLogger.Fatalf("Sqlite init: %s", err)
	} else {
		appLogger.Infof("Sqlite connected, Status: %#v", sqliteDB.Stats())
	}
	defer sqliteDB.Close()

	redisClient := redis.NewRedisClient(cfg, appLogger)
	if redisClient == nil {
		appLogger.Errorf("Redis Client init: %s", err)
	}
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	// awsClient, err := aws.NewAWSClient(cfg.AWS.Endpoint, cfg.AWS.MinioAccessKey, cfg.AWS.MinioSecretKey, cfg.AWS.UseSSL)
	// if err != nil {
	// 	appLogger.Errorf("AWS Client init: %s", err)
	// }
	// appLogger.Info("AWS S3 connected")

	jaegerCfgInstance := jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           cfg.Jaeger.LogSpans,
			LocalAgentHostPort: cfg.Jaeger.Host,
		},
	}

	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)
	if err != nil {
		log.Fatal("cannot create tracer", err)
	}
	appLogger.Info("Jaeger connected")

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	appLogger.Info("Opentracing connected")

	s := server.NewServer(cfg, sqliteDB, redisClient, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
