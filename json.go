package main

import "encoding/json"

type EditorJS struct {
	Blocks []EditorJSBlock `json:"blocks"`
}

type EditorJSBlock struct {
	Type string       `json:"type"`
	Data EditorJSData `json:"data"`
}

type EditorJSData struct {
	Text           string   `json:"text",omitempty`
	Level          int      `json:"level,omitempty" `
	Style          string   `json:"style,omitempty" `
	Items          []string `json:"items,omitempty" `
	File           FileData `json:"file,omitempty" `
	Caption        string   `json:"caption,omitempty"`
	WithBorder     bool     `json:"withBorder,omitempty"`
	Stretched      bool     `json:"stretched,omitempty"`
	WithBackground bool     `json:"withBackground,omitempty"`
}

type FileData struct {
	URL string `json:"url"`
}

func ParseEditorJSON(editorJS string) EditorJS {
	var result EditorJS
	json.Unmarshal([]byte(editorJS), &result)

	return result
}
