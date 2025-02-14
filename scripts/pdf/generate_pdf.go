package pdf

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/andrefrco/resume/scripts/serve"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func GeneratePDF() {
	mux := serve.NewMux()

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to start temporary server: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	srv := &http.Server{Handler: mux}

	go func() {
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error in temporary server: %v", err)
		}
	}()
	defer srv.Shutdown(context.Background())

	maxRetries := 10
	waitTime := 100 * time.Millisecond

	serverReady := false
	for i := 0; i < maxRetries; i++ {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err == nil && resp.StatusCode == 200 {
			resp.Body.Close()
			serverReady = true
			break
		}
		time.Sleep(waitTime)
	}

	if !serverReady {
		log.Fatal("Server took too long to start")
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err = chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("http://localhost:%d", port)),
		chromedp.EmulateViewport(1280, 1024),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatal("Error generating PDF:", err)
	}

	if err := os.WriteFile("resume.pdf", buf, 0644); err != nil {
		log.Fatal("Error saving PDF:", err)
	}

	log.Println("PDF successfully generated: resume.pdf")
}
