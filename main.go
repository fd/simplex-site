package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(index))
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, index_html)
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

  .Simplex {
    position: absolute;
    width: 12000px;
    left:  -6000px;
    top:   -160px;


    text-align:center;
    white-space: pre;
    font-family: 'Wire One', sans-serif;
    font-size: 160px;
  }

  </style>
</head>
<body>

<div class="Wrapper">
<div class="Simplex">________________________________________________________________________________________________________________________________Simplex________________________________________________________________________________________________________________________________</div>
</div>

</body>
</html>`
