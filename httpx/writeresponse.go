package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

// WriteResponse - write a http.Response, utilizing the content, status code, and headers
// Content types supported: []byte, string, error, io.Reader, io.ReadCloser. Other types will be treated as JSON and serialized, if
// the headers content type is JSON. If not JSON, then an error will be raised.
func WriteResponse[E core.ErrorHandler](w http.ResponseWriter, headers any, statusCode int, content any, reqHeader http.Header) (contentLength int64) {
	var e E

	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	SetHeaders(w, headers)
	if content == nil {
		w.WriteHeader(statusCode)
		return 0
	}
	if len(w.Header().Get(ContentEncoding)) != 0 {
		reqHeader.Set(AcceptEncoding, "")
	}
	writer, status0 := io.NewEncodingWriter(w, reqHeader)
	if !status0.OK() {
		e.Handle(status0, core.RequestId(w.Header()))
		w.WriteHeader(status0.HttpCode())
		return 0
	}
	if writer.ContentEncoding() != io.NoneEncoding {
		w.Header().Add(ContentEncoding, writer.ContentEncoding())
	}
	w.WriteHeader(statusCode)
	contentLength, status0 = writeContent(writer, content, w.Header().Get(ContentType))
	_ = writer.Close()
	if !status0.OK() {
		e.Handle(status0, core.RequestId(w.Header()))
	}
	return contentLength
}

func CreateAcceptEncodingHeader() http.Header {
	out := make(http.Header)
	out.Add(AcceptEncoding, AcceptEncodingValue)
	return out
}

/*
	ct := GetContentType(headers)
	buf, status0 := core.Bytes(content, ct)
	if !status0.OK() {
		e.Handle(status0, status.RequestId(), writeLoc)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	SetHeaders(w, headers)
	if len(ct) == 0 {
		w.Header().Set(ContentType, http.DetectContentType(buf))
	}
	w.WriteHeader(status.Http())

*/
/*
	bytes, err := w.Write(buf)
	if err != nil {
		e.Handle(core.NewStatusError(http.StatusInternalServerError, writeLoc, err), status.RequestId(), "")
	}
	if bytes != len(buf) {
		e.Handle(core.NewStatusError(http.StatusInternalServerError, writeLoc, errors.New(fmt.Sprintf("error on ResponseWriter().Write() -> [got:%v] [want:%v]\n", bytes, len(buf)))), status.RequestId(), "")
	}

*/

/*
func writeStatusContent[E core.ErrorHandler](w http.ResponseWriter, status core.Status, location string) {
	var e E

	if status.Content() == nil {
		return
	}
	ct := status.ContentHeader().Get(ContentType)
	buf, status1 := core.Bytes(status.Content(), ct)
	if !status1.OK() {
		e.Handle(status, status.RequestId(), location+writeStatusContentLoc)
		return
	}
	if len(ct) == 0 {
		ct = http.DetectContentType(buf)
	}
	w.Header().Set(ContentType, ct)
	_, err := w.Write(buf)
	if err != nil {
		e.Handle(core.NewStatusError(http.StatusInternalServerError, location+writeStatusContentLoc, err), "", "")
	}
}


*/
