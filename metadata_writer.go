package gotenberg

import "github.com/dcaraxes/gotenberg-go-client/document"

type WriteMetadataRequest struct {
	pdfs []document.Document

	*baseRequest
}

func NewWriteMetadataRequest(pdfs ...document.Document) *WriteMetadataRequest {
	return &WriteMetadataRequest{
		pdfs:        pdfs,
		baseRequest: newBaseRequest(),
	}
}

func (wmd *WriteMetadataRequest) endpoint() string {
	return "/forms/pdfengines/metadata/write"
}

func (wmd *WriteMetadataRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range wmd.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (wmd *WriteMetadataRequest) Metadata(md []byte) {
	wmd.fields[fieldMetadata] = string(md)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MainRequester(new(WriteMetadataRequest))
)