package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tardigodev/tardigo-core/pkg/constants"
	"github.com/tardigodev/tardigo/internal/plugins"
)

var SUPPORTED_PLUGINS = []constants.PluginType{
	constants.PluginTypeSourceStorage,
	constants.PluginTypeTargetStorage,
	constants.PluginTypeSourceParser,
	constants.PluginTypeTargetParser,
	constants.PluginTypeProcessor,
}

// pluginCmd represents the `plugin` command
var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "build, install or manage your tardigo plugin project",
	Long:  ``,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		pluginType := cmd.Flag("pluginType").Value.String()
		if pluginType != "" {
			var isSupported bool
			for _, plugin := range SUPPORTED_PLUGINS {
				if pluginType == string(plugin) {
					isSupported = true
				}
			}
			if !isSupported {
				return fmt.Errorf("unsupported plugin type")
			}
		}
		return nil
	},
}

const (
	ErrorDetectionFailed    = "expected atleast one of - source_parser.go, target_parser.go, source_storage.go, target_storage.go, processor.go"
	ErrorBuildFailed        = "make sure the provided plugin file is valid and you have `go` installed in your system"
	ErrorVerificationFailed = "make sure your plugin file implements the required methods for that plugin"
)

// buildCmd represents the `plugin build“ command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build your plugin project",
	Long: `
	Detect, build and verify your current tardigo plugin project.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		pluginType := cmd.Flag("pluginType").Value.String()
		shouldInstall := cmd.Flag("install").Value.String() == "true"
		var detectedPlugins []constants.PluginType
		if pluginType == "" {
			detectedPlugins = plugins.DetectPlugins()
			if len(detectedPlugins) == 0 {
				log.Fatalf("No plugin files detected in current directory - %s\n", ErrorDetectionFailed)
				os.Exit(1)
			}
		} else {
			detectedPlugins = append(detectedPlugins, constants.PluginType(pluginType))
		}

		for _, pluginType := range detectedPlugins {
			pluginBuildPath, pluginGoPath := plugins.GetSoFileBuildPath(pluginType), plugins.GetGoFilePath(pluginType)
			err := plugins.BuildPlugin(pluginGoPath, pluginBuildPath)
			if err != nil {
				log.Fatalf("failed to build plugin %s: %s - %s\n", pluginType, err, ErrorBuildFailed)
				os.Exit(1)
			}

			err = plugins.VerifyPlugin(pluginType, pluginBuildPath)
			if err != nil {
				log.Fatalf("failed to verify plugin %s: %s - %s\n", pluginType, err, ErrorVerificationFailed)
				os.Exit(1)
			}

			err = plugins.GenerateMetadata(pluginType, pluginBuildPath)
			if err != nil {
				log.Fatalf("failed to generate metadata for plugin %s: %s\n", pluginType, err)
				os.Exit(1)
			}

			if shouldInstall {
				pluginName, err := plugins.GetModuleName()
				if err != nil {
					log.Fatalf("failed to get plugin name: %s\n", err)
					os.Exit(1)
				}

				err = plugins.InstallPlugin(pluginType, pluginName)
				if err != nil {
					log.Fatalf("failed to install plugin %s: %s\n", pluginType, err)
					os.Exit(1)
				}
			}
		}
	},
}

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "generate template of a plugin",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if pluginType := cmd.Flag("pluginType").Value.String(); pluginType == "" {
			return fmt.Errorf("--pluginType is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pluginType := constants.PluginType(cmd.Flag("pluginType").Value.String())
		plugins.GenerateTemplate(pluginType)
	},
}

func init() {
	pluginCmd.PersistentFlags().StringP("pluginType", "t", "", fmt.Sprintf("supported plugin types : %s", SUPPORTED_PLUGINS))
	buildCmd.Flags().BoolP("install", "i", false, "install plugin after building")

	rootCmd.AddCommand(pluginCmd)
	pluginCmd.AddCommand(buildCmd)
	pluginCmd.AddCommand(templateCmd)
}
