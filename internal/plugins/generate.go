package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tardigodev/tardigo-core/pkg/constants"
	"github.com/tardigodev/tardigo/internal/templates"
)

func GenerateMetadata(pluginType constants.PluginType, pluginBuildPath string) error {
	log.Printf("generating metadata for plugin : %s from %s", pluginType, pluginBuildPath)
	plug, err := LoadGenericPlugin(pluginType, pluginBuildPath)
	if err != nil {
		return fmt.Errorf("failed to load plugin from path %s of type %s: %w", pluginBuildPath, pluginType, err)
	}

	configBuildPath := GetConfigBuildPath(pluginType)

	err = GenerateJSON(plug, configBuildPath)
	if err != nil {
		return fmt.Errorf("failed to generate config file: %w", err)
	}

	detailBuildPath := GetDetailBuildPath(pluginType)

	err = GenerateJSON(plug.GetPluginDetail(), detailBuildPath)
	if err != nil {
		return fmt.Errorf("failed to generate detail file: %w", err)
	}

	return nil
}

func GenerateJSON(item any, outputJsonPath string) error {
	jsonBytes, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(outputJsonPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}
	log.Printf("generated json at %s", outputJsonPath)
	return nil
}

func GenerateTemplate(pluginType constants.PluginType) {
	generateGoFile := func(content string) {
		goFilePath := GetGoFilePath(pluginType)
		if checkIfFileExists(goFilePath) {
			log.Fatalf("file %s already exists, delete it before generating template", goFilePath)
		}
		err := os.WriteFile(goFilePath, []byte(content), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	switch pluginType {
	case constants.PluginTypeSourceParser:
		generateGoFile(templates.SOURCE_PARSER_TEMPLATE)
	case constants.PluginTypeTargetParser:
		generateGoFile(templates.TARGET_PARSER_TEMPLATE)
	default:
		log.Fatalf("unknown plugin type %s", pluginType)
	}

	log.Printf("generated template for plugin : %s", pluginType)
}
