package ports

import log "github.com/sirupsen/logrus"

// ILogger creates a PORT for the logger
type ILogger interface {
	MakeLogger(filename string, display bool) *log.Logger
	SetFormater()
	Hook() *log.Logger
}
