package main

import (
	"github.com/Snap-for-Windows/snap-plugin-collector-active-directory/activedirectory"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	pluginName    = "activedirectory-collector"
	pluginVersion = 1
)

//plugin bootstrap
func main() {
	plugin.StartCollector(
		activedirectory.ActiveDirectoryCollector{},
		pluginName,
		pluginVersion)
}
