package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func HTML(input string, options ...Options) string {
	var markdownOptions Options

	if len(options) > 0 {
		markdownOptions = options[0]
	}

	var result []string
	editorJSAST := ParseEditorJSON(input)

	for _, el := range editorJSAST.Blocks {

		data := el.Data

		switch el.Type {

		case "header":
			result = append(result, generateHTMLHeader(data))

		case "paragraph":
			result = append(result, generateHTMLParagraph(el.Data))

		case "list":
			result = append(result, generateMDList(data))

		case "image":
			result = append(result, generateHTMLImage(data, markdownOptions))

		case "rawTool":
			result = append(result, data.HTML)

		case "delimiter":
			result = append(result, "---")

		case "table":
			result = append(result, generateMDTable(data))

		case "caption":
			result = append(result, generateMDCaption(data))

		default:
			log.Fatal("Unknown data type: " + el.Type)
		}

	}

	return strings.Join(result[:], "\n\n")
}

func generateHTMLHeader(el EditorJSData) string {
	level := strconv.Itoa(el.Level)
	return fmt.Sprintf("<h%s>%s</h%s>", level, el.Text, level)
}

func generateHTMLParagraph(el EditorJSData) string {
	return fmt.Sprintf("<p>%s</p>", el.Text)
}

func generateHTMLList(el EditorJSData) string {
	var result []string

	if el.Style == "unordered" {
		result = append(result, "<ul>")

		for _, el := range el.Items {
			result = append(result, "  <li>"+el+"</li>")
		}

		result = append(result, "</ul>")
	} else {
		result = append(result, "<ol>")

		for _, el := range el.Items {
			result = append(result, "  <li>"+el+"</li>")
		}

		result = append(result, "</ol>")
	}

	return strings.Join(result[:], "\n")
}

func generateHTMLImage(el EditorJSData, options Options) string {
	classes := options.Image.Classes
	withBorder := classes.WithBorder
	stretched := classes.Stretched
	withBackground := classes.WithBackground

	if withBorder == "" && el.WithBorder {
		withBorder = "editorjs-with-border"
	}

	if stretched == "" && el.Stretched {
		stretched = "editorjs-stretched"
	}

	if withBackground == "" && el.WithBackground {
		withBackground = "editorjs-withBackground"
	}

	return fmt.Sprintf(`<img src="%s" alt="%s" class="%s %s %s" />`, el.File.URL, options.Image.Caption, withBorder, stretched, withBackground)
}
