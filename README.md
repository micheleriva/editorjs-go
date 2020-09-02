<img src="/misc/cover.png" alt="Editorjs.go" />

A simple library which converts **[Editor.js](https://editorjs.io)** JSON output to **Markdown** or **HTML**.

## Installation

```bash
go get github.com/micheleriva/editorjs-go
```

## Usage

Let's suppose that we have the following Editor.js output saved in a file called `editorjs_output.json`:

```json
{
  "blocks": [
    {
      "type" : "header",
      "data" : {
        "text" : "Editor.js",
        "level" : 2
      }
    },
    {
      "type" : "paragraph",
      "data" : {
        "text" : "Hey. Meet the new Editor. On this page you can see it in action — try to edit this text."
      }
    }
  ]
}
```

```go
package main

import (
  "fmt"
  editorjs "github.com/micheleriva/editorjs-go"
	"io/ioutil"
	"log"
)

func main() {
  myJSON, err := ioutil.ReadFile("./editorjs_output.json")
	if err != nil {
		log.Fatal(err)
  }
  
	resultMarkdown := editorjs.Markdown(string(data))
  resultHTML := editorjs.HTML(string(data))

  fmt.Println("=== MARKDOWN ===\n")
  fmt.Println(resultMarkdown)

  fmt.Println("=== HTML ===\n")
  fmt.Println(resultHTML)
}
```

It will generate the following output:

```
=== MARKDOWN ==="

## Editor.js

Hey. Meet the new Editor. On this page you can see it in action — try to edit this text.

=== HTML ===

<h2> Editor.js </h2>
<p>Hey. Meet the new Editor. On this page you can see it in action — try to edit this text.</p>
```

## License
[GPLv3](/LICENSE.md)