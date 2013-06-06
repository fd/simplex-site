package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(index))
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Forwarded-Proto") != "https" {
		http.Redirect(res, req, "http://simplex.sh"+req.URL.Path, 301)
		return
	}

	if strings.HasPrefix(req.Host, "www.") {
		http.Redirect(res, req, "http://simplex.sh"+req.URL.Path, 301)
		return
	}

	str := strings.Replace(index_html, "{{.Path}}", req.URL.Path, -1)
	fmt.Fprintln(res, str)
}

const index_html = `<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8" />
  <title>Simplex</title>

  <meta name="go-import" content="simplex.sh git git://github.com/fd/simplex">

  <link href='https://fonts.googleapis.com/css?family=Wire+One' rel='stylesheet' type='text/css'>
  <style type="text/css" media="all">

  html, body {
    margin: 0;
    padding: 0;
    height: 100%;
    width: 100%;
    position: relative;
    overflow:hidden;
  }

  .Wrapper {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 0;
    height: 0;
  }

  .Content {
    position: absolute;
    width: 12000px;
    left:  -6000px;
    top:   -160px;
    text-align: center;
    font-family: 'Wire One', sans-serif;
  }


  .Simplex {
    font-size:   160px;
  }

  </style>
</head>
<body>

<div class="Wrapper">
  <div class="Content">
    <span class="Simplex">________________________________________________________________________________________________________________________________Simplex________________________________________________________________________________________________________________________________</span>

    <br>

    <a href="http://godoc.org/simplex.sh{{.Path}}"><var>import "simplex.sh{{.Path}}"</var></a>
  </div>
</div>

</body>
</html>`
