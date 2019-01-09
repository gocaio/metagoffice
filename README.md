# Package to extract metadata of Office files

[![Travis build](https://travis-ci.com/kevinborras/metagoffice.svg?branch=master)](https://travis-ci.com/kevinborras/metagoffice)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinborras/metagoffice)](https://goreportcard.com/badge/github.com/kevinborras/metagoffice)

This package is intended to be a powerful alternative to other solutions for extract metadata from Office files.

The main features of metagoffice are:

* Read metadata of xlsx files.
* Read metadata of docx files.
* Read metadata of pptx files.

## How to use ?
---

```golang
package main

import (
	"fmt"

	"github.com/kevinborras/metagoffice"
)

func main() {
	
	file, err := os.Open("document.docx")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	content, err := metagoffice.GetContent(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title: ", content.Title)
	fmt.Println("Subject: ", content.Subject)
	fmt.Println("Creator: ", content.Creator)
	fmt.Println("Keywords: ", content.Keywords)
	fmt.Println("Description: ", content.Description)
	fmt.Println("Last modified by: ", content.LastModifiedBy)
	fmt.Println("Revision: ", content.Revision)
	fmt.Println("Created: ", content.Created)
	fmt.Println("Modified: ", content.Modified)
	fmt.Println("Category: ", content.Category)

}
```

Output
```
Title:  Test Document
Subject:
Creator:  Soler, Kevin
Keywords:  This is a tag
Description:  This is a comment
Last modified by:  Soler, Kevin
Revision:  4
Created:  2018-10-12T14:01:00Z
Modified:  2018-10-15T11:23:00Z
Category:
```
