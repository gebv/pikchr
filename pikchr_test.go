package pikchr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name  string
		src   string
		in    []Option
		want  *RenderResult
		want1 bool
	}{
		{name: "empty", want: &RenderResult{Data: "<!-- empty pikchr diagram -->\n"}, want1: false},
		{name: "ok", src: `box "some box"`, want: &RenderResult{Data: `<svg xmlns='http://www.w3.org/2000/svg' viewBox="0 0 112.32 76.32">
<path d="M2,74L110,74L110,2L2,2Z"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
<text x="56" y="38" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">some box</text>
</svg>
`, Width: 112, Height: 76}, want1: true},
		{name: "dark", src: `box "some box"`, in: []Option{Dark()}, want: &RenderResult{Data: `<svg xmlns='http://www.w3.org/2000/svg' viewBox="0 0 112.32 76.32">
<path d="M2,74L110,74L110,2L2,2Z"  style="fill:none;stroke-width:2.16;stroke:rgb(255,255,255);" />
<text x="56" y="38" text-anchor="middle" fill="rgb(255,255,255)" dominant-baseline="central">some box</text>
</svg>
`, Width: 112, Height: 76}, want1: true},
		{name: "classname", src: `box "some box"`, in: []Option{SVGClass("foobar")}, want: &RenderResult{Data: `<svg xmlns='http://www.w3.org/2000/svg' class="foobar" viewBox="0 0 112.32 76.32">
<path d="M2,74L110,74L110,2L2,2Z"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
<text x="56" y="38" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">some box</text>
</svg>
`, Width: 112, Height: 76}, want1: true},
		{name: "sintax error", src: `boooooox "some box"`, in: []Option{}, want: &RenderResult{Data: "/*    1 */  boooooox \"some box\"\n                    ^^^^^^^^^^\nERROR: syntax error\n", Width: -1, Height: -1}, want1: false},
		{name: "sintax error as html", src: `boooooox "some box"`, in: []Option{HTMLError()}, want: &RenderResult{Data: "<div><pre>\n/*    1 */  boooooox \"some box\"\n                    ^^^^^^^^^^\nERROR: syntax error\n</pre></div>\n", Width: -1, Height: -1}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Render(tt.src, tt.in...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Render() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Render() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Example() {
	in := `arrow right 200% "Markdown" "Source"
box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
arrow right 200% "HTML+SVG" "Output"
arrow <-> down from last box.s
box same "Pikchr" "Formatter" "(pikchr.c)" fit
	`
	res, ok := Render(in)
	fmt.Println("Success?", ok)
	fmt.Println("Width =", res.Width)
	fmt.Println("Height =", res.Height)
	fmt.Println()
	fmt.Println(res.Data)

	// Output:
	// Success? true
	// Width = 423
	// Height = 217

	// <svg xmlns='http://www.w3.org/2000/svg' viewBox="0 0 423.821 217.44">
	// <polygon points="146,37 134,41 134,33" style="fill:rgb(0,0,0)"/>
	// <path d="M2,37L140,37"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
	// <text x="74" y="25" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Markdown</text>
	// <text x="74" y="49" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Source</text>
	// <path d="M161,72L258,72A15 15 0 0 0 273 57L273,17A15 15 0 0 0 258 2L161,2A15 15 0 0 0 146 17L146,57A15 15 0 0 0 161 72Z"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
	// <text x="209" y="17" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Markdown</text>
	// <text x="209" y="37" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Formatter</text>
	// <text x="209" y="57" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">(markdown.c)</text>
	// <polygon points="417,37 405,41 405,33" style="fill:rgb(0,0,0)"/>
	// <path d="M273,37L411,37"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
	// <text x="345" y="25" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">HTML+SVG</text>
	// <text x="345" y="49" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Output</text>
	// <polygon points="209,72 214,84 205,84" style="fill:rgb(0,0,0)"/>
	// <polygon points="209,144 205,133 214,133" style="fill:rgb(0,0,0)"/>
	// <path d="M209,78L209,138"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
	// <path d="M176,215L243,215A15 15 0 0 0 258 200L258,159A15 15 0 0 0 243 144L176,144A15 15 0 0 0 161 159L161,200A15 15 0 0 0 176 215Z"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
	// <text x="209" y="159" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Pikchr</text>
	// <text x="209" y="180" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Formatter</text>
	// <text x="209" y="200" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">(pikchr.c)</text>
	// </svg>
}
