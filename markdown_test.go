package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHeader(t *testing.T) {

	// Test H1 Header
	input1 := EditorJSData{
		Text:  "Level 1 Header",
		Level: 1,
	}

	expected1 := "# Level 1 Header"
	actual1 := generateHeader(input1)

	// Test H2 Header
	input2 := EditorJSData{
		Text:  "Level 2 Header",
		Level: 2,
	}

	expected2 := "## Level 2 Header"
	actual2 := generateHeader(input2)

	// Test H3 Header
	input3 := EditorJSData{
		Text:  "Level 3 Header",
		Level: 3,
	}

	expected3 := "### Level 3 Header"
	actual3 := generateHeader(input3)

	// Test H4 Header
	input4 := EditorJSData{
		Text:  "Level 4 Header",
		Level: 4,
	}

	expected4 := "#### Level 4 Header"
	actual4 := generateHeader(input4)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
	assert.Equal(t, expected3, actual3)
	assert.Equal(t, expected4, actual4)
}

func TestGenerateUnorderedList(t *testing.T) {
	input := EditorJSData{
		Style: "unordered",
		Items: []string{"first", "second", "third"},
	}

	expected := `- first
- second
- third`

	actual := generateList(input)

	assert.Equal(t, expected, actual)
}

func TestGenerateOrderedList(t *testing.T) {
	input := EditorJSData{
		Style: "ordered",
		Items: []string{"first", "second", "third"},
	}

	expected := `1. first
2. second
3. third`

	actual := generateList(input)

	assert.Equal(t, expected, actual)
}

func TestGenerateImageWithoutOptions(t *testing.T) {
	input := EditorJSData{
		File: FileData{
			URL: "https://example.com/img.png",
		},
	}

	expected := `![](https://example.com/img.png)`
	actual := generateImage(input, Options{})

	assert.Equal(t, expected, actual)
}

func TestGenerateImageWithPartialOptions(t *testing.T) {
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

	expected := `![My beautiful image](https://example.com/img.png)`
	actual := generateImage(input, options)

	assert.Equal(t, expected, actual)
}

func TestGenerateImageWithFullOptions(t *testing.T) {
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
	actual := generateImage(input, options)

	assert.Equal(t, expected, actual)
}
