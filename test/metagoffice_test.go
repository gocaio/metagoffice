/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package test

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/gocaio/metagoffice"
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
	file, err := os.Open("document.docx")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	actualResult, err := metagoffice.GetContent(file)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(docExpectedResult, actualResult) {
		t.Fatalf("Expected %s but got %s", docExpectedResult, actualResult)
	}
}

func TestExcelDocumment(t *testing.T) {
	file, err := os.Open("Book1.xslx")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	actualResult, err := metagoffice.GetContent(file)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(excelExpectedResult, actualResult) {
		t.Fatalf("Expected %s but got %s", excelExpectedResult, actualResult)
	}
}

func TestPowerDocumment(t *testing.T) {
	file, err := os.Open("Presentation1.pptx")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	actualResult, err := metagoffice.GetContent(file)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(powerExpectedResult, actualResult) {
		t.Fatalf("Expected %s but got %s", powerExpectedResult, actualResult)
	}
}
