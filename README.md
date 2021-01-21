# pikchr

![CI](https://github.com/gebv/pikchr/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/gebv/pikchr)](https://goreportcard.com/report/github.com/gebv/pikchr)

This repository is a wrapper around the `pikchr.c` for use in Golang.

`pikchr.c` and `pikchr.h` version downloaded from that [https://pikchr.org/home/dir?ci=tip&type=tree](https://pikchr.org/home/dir?ci=tip&type=tree) on 2020-12-04 21:07:42.


# What is pikchr?

> Pikchr (pronounced "picture") is a PIC-like markup language for diagrams in technical documentation. Pikchr is designed to be embedded in fenced code blocks of Markdown or similar mechanisms of other documentation markup languages.
>
> For example, the diagram:
>
> <img width="423" height="217" src=".github/demo.svg" />
>
> Is generated by 7 lines of Markdown:
>
>   ``` pikchr
>   arrow right 200% "Markdown" "Source"
>   box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
>   arrow right 200% "HTML+SVG" "Output"
>   arrow <-> down 70% from last box.s
>   box same "Pikchr" "Formatter" "(pikchr.c)" fit
>   ```

[More deatils on the official website](https://pikchr.org/home/doc/trunk/homepage.md)


# Example of the use

```go
package main

import (
	"fmt"

	"github.com/gebv/pikchr"
)

func main() {
	in := `arrow right 200% "Markdown" "Source"
box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
arrow right 200% "HTML+SVG" "Output"
arrow <-> down from last box.s
box same "Pikchr" "Formatter" "(pikchr.c)" fit
	`
	res, ok := pikchr.Render(
        in,
        // pikchr.Dark(), // render the image in dark mode
        // pikchr.SVGClass("foobar"), // add class="%s" to <svg>
        // pikchr.HTMLError(), // wrap the error message text with a html <div><pre>
        )
	fmt.Println("Success?", ok)
	fmt.Println("Width =", res.Width)
	fmt.Println("Height =", res.Height)
	fmt.Println()
    fmt.Println(res.Data)
}

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

```

