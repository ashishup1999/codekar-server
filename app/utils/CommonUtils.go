package utils

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
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

func GetProjectThumbnail(html, css, javascript string) ([]byte, error) {
	//creating a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	//image buffer and htmlText
	var imgBuff []byte
	htmlText := GetProjectPreview(html, css, javascript)

	//capturing screenshot
	err := chromedp.Run(ctx, chromedp.Navigate("data:text/html,"+htmlText), chromedp.CaptureScreenshot(&imgBuff))
	if err != nil {
		return nil, err
	}

	return imgBuff, nil

}
