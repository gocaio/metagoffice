package main

import (
	"fmt"

	"github.com/kevinborras/metagoffice"
)

func main() {
	
	content := metagoffice.GetContent("example/document.docx")

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
