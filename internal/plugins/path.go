package plugins

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tardigodev/tardigo-core/pkg/constants"
)

type PluginDirType int

const (
	PluginBuildDirType PluginDirType = iota
	PluginInstallDirType
)

const (
	fmtBuildDir   = ".build/%s/"
	fmtInstallDir = ".tardigo/plugins/%s/%s/"

	fmtGoFile  = "%s.go"
	soFile     = "plugin.so"
	configFile = "config.json"
	detailFile = "detail.json"
)

func GetGoFilePath(pluginType constants.PluginType) string {
	return filepath.Join("./", fmt.Sprintf(fmtGoFile, pluginType))
}

func GetSoFileBuildPath(pluginType constants.PluginType) string {
	return filepath.Join(GetBuildDir(pluginType), soFile)
}

func GetSoFileInstallPath(pluginType constants.PluginType, pluginName string) string {
	return filepath.Join(GetInstallDir(pluginType, pluginName), soFile)
}

func GetConfigBuildPath(pluginType constants.PluginType) string {
	return filepath.Join(GetBuildDir(pluginType), configFile)
}

func GetConfigInstallPath(pluginType constants.PluginType, pluginName string) string {
	return filepath.Join(GetInstallDir(pluginType, pluginName), configFile)
}

func GetDetailBuildPath(pluginType constants.PluginType) string {
	return filepath.Join(GetBuildDir(pluginType), detailFile)
}

func GetDetailInstallPath(pluginType constants.PluginType, pluginName string) string {
	return filepath.Join(GetInstallDir(pluginType, pluginName), detailFile)
}

func GetBuildDir(pluginType constants.PluginType) string {
	return fmt.Sprintf(fmtBuildDir, pluginType)
}

func GetInstallDir(pluginType constants.PluginType, pluginName string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, fmt.Sprintf(fmtInstallDir, pluginName, pluginType))
}
