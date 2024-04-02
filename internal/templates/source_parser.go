package templates

const SOURCE_PARSER_TEMPLATE = `package main

import (
	"io"

	"github.com/tardigodev/tardigo-core/pkg"
	"github.com/tardigodev/tardigo-core/pkg/constants"
	"github.com/tardigodev/tardigo-core/pkg/dtypes"
	"github.com/tardigodev/tardigo-core/pkg/objects"
)

type sourceParserPlugin struct {
	SampleConfig string
}

func (pp sourceParserPlugin) GetRecord(reader io.Reader, readerDetail objects.ReaderDetail, addRecord pkg.AddRecord) error {
	return nil
}

func (pp sourceParserPlugin) GetSchema(reader io.Reader, readerDetail objects.ReaderDetail) (dtypes.Schema, error) {
	return nil, nil
}

func (pp sourceParserPlugin) GetPluginDetail() objects.PluginDetail {
	return objects.PluginDetail{
		PluginName: "template_source_parser",
		PluginType: constants.PluginTypeSourceParser,
	}
}

var SourceParserPlugin = sourceParserPlugin{
	SampleConfig: "defaults",
}
`
