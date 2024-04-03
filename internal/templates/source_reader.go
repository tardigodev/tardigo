package templates

const SOURCE_READER_TEMPLATE = `package main

import (
	"io"

	"github.com/tardigodev/tardigo-core/pkg/constants"
	"github.com/tardigodev/tardigo-core/pkg/objects"
)

type sourceReaderPlugin struct {
	SampleConfig string
}

func (rp sourceReaderPlugin) GetReader(putReader func(io.Reader, objects.ReaderDetail) error) error {
	return nil
}

func (rp sourceReaderPlugin) GetPluginDetail() objects.PluginDetail {
	return objects.PluginDetail{
		PluginName: "template_source_reader",
		PluginType: constants.PluginTypeSourceReader,
	}
}

var SourceReaderPlugin = sourceReaderPlugin{
	SampleConfig: "defaults",
}
`
