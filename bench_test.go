package main

import (
	"html/template"
	"io/ioutil"
	"testing"
	texttemplate "text/template"
)

func BenchmarkTextTemplateParse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := texttemplate.New("").Parse(tpl)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHTMLTemplateExecute(b *testing.B) {

	type Link struct {
		Title string
		URL   template.URL
	}

	links := []Link{
		Link{Title: "A", URL: "http://a"},
		Link{Title: "B", URL: "http://b"},
		Link{Title: "C", URL: "http://c"},
	}

	data := map[string]interface{}{
		"Links": links,
	}

	templ, err := template.New("").Parse(tpl)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := templ.Execute(ioutil.Discard, data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

const tpl = `
<!doctype html>
<html class="no-js" lang="">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="x-ua-compatible" content="ie=edge">
        <title></title>
        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <link rel="apple-touch-icon" href="apple-touch-icon.png">
        <!-- Place favicon.ico in the root directory -->

        <link rel="stylesheet" href="css/normalize.css">
        <link rel="stylesheet" href="css/main.css">
        <script src="js/vendor/modernizr-2.8.3.min.js"></script>
    </head>
    <body>
        <!--[if lt IE 8]>
            <p class="browserupgrade">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
        <![endif]-->

        <!-- Add your site or application content here -->
        <p>Hello world! This is HTML5 Boilerplate.</p>
		<ol>
		{{ range .Links }}
		<li><a href="{{ .URL }}">{{ .Title }}</a></li>
		{{ end }}
		</ol>

        <script src="https://code.jquery.com/jquery-1.12.0.min.js"></script>
        <script>window.jQuery || document.write('<script src="js/vendor/jquery-1.12.0.min.js"><\/script>')</script>
        <script src="js/plugins.js"></script>
        <script src="js/main.js"></script>

        <!-- Google Analytics: change UA-XXXXX-X to be your site's ID. -->
        <script>
            (function(b,o,i,l,e,r){b.GoogleAnalyticsObject=l;b[l]||(b[l]=
            function(){(b[l].q=b[l].q||[]).push(arguments)});b[l].l=+new Date;
            e=o.createElement(i);r=o.getElementsByTagName(i)[0];
            e.src='https://www.google-analytics.com/analytics.js';
            r.parentNode.insertBefore(e,r)}(window,document,'script','ga'));
            ga('create','UA-XXXXX-X','auto');ga('send','pageview');
        </script>
    </body>
</html>

`
