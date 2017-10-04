# website
The new Panux website

## Using the included scripts
### Deploying site locally
First, install quickserve:

`go get github.com/jadr2ddude/quickserve`

Then to update the site and run quickserve:

`make run`

The pages will now be served on port 8080. Note that with this command the root path will show a file list instead of the homepage (Home.html).
### Format of page files
The files in the pages are split into two parts by an `@---@`. Everything before is processed as YAML, used for templating the page. In the YAML portion, the two values you can set are:

* title: the title of the page (defaults to the name of the file with the extension removed)
* data: user-defined data available in templating as .Data; internally a:
```go
map[string]interface{}
```

The part below the `@---@` is run through Go [text/template](https://golang.org/pkg/text/template/) with the YAML data passed as {{.}}. You can use the data portion to generate html here.
### site.tmpl
In the root directory of this repo is a file called site.tmpl. This is a Go [text/template](https://golang.org/pkg/text/template/) which is run with the {{.}} set to the YAML values in each page. In addition {{.Page}} is set to the page text after being templated before.
## What the site uses
This site uses Materialize for the page layout/elements and jQuery for scripting.
