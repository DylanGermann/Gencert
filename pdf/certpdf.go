package pdf

import (
	"fmt"
	"os"
	"path"

	"gencert/cert"
	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: outputdir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Background
	background(pdf)

	// Header
	header(pdf, &cert)
	pdf.Ln(30)

	// Body
	body(pdf, &cert)

	// Footer
	footer(pdf)

	//save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	imageWidth := 50.0
	filename := "img/stamp.png"
	pageWidth, pageHeight := pdf.GetPageSize()
	x := pageWidth - imageWidth - 20.0
	y := pageHeight - imageWidth - 10.0
	pdf.ImageOptions(filename, x, y, imageWidth, 0, false, opts, 0, "")
}

func body(pdf *gofpdf.Fpdf, c *cert.Cert) {
	// Body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelPresented, "C")
	pdf.Ln(30)

	// Body - Student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, c.Name, "C")
	pdf.Ln(30)

	// Body - Participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelParticipation, "C")
	pdf.Ln(30)

	// Body - Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, c.LabelDate, "C")
	pdf.Ln(30)
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(filename, x+margin, 20, imageWidth, 0, false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename, x-margin, 20, imageWidth, 0, false, opts, 0, "")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png", 0, 0, pageWidth, pageHeight, false, opts, 0, "")
}
