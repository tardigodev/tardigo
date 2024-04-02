package templates

const TARGET_PARSER_TEMPLATE = `package main

import (
	"io"

	"github.com/tardigodev/tardigo-core/pkg"
	"github.com/tardigodev/tardigo-core/pkg/constants"
	"github.com/tardigodev/tardigo-core/pkg/dtypes"
	"github.com/tardigodev/tardigo-core/pkg/objects"
)

type targetParserPlugin struct {
	SampleConfig string
}

func (tp targetParserPlugin) PutRecord(writer io.Writer, writerDetail objects.WriterDetail, record any, recordDetail objects.RecordDetail, schema dtypes.Schema, addErrorRecord pkg.AddRecord) error {
	return nil
}

func (tp targetParserPlugin) ConvertSchema(schema dtypes.Schema) (dtypes.Schema, error) {
	return nil, nil
}

func (tp targetParserPlugin) GetPluginDetail() objects.PluginDetail {
	return objects.PluginDetail{
		PluginName: "template_target_target",
		PluginType: constants.PluginTypeTargetParser,
	}
}

var TargetParserPlugin = targetParserPlugin{
	SampleConfig: "defaults",
}
`
