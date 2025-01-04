package main

import (
	"system-Info-collector/pkg/logger"
	"system-Info-collector/pkg/web"
)

var (
	_buildDate, _buildRevision, _buildRevisionShort, _buildBranch, _buildTag string
)

func main() {
	fnc := "main"

	//build info print
	logger.Info(fnc, "%s: build info", "system-agent")
	logger.Info(fnc, "buildDate: %s", _buildDate)
	logger.Info(fnc, "buildRevision: %s (%s)", _buildRevisionShort, _buildRevision)
	logger.Info(fnc, "buildBranch: %s", _buildBranch)
	logger.Info(fnc, "buildTag: %s", _buildTag)

	//web server start
	web.StartWebServer()
}
