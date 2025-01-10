package cmd

import "system-Info-collector/pkg/web"

func StartSystemAgent() {
	//web server start
	web.StartWebServer()
}
