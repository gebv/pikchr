package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gebv/pikchr"
	"github.com/pkg/errors"
)

func main() {
	http.Handle("/", &renderServer{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type renderServer struct{}

func (s *renderServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tn := time.Now()
	execDuration := func() {
		execMS := time.Since(tn).Milliseconds()
		w.Header().Set("x-exec-duration-ms", strconv.Itoa(int(execMS)))
	}

	w.Header().Set("server", "PikchrRenderServer")

	if r.Method == http.MethodGet {
		responseWelcomeText(w)
		return
	}

	if r.Method != http.MethodPost {
		responseErr(w, http.StatusMethodNotAllowed, errors.New("allowd only POST method"))
		return
	}

	in := generateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		responseErr(w, http.StatusMethodNotAllowed, errors.Wrap(err, "failed decode body"))
		return
	}

	if in.DiagramSrc == "" {
		responseErr(w, http.StatusBadRequest, fmt.Errorf("invalid arguments: empty src of diagram"))
		return
	}

	opts := []pikchr.Option{}
	if in.Dark {
		opts = append(opts, pikchr.Dark())
	}
	if in.ClassName != nil && *in.ClassName != "" {
		opts = append(opts, pikchr.SVGClass(*in.ClassName))
	}

	res, ok := pikchr.Render(in.DiagramSrc, opts...)
	if !ok {
		responseErr(w, http.StatusOK, fmt.Errorf("failed render: %v", res.Data))
		return
	}

	svgB64 := stob64(res.Data)
	svgImg := fmt.Sprintf(`<img height="%d" width="%d" src="data:image/svg+xml;base64,%s"></img>`, res.Height, res.Width, svgB64)

	execDuration()
	renderJSON(w, http.StatusOK, map[string]interface{}{
		"success":        true,
		"svg_raw":        res.Data,
		"svg_raw_base64": svgB64,
		"img_inline_svg": svgImg,
		"svg_width":      res.Width,
		"svg_height":     res.Height,
	})
}

type generateRequest struct {
	DiagramSrc string  `json:"diagram_src"`
	Dark       bool    `json:"dark,omitempty"`
	ClassName  *string `json:"class_name,omitempty"`
}

func renderJSON(w http.ResponseWriter, status int, dat interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(dat)
	if err != nil {
		log.Println("failed encode json:", err)
	}
}

func responseErr(w http.ResponseWriter, status int, err error) {
	if err == nil {
		return
	}
	renderJSON(w, status, map[string]interface{}{
		"success": false,
		"err":     err.Error(),
	})
}

func responseWelcomeText(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "<p>This is render server of pichkr diagrams</p>")
	fmt.Fprintln(w, "<p>More deatils to follow link</p>")
	fmt.Fprintln(w, "<a target=\"_blank\" href=\"https://github.com/gebv/pikchr/blob/master/renderserver/\">https://github.com/gebv/pikchr</a>")
}

func stob64(in string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(in))
}
