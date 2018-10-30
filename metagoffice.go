package metagoffice

import (
        "archive/zip"
        "bufio"
        "encoding/xml"
        "log"
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
func GetContent(document string) (fields XMLContent) {
        dot := strings.Index(document, ".")
        zipFile := document[:dot] + ".zip"
        os.Rename(document, zipFile)

        read, err := zip.OpenReader(zipFile)
        if err != nil {
                msg := "Failed to open: %s"
                log.Fatalf(msg, err)
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
                                log.Fatal(err)
                        }
                        defer rc.Close()
                }
        }
        d := XMLContent{}
        if err := xml.Unmarshal([]byte(xmlFile), &d); err != nil {
                log.Fatal(err)
        }
        read.Close()
        os.Rename(zipFile, document)
        return d
}
