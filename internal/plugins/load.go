package plugins

import (
	"fmt"
	"plugin"

	"github.com/tardigodev/tardigo-core/pkg"
	"github.com/tardigodev/tardigo-core/pkg/constants"
)

func LoadGenericPlugin(pluginType constants.PluginType, pluginPath string) (pkg.Plugin, error) {
	switch pluginType {
	case constants.PluginTypeSourceParser:
		pluginSymbol, err := LoadSourceParserPlugin(pluginPath)
		if err != nil {
			return nil, err
		}
		return pluginSymbol, nil
	case constants.PluginTypeTargetParser:
		pluginSymbol, err := LoadTargetParserPlugin(pluginPath)
		if err != nil {
			return nil, err
		}
		return pluginSymbol, nil
	}
	return nil, fmt.Errorf("failed to load plugin from path %s of type %s", pluginPath, pluginType)
}

func LoadSourceParserPlugin(pluginPath string) (pkg.SourceParserPlugin, error) {
	plugin, err := loadPlugin(pluginPath)
	if err != nil {
		return nil, err
	}

	pluginSymbol, err := plugin.Lookup("SourceParserPlugin")
	if err != nil {
		return nil, err
	}

	if sourceParserPlugin, ok := pluginSymbol.(pkg.SourceParserPlugin); ok {
		return sourceParserPlugin, nil
	}
	return nil, fmt.Errorf("failed to load plugin")
}

func LoadTargetParserPlugin(pluginPath string) (pkg.TargetParserPlugin, error) {
	plugin, err := loadPlugin(pluginPath)
	if err != nil {
		return nil, err
	}

	pluginSymbol, err := plugin.Lookup("TargetParserPlugin")
	if err != nil {
		return nil, err
	}

	if targetParserPlugin, ok := pluginSymbol.(pkg.TargetParserPlugin); ok {
		return targetParserPlugin, nil
	}
	return nil, fmt.Errorf("failed to load plugin")
}

func loadPlugin(pluginPath string) (*plugin.Plugin, error) {
	plugin, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, err
	}
	return plugin, nil
}
