package test

import (
	"reflect"
	"testing"

	"github.com/kevinborras/metagoffice"
)

var docExpectedResult = metagoffice.XMLContent{
	Title:          "Test Document",
	Creator:        "Soler, Kevin",
	Keywords:       "This is a tag",
	Description:    "This is a comment",
	LastModifiedBy: "Soler, Kevin",
	Revision:       "4",
	Created:        "2018-10-12T14:01:00Z",
	Modified:       "2018-10-15T11:23:00Z",
}

var excelExpectedResult = metagoffice.XMLContent{
	Title:          "This is a title",
	Creator:        "Soler, Kevin",
	Keywords:       "1234",
	LastModifiedBy: "Soler, Kevin",
	Created:        "2018-10-15T10:51:30Z",
	Modified:       "2018-10-15T10:52:41Z",
	Category:       "Confidential",
}

var powerExpectedResult = metagoffice.XMLContent{
	Title:          "PowerPoint Title",
	Creator:        "Soler, Kevin",
	Keywords:       "This is a tag",
	LastModifiedBy: "Soler, Kevin",
	Revision:       "1",
	Created:        "2018-10-15T11:47:28Z",
	Modified:       "2018-10-15T11:48:22Z",
	Category:       "comercial",
}

func TestDocxDocumment(t *testing.T) {
	actualResult := metagoffice.GetContent("test/document.docx")

	if !reflect.DeepEqual(docExpectedResult, actualResult) {
		t.Fatalf("Expected %s but got %s", docExpectedResult, actualResult)
	}
}

func TestExcelDocumment(t *testing.T) {
	actualResult := metagoffice.GetContent("test/Book1.xlsx")

	if !reflect.DeepEqual(excelExpectedResult, actualResult) {
		t.Fatalf("Expected %s but got %s", excelExpectedResult, actualResult)
	}
}

func TestPowerDocumment(t *testing.T) {
	actualResult := metagoffice.GetContent("test/Presentation1.pptx")

	if !reflect.DeepEqual(powerExpectedResult, actualResult) {
		t.Fatalf("Expected %s but got %s", powerExpectedResult, actualResult)
	}
}
