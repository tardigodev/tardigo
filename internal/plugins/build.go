package plugins

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/tardigodev/tardigo-core/pkg/constants"
)

func DetectPlugins() []constants.PluginType {
	detectedPlugins := []constants.PluginType{}

	addIfExist := func(pluginType constants.PluginType) {
		if checkIfFileExists(GetGoFilePath(pluginType)) {
			detectedPlugins = append(detectedPlugins, pluginType)
		}
	}

	addIfExist(constants.PluginTypeSourceParser)
	addIfExist(constants.PluginTypeTargetParser)
	addIfExist(constants.PluginTypeSourceReader)
	addIfExist(constants.PluginTypeTargetWriter)

	return detectedPlugins
}

func BuildPlugin(pluginFile string, pluginBuildPath string) error {
	log.Printf("building plugin : %s to %s", pluginFile, pluginBuildPath)
	_, err := runCommand("go", "build", "-trimpath", "-buildmode=plugin", "-o", pluginBuildPath, pluginFile)
	if err != nil {
		return fmt.Errorf("failed to build plugin %s to %s: %w", pluginFile, pluginBuildPath, err)
	}

	return nil
}

func VerifyPlugin(pluginType constants.PluginType, pluginBuildPath string) error {
	log.Printf("verifying plugin : %s from %s", pluginType, pluginBuildPath)
	_, err := LoadGenericPlugin(pluginType, pluginBuildPath)
	if err != nil {
		return err
	}
	return nil
}

func InstallPlugin(pluginType constants.PluginType, pluginName string) error {
	log.Printf("installing plugin : %s from %s", pluginType, pluginName)
	pluginBuildPath, pluginInstallPath := GetSoFileBuildPath(pluginType), GetSoFileInstallPath(pluginType, pluginName)
	if !checkIfFileExists(pluginBuildPath) {
		return fmt.Errorf("plugin does not exist at '%s', make sure to build the plugin first", pluginBuildPath)
	}

	configBuildPath, configInstallPath := GetConfigBuildPath(pluginType), GetConfigInstallPath(pluginType, pluginName)
	if !checkIfFileExists(configBuildPath) {
		return fmt.Errorf("config file does not exist at '%s', make sure to build the plugin first", configBuildPath)
	}

	detailBuildPath, detailInstallPath := GetDetailBuildPath(pluginType), GetDetailInstallPath(pluginType, pluginName)
	if !checkIfFileExists(detailBuildPath) {
		return fmt.Errorf("detail file does not exist at '%s',make sure to build the plugin first", detailBuildPath)
	}

	if err := os.MkdirAll(GetInstallDir(pluginType, pluginName), os.ModePerm); err != nil {
		return err
	}

	if err := os.Rename(pluginBuildPath, pluginInstallPath); err != nil {
		return err
	}

	if err := os.Rename(configBuildPath, configInstallPath); err != nil {
		return err
	}

	if err := os.Rename(detailBuildPath, detailInstallPath); err != nil {
		return err
	}

	if err := os.Remove(GetBuildDir(pluginType)); err != nil {
		return err
	}
	return nil
}

func GetModuleName() (string, error) {
	moduleName, err := runCommand("go", "list", "-m")
	if err != nil {
		return "", fmt.Errorf("failed to get module name: %w", err)
	}
	return strings.Trim(string(moduleName), "\n"), nil
}

func checkIfFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func runCommand(args ...string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)
	log.Printf("executing command : `%s`\n", cmd.String())

	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return out, err
	}

	return out, nil
}
