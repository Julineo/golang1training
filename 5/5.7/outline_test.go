package main

import (
	"testing"
	"fmt"
	"bytes"
	"strings"
	"io"
	"os"

	"golang.org/x/net/html"
)

func TestOutline (t *testing.T) {
	input := `
<!DOCTYPE html>
<html>
    <head>
        <link rel="stylesheet" href="css/bootstrap.min.css">
        <script type="text/javascript" src="js/jquery-1.12.0.min.js"></script>
        <script type="text/javascript" src="js/bootstrap.min.js"></script>
    </head>
    <body>
        <div class="container">
            <div class="row">
                <div class="col-md-6"><img src="http://placehold.it/100x100"></div>
                <div class="col-md-6 text-right">
                    <h1>Alexey Dvoretskiy</h1>
                    <h3>Software Engineer</h3>
                </div>
            </div>
        </div>
    </body>
</html>
`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var out io.Writer = os.Stdout
	out = new(bytes.Buffer) //to capture output	
	forEachNode(doc, startElement, endElement)
	fmt.Println(out.(*bytes.Buffer).String())

	_, err = html.Parse(strings.NewReader(out.(*bytes.Buffer).String()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
