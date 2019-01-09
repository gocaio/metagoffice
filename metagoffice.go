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

package metagoffice

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"errors"
	"os"
	"strings"
)

//XMLContent contains the fields of te file core.xml
type XMLContent struct {
	Title          string `xml:"title"`
	Subject        string `xml:"subject"`
	Creator        string `xml:"creator"`
	Keywords       string `xml:"keywords"`
	Description    string `xml:"description"`
	LastModifiedBy string `xml:"lastModifiedBy"`
	Revision       string `xml:"revision"`
	Created        string `xml:"created"`
	Modified       string `xml:"modified"`
	Category       string `xml:"category"`
}

//GetContent function
func GetContent(document *os.File) (fields XMLContent, err error) {
	fileName := document.Name()
	d := XMLContent{}
	dot := strings.Index(fileName, ".")
	before := fileName[:dot] + ".zip"
	err = os.Rename(fileName, before)
	if err != nil {
		return fields, errors.New("Failed to rename as .zip")
	}

	read, err := zip.OpenReader(before)
	if err != nil {
		os.Rename(before, fileName)
		return fields, errors.New("Failed to open the file")
	}

	var xmlFile string
	for _, file := range read.File {
		if file.Name == "docProps/core.xml" {
			rc, err := file.Open()
			scanner := bufio.NewScanner(rc)
			for scanner.Scan() {
				xmlFile += scanner.Text()
			}
			if err != nil {
				os.Rename(before, fileName)
				return fields, errors.New("Failed to open docProps/core.xml")
			}
			defer rc.Close()
		}
	}
	if err := xml.Unmarshal([]byte(xmlFile), &d); err != nil {
		os.Rename(before, fileName)
		return fields, errors.New("Failed to Unmarshal")
	}

	read.Close()
	os.Rename(before, fileName)
	return d, err
}
