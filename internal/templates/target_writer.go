package templates

const TARGET_WRITER_TEMPLATE = `package main
import (
	"io"

	"github.com/tardigodev/tardigo-core/pkg/constants"
	"github.com/tardigodev/tardigo-core/pkg/objects"
)

type targetWriterPlugin struct {
	SampleConfig string
}

func (tp targetWriterPlugin) GetWriter(putWriter func(io.Writer, objects.WriterDetail) error) error {
	return nil
}

func (rp targetWriterPlugin) GetPluginDetail() objects.PluginDetail {
	return objects.PluginDetail{
		PluginName: "template_target_writer",
		PluginType: constants.PluginTypeTargetWriter,
	}
}

var TargetWriterPlugin = targetWriterPlugin{
	SampleConfig: "defaults",
}
`
