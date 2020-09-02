package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHeaderHTML(t *testing.T) {

	// Test H1 Header
	input1 := EditorJSData{
		Text:  "Level 1 Header",
		Level: 1,
	}

	expected1 := "<h1>Level 1 Header</h1>"
	actual1 := generateHTMLHeader(input1)

	// Test H2 Header
	input2 := EditorJSData{
		Text:  "Level 2 Header",
		Level: 2,
	}

	expected2 := "<h2>Level 2 Header</h2>"
	actual2 := generateHTMLHeader(input2)

	// Test H3 Header
	input3 := EditorJSData{
		Text:  "Level 3 Header",
		Level: 3,
	}

	expected3 := "<h3>Level 3 Header</h3>"
	actual3 := generateHTMLHeader(input3)

	// Test H4 Header
	input4 := EditorJSData{
		Text:  "Level 4 Header",
		Level: 4,
	}

	expected4 := "<h4>Level 4 Header</h4>"
	actual4 := generateHTMLHeader(input4)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
	assert.Equal(t, expected3, actual3)
	assert.Equal(t, expected4, actual4)
}

func TestGenerateHTMLParagraph(t *testing.T) {
	input := EditorJSData{
		Text: "I am a paragraph!",
	}

	expected := "<p>I am a paragraph!</p>"
	actual := generateHTMLParagraph(input)

	assert.Equal(t, expected, actual)
}

func TestGenerateHTMLUnorderedList(t *testing.T) {
	input := EditorJSData{
		Style: "unordered",
		Items: []string{"first", "second", "third"},
	}

	expected := `<ul>
  <li>first</li>
  <li>second</li>
  <li>third</li>
</ul>`

	actual := generateHTMLList(input)

	assert.Equal(t, expected, actual)
}

func TestGenerateHTMLOrderedList(t *testing.T) {
	input := EditorJSData{
		Style: "ordered",
		Items: []string{"first", "second", "third"},
	}

	expected := `<ol>
  <li>first</li>
  <li>second</li>
  <li>third</li>
</ol>`

	actual := generateHTMLList(input)

	assert.Equal(t, expected, actual)
}

func TestGenerateImageWithoutOptionsHTML(t *testing.T) {
	input := EditorJSData{
		File: FileData{
			URL: "https://example.com/img.png",
		},
	}

	expected := `<img src="https://example.com/img.png" alt="" class="  " />`
	actual := generateHTMLImage(input, Options{})

	assert.Equal(t, expected, actual)
}

func TestGenerateImageWithPartialOptionsHTML(t *testing.T) {
	input := EditorJSData{
		File: FileData{
			URL: "https://example.com/img.png",
		},
	}

	options := Options{
		Image: ImageOptions{
			Caption: "My beautiful image",
		},
	}

	expected := `<img src="https://example.com/img.png" alt="My beautiful image" class="  " />`
	actual := generateHTMLImage(input, options)

	assert.Equal(t, expected, actual)
}

func TestGenerateImageWithFullOptionsHTML(t *testing.T) {
	input := EditorJSData{
		File: FileData{
			URL: "https://example.com/img.png",
		},
	}

	options := Options{
		Image: ImageOptions{
			Caption: "My beautiful image",
			Classes: ImageClasses{
				Stretched:      "streched-class",
				WithBackground: "with-background-class",
				WithBorder:     "with-border-class",
			},
		},
	}

	expected := `<img src="https://example.com/img.png" alt="My beautiful image" class="with-border-class streched-class with-background-class" />`
	actual := generateHTMLImage(input, options)

	assert.Equal(t, expected, actual)
}
