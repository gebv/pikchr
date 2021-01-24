# pikchr render server

Two methods
* generation SVG - `POST /`
* home page with info - `GET /`

## POST `/` - generation SVG

* Base path: `/`
* Method: `POST`
* Request\response content type: `application/json`

Request structure

| Field       | Type        | Description |
| ----------- | ----------- | ----------- |
| `diagram_src` | `string` | (required) Source code of pikchr diagram. |
| `dark` | `bool` | Dark mode. Default `false`. |
| `class_name` | `string` | Class name for svg tag. Default `null`. |

### Successfully generated SVG

The data is valid. The request is correct. Successfully generated SVG

Response structure

| Field       | Type        | Description |
| ----------- | ----------- | ----------- |
| `success` | `bool` | (required) `true` if successfully generated SVG. |
| `svg_raw` | `string` | (required) SVG raw data. |
| `svg_raw_base64` | `string` | (required) Encoded base64 from SVG raw data. |
| `img_inline_svg` | `string` | (required) `img` tag with inline svg data `<img height=".." width=".." src="data:image/svg+xml;base64,.."></img>`. |
| `svg_width` | `number` | (required) SVG width. |
| `svg_height` | `number` | (required) SVG height. |

<details>
  <summary>example CURL request</summary>

```log
curl -v -XPOST -d '{"diagram_src":"text \"some title\"\nbox \"some box\"\n"}' -H 'content-type: application/json' http://127.0.0.1:59112/
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 59112 (#0)
> POST / HTTP/1.1
> Host: 127.0.0.1:59112
> User-Agent: curl/7.64.1
> Accept: */*
> content-type: application/json
> Content-Length: 57
>
* upload completely sent off: 57 out of 57 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json
< Server: PikchrRenderServer
< X-Exec-Duration-Ms: 0
< Date: Sun, 24 Jan 2021 06:50:16 GMT
< Content-Length: 1699
<
{"img_inline_svg":"\u003cimg height=\"76\" width=\"201\" src=\"data:image/svg+xml;base64,PHN2ZyB4bWxucz0naHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmcnIHZpZXdCb3g9IjAgMCAyMDEuMjU0IDc2LjMyIj4KPHRleHQgeD0iNDMiIHk9IjM4IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJyZ2IoMCwwLDApIiBkb21pbmFudC1iYXNlbGluZT0iY2VudHJhbCI+c29tZcKgdGl0bGU8L3RleHQ+CjxwYXRoIGQ9Ik05MSw3NEwxOTksNzRMMTk5LDJMOTEsMloiICBzdHlsZT0iZmlsbDpub25lO3N0cm9rZS13aWR0aDoyLjE2O3N0cm9rZTpyZ2IoMCwwLDApOyIgLz4KPHRleHQgeD0iMTQ1IiB5PSIzOCIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZmlsbD0icmdiKDAsMCwwKSIgZG9taW5hbnQtYmFzZWxpbmU9ImNlbnRyYWwiPnNvbWXCoGJveDwvdGV4dD4KPC9zdmc+Cg\"\u003e\u003c/img\u003e","success":true,"svg_height":76,"svg_raw":"\u003csvg xmlns='http://www.w3.org/2000/svg' viewBox=\"0 0 201.254 76.32\"\u003e\n\u003ctext x=\"43\" y=\"38\" text-anchor=\"middle\" fill=\"rgb(0,0,0)\" dominant-baseline=\"central\"\u003esome title\u003c/text\u003e\n\u003cpath d=\"M91,74L199,74L199,2L91,2Z\"  style=\"fill:none;stroke-width:2.16;stroke:rgb(0,0,0);\" /\u003e\n\u003ctext x=\"145\" y=\"38\" text-anchor=\"middle\" fill=\"rgb(0,0,0)\" dominant-baseline=\"central\"\u003esome box\u003c/text\u003e\n\u003c/svg\u003e\n","svg_raw_base64":"PHN2ZyB4bWxucz0naHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmcnIHZpZXdCb3g9IjAgMCAyMDEuMjU0IDc2LjMyIj4KPHRleHQgeD0iNDMiIHk9IjM4IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJyZ2IoMCwwLDApIiBkb21pbmFudC1iYXNlbGluZT0iY2VudHJhbCI+c29tZcKgdGl0bGU8L3RleHQ+CjxwYXRoIGQ9Ik05MSw3NEwxOTksNzRMMTk5LDJMOTEsMloiICBzdHlsZT0iZmlsbDpub25lO3N0cm9rZS13aWR0aDoyLjE2O3N0cm9rZTpyZ2IoMCwwLDApOyIgLz4KPHRleHQgeD0iMTQ1IiB5PSIzOCIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZmlsbD0icmdiKDAsMCwwKSIgZG9taW5hbnQtYmFzZWxpbmU9ImNlbnRyYWwiPnNvbWXCoGJveDwvdGV4dD4KPC9zdmc+Cg","svg_width":201}
* Connection #0 to host 127.0.0.1 left intact
* Closing connection 0
```
</details>


### Error queries

General structure

| Field       | Type        | Description |
| ----------- | ----------- | ----------- |
| `success` | `bool` | (required) `false` if not generated SVG. |
| `err` | `string` | (required) Details error. |

#### Syntax error in the source code of the diagram

Response structure

| Field       | Type        | Description |
| ----------- | ----------- | ----------- |
| `success` | `bool` | (required) `false` if not generated SVG. |
| `err` | `string` | (required) Begin with `failed render: ` and details error. |
