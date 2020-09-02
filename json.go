package main

import (
	"encoding/json"
	"log"
)

type EditorJS struct {
	Blocks []EditorJSBlock `json:"blocks"`
}

type EditorJSBlock struct {
	Type string       `json:"type"`
	Data EditorJSData `json:"data"`
}

type EditorJSData struct {
	Text           string     `json:"text",omitempty`
	Level          int        `json:"level,omitempty" `
	Style          string     `json:"style,omitempty" `
	Items          []string   `json:"items,omitempty" `
	File           FileData   `json:"file,omitempty" `
	Caption        string     `json:"caption,omitempty"`
	WithBorder     bool       `json:"withBorder,omitempty"`
	Stretched      bool       `json:"stretched,omitempty"`
	WithBackground bool       `json:"withBackground,omitempty"`
	HTML           string     `json:"html,omitempty"`
	Content        [][]string `json:"content,omitempty"`
	Alignment      string     `json:"alignment,omitempty"`
}

type FileData struct {
	URL string `json:"url"`
}

func ParseEditorJSON(editorJS string) EditorJS {
	var result EditorJS

	err := json.Unmarshal([]byte(editorJS), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
