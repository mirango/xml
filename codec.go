package xml

import (
	"encoding/xml"
	"io"

	"github.com/mirango/framework"
)

type XML struct{}

func New() *XML {
	return &XML{}
}

func (XML) Encode(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (XML) EncodeTo(w io.Writer, v interface{}) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(v)
}

func (XML) EncodingType() string {
	return framework.MIME_XML
}

func (XML) Decode(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func (XML) DecodeFrom(r io.Reader, v interface{}) error {
	dec := xml.NewDecoder(r)
	return dec.Decode(v)
}

func (XML) DecodingType() string {
	return framework.MIME_XML
}

func (XML) MIMEType() string {
	return framework.MIME_XML
}

type XMLIndented struct{}

func NewIndented() *XMLIndented {
	return &XMLIndented{}
}

func (XMLIndented) Encode(v interface{}) ([]byte, error) {
	return xml.MarshalIndent(v, "", " ")
}

func (XMLIndented) EncodeTo(w io.Writer, v interface{}) error {
	b, err := xml.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func (XMLIndented) EncodingType() string {
	return framework.MIME_XML
}

func (XMLIndented) Decode(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func (XMLIndented) DecodeFrom(r io.Reader, v interface{}) error {
	dec := xml.NewDecoder(r)
	return dec.Decode(v)
}

func (XMLIndented) DecodingType() string {
	return framework.MIME_XML
}

func (XMLIndented) MIMEType() string {
	return framework.MIME_XML
}
