package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/talon-one/talang"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger := initLogger()
	port := initPort(logger)
	webRoot := initWebRoot(logger)
	defer logger.Sync()

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(webRoot))
	mux.Handle("/", fs)
	mux.HandleFunc("/healthcheck", healthCheck(logger))
	mux.HandleFunc("/api/parse", handleParse(logger))

	logger.Info("Talangland Server Starting", zap.Int64("Port", port), zap.String("WebRoot", webRoot))

	var httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Panic("http server closed with error", zap.Error(err))
	}
}

func initLogger() *zap.Logger {
	logLevel := zap.InfoLevel
	logLevelEnv := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	if logLevelEnv == "DEBUG" {
		logLevel = zap.DebugLevel
	}

	config := zap.NewDevelopmentEncoderConfig()
	config.MessageKey = "log"

	return zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), logLevel))
}

func initPort(logger *zap.Logger) int64 {
	port := int64(9876)
	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		maybePort, err := strconv.ParseInt(portEnv, 10, 16)
		if err != nil || port <= 0 || port > 65535 {
			logger.Panic("Environment variable PORT is not valid")
		}
		port = maybePort
	}

	return port
}

func initWebRoot(logger *zap.Logger) string {
	webRoot := "static"
	webRootEnv := os.Getenv("WEB_ROOT")
	if webRootEnv != "" {
		webRoot = webRootEnv
	}

	absRoot, err := filepath.Abs(webRoot)
	if err != nil {
		logger.Panic("Couldn't initiate absolute web root", zap.Error(err))
	}

	return absRoot
}

func handleParse(logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		talangParam, ok := values["talang"]
		if !ok {
			http.Error(w, "No talang parameter received", http.StatusBadRequest)
			return
		}

		input := talangParam[0]
		interp := talang.MustNewInterpreter()

		parsed, err := talang.Parse(input)
		if err != nil {
			errorText := err.Error()
			logger.Error("Failed to parse input talang", zap.String("talang", input), zap.String("Error", errorText))
			http.Error(w, fmt.Sprintf("talang couldn't be parsed: %s", errorText), http.StatusBadRequest)
			return
		}

		if err := interp.Evaluate(parsed); err != nil {
			errorText := err.Error()
			logger.Error("Failed to evaluate input talang", zap.String("talang", input), zap.String("Error", errorText))
			http.Error(w, fmt.Sprintf("talang couldn't be evaluated: %s", errorText), http.StatusInternalServerError)
		}

		output := parsed.Stringify()
		logger.Debug(fmt.Sprintf("input: %s ==> output: %s || type: %s", input, output, parsed.Kind.String()))

		_, err = w.Write([]byte(output))
		if err != nil {
			logger.Error("Unexpected error while writing output", zap.String("Error", err.Error()))
			http.Error(w, "Unexpected Internal Error", http.StatusInternalServerError)
			return
		}
	}
}

func healthCheck(logger *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("health-check request received")
		w.WriteHeader(http.StatusOK)
		return
	}
}
