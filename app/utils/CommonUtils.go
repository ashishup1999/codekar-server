package utils

import (
	"fmt"
)

func GetProjectPreview(html, css, javascript string) string {
	var previewString = fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Your Title Here</title>
	  <style>
		*{
			overflow: hidden;
		}
		html {
		  width: 100%%;
		  height: 100%%;
		}
		body {
			width: 100%%;
			height: 100%%;
			margin:0;
			padding:0;
			box-sizing: border-box;
			pointer-events: none;
		}
		%s
	  </style>
	</head>
	<body>
	  %s
	  <script>
	  %s
	  </script>
	</body>
	</html>`, css, html, javascript)
	return previewString
}
