package converts

import (
    "bytes"
    "errors"
    "fmt"
    "io"
    "net/http"
    "regexp"
    "strconv"
    "strings"

    "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type ConvertRequest struct {
    Url string
}

var (
    verifyExtension = regexp.MustCompile("^.*?(\\.\\w{1,4})?$")
    AcceptsSuffix   = strings.Join([]string{".html", ".htm"}, ",")
    AcceptsType     = strings.Join([]string{"text/html"}, ",")
)

func (c ConvertRequest) ToPdf() ([]byte, error) {

    if err := verifySuffix(c.Url); err != nil {
        return nil, err
    }

    body, err := downloadAndVerify(c.Url)
    if err != nil {
        return nil, err
    }

    pdf, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        return nil, err
    }

    // Set global options
    pdf.Dpi.Set(600)
    pdf.Orientation.Set(wkhtmltopdf.OrientationPortrait)
    pdf.Grayscale.Set(true)
    pdf.NoCollate.Set(false)
    pdf.PageSize.Set(wkhtmltopdf.PageSizeA4)
    pdf.MarginBottom.Set(0)
    pdf.MarginLeft.Set(0)
    pdf.MarginTop.Set(0)
    pdf.MarginRight.Set(0)

    // Create a new input page from an URL
    page := wkhtmltopdf.NewPageReader(body)
    page.DisableSmartShrinking.Set(false)
    page.DisableExternalLinks.Set(true)
    page.DisableJavascript.Set(true)
    // Set options for this page
    page.Zoom.Set(0.95)

    // Add to document
    pdf.AddPage(page)

    outBuf := new(bytes.Buffer)
    pdf.SetOutput(outBuf)

    // Create PDF document in internal buffer
    err = pdf.Create()

    if err != nil {
        return nil, err
    }

    b := pdf.Bytes()
    if len(b) != 0 {
        return nil, errors.New(fmt.Sprintf("expected to have zero bytes in internal buffer, have %d", len(b)))
    }

    return outBuf.Bytes(), nil
}

func downloadAndVerify(url string) (io.ReadCloser, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    if v := convertOrDefault(resp.Header.Get("Content-Length"), 0); v <= 0 {
        return nil, errors.New("error or empty content-type")
    }

    if !strings.Contains(AcceptsType, resp.Header.Get("Content-Type")) {
        return nil, errors.New(fmt.Sprintf("content-type not accept %s", resp.Header.Get("Content-Type")))
    }

    return resp.Body, nil
}

func verifySuffix(url string) error {
    v := verifyExtension.FindAllStringSubmatch(url, -1)[0]
    if strings.Contains(AcceptsSuffix, v[1]) || v[1] == "" {
        return nil
    }
    return errors.New(fmt.Sprintf("not support extension %s", v[1]))
}

func convertOrDefault(str string, def int) int {
    v, err := strconv.Atoi(str)
    if err != nil {
        return def
    }
    return v
}
