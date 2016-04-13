package xml

import (
	"encoding/xml"
	"io"

	"github.com/mirango/framework"
)

type Renderer struct{}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) MimeType() string {
	return framework.MIME_XML
}

func (r *Renderer) RenderToWriter(wr io.Writer, c framework.Context, data interface{}) error {
	b, err := r.Render(c, data)
	if err != nil {
		return err
	}
	_, err = wr.Write(b)
	return err
}

func (r *Renderer) Render(c framework.Context, data interface{}) ([]byte, error) {
	if c.IsIndented() {
		return xml.MarshalIndent(data, "", " ")
	}
	return xml.Marshal(data)
}
