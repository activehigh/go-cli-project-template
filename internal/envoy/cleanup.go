package envoy

import (
	"time"

	"github.com/imroc/req/v3"
	"go.uber.org/zap"
)

const envoyQuitURL = "http://127.0.0.1:15000/quitquitquit"
const quitTimeoutSeconds = 30

// Cleanup calls the Envoy quit endpoint when the application is ready to shut down.
func Cleanup() error {
	_, err := req.SetTimeout(time.Second * quitTimeoutSeconds).R().Post(envoyQuitURL)

	if err != nil {
		zap.S().Error("Error creating Envoy quit request: %w", zap.Error(err))
		return err
	}

	zap.S().Debug("Envoy quit request successful")

	return nil
}
