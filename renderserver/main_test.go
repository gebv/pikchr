package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderServer(t *testing.T) {
	app := &renderServer{}
	ts := httptest.NewServer(app)
	defer ts.Close()

	formatURL := func(urlPath string) string {
		return "http://" + ts.Listener.Addr().String() + urlPath
	}
	t.Log(formatURL("/"))

	cases := []struct {
		name           string
		method         string
		diagramSrc     string
		wantStatusCode int
		wantSuccess    bool
		wantSVGRaw     string
		wantErr        string
	}{
		{name: "empty scr", method: http.MethodPost, wantSuccess: false, wantStatusCode: http.StatusBadRequest, wantErr: "invalid arguments: empty src of diagram"},
		{name: "invalid sintax", method: http.MethodPost, diagramSrc: "foobar", wantSuccess: false, wantStatusCode: http.StatusOK, wantErr: "failed render: /*    1 */  foobar\n           \nERROR: syntax error\n"},
		{name: "ok", method: http.MethodPost, diagramSrc: `text "some title"
box "some box"
`, wantSuccess: true, wantStatusCode: http.StatusOK, wantSVGRaw: `<svg xmlns='http://www.w3.org/2000/svg' viewBox="0 0 201.254 76.32">
<text x="43" y="38" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">some title</text>
<path d="M91,74L199,74L199,2L91,2Z"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
<text x="145" y="38" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">some box</text>
</svg>
`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			reqData := map[string]interface{}{
				"diagram_src": c.diagramSrc,
			}
			reqBytes, err := json.Marshal(reqData)
			require.NoError(t, err, "encode request")

			t.Logf("request json: %q", string(reqBytes))
			req, err := http.NewRequest(c.method, formatURL("/"), bytes.NewReader(reqBytes))
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			app.ServeHTTP(rec, req)

			assert.EqualValues(t, c.wantStatusCode, rec.Result().StatusCode)

			t.Logf("response json: %q", rec.Body.Bytes())

			resData := map[string]interface{}{}
			err = json.Unmarshal(rec.Body.Bytes(), &resData)
			require.NoError(t, err, "decode response")

			require.EqualValues(t, c.wantSuccess, resData["success"])
			if resData["success"].(bool) {
				assert.EqualValues(t, c.wantSVGRaw, resData["svg_raw"])
			} else {
				assert.EqualValues(t, c.wantErr, resData["err"])
			}
		})
	}
}
