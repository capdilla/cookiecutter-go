package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/* Logger based on zap logger
 * How to use:
 * 1.- Run "Init(serviceName, mode)" function in order to setup the logger and instantiate into
 * 	a global variable.
 * 2.- Depending on whether the performance is critical for your task:
 * 	2.1.- Critical: Run "Logger()" in order to retrieve a well configured logger, you will get
 * 		a *zap.Logger so you must provide the type of each param to log, this way will provide
 *		the best performance at moment of logging.
 * 		Aditionally you can use "zap.L()" but you must run previously the step #1, if you don't
 * 		your logs won't beproperly printed.
 *	2.2.- Nice but not critical: Run "SLogger()" in order to get a "zap.SugaredLoger", which one
 *		has a lightly less performance but you won't need pass the type per each logged param.
 *
 * See https://github.com/uber-go/zap for more information.
 */

var (
	logger  *zap.Logger
	slogger *zap.SugaredLogger
)

// The next constants represents the two differents modes of logging that we will supporting.
const (
	// release mode will be used for production logging, so we will get structured logs in
	// order to support machine analizys.
	release = "release"
	// debug mode will be used on development environment, so we will get unstructured logs
	debug = "debug"
)

// InitLogger will setup & init the logger, the enconding of the logs will be determinated by
// the "mode", we will support only "release" and "debug" modes
func InitLogger() {
	mode := os.Getenv("GIN_MODE")

	// Setup logger
	switch mode {
	case release:
		logger, _ = zap.Config{
			Encoding:    "json",
			Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
			OutputPaths: []string{"stdout"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}.Build()
	default:
		logger, _ = zap.Config{
			Encoding:    "console",
			Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
			OutputPaths: []string{"stdout"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalColorLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}.Build()
	}

	// Setup microservice name
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "merchant-auth"
	}

	logger = logger.With(
		zap.String("service", serviceName),
	)

	// Replace global logger of zap, so we can use "zap.L()"
	zap.ReplaceGlobals(logger)

	// Create Sugared Logger, used when performance is nice but not critical
	slogger = logger.Sugar()

	slogger.Infof("logger initialized for %s", mode)
}

// Logger will return a *zap.Logger, use it when the performance is critical
func Logger() *zap.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

// SugaredLogger will return a *zap.SugaredLogger (allow to pass argument without cast), use it when the performance is nice but not critical
func SugaredLogger() *zap.SugaredLogger {
	if slogger == nil {
		InitLogger()
	}
	return slogger
}
