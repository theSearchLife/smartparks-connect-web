package device_template

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

const templateDir = "template_files"

type FwVersion struct {
	Type  string `json:"type"`
	Major string `json:"major"`
	Minor string `json:"minor"`
}
type Hardware struct {
	Type    string `json:"type"`
	Version struct {
		Major string `json:"major"`
		Minor string `json:"minor"`
	} `json:"version"`
}
type DeviceTemplate struct {
	FileName  string          `json:"file_name"`
	FwVersion FwVersion       `json:"fw_version"`
	Hardware  Hardware        `json:"hardware"`
	Settings  json.RawMessage `json:"settings"`
	Commands  json.RawMessage `json:"commands"`
	Ports     json.RawMessage `json:"ports"`
}

var Templates = make(map[string]*DeviceTemplate)

func init() {
	scanTemplateDir()
}

func scanTemplateDir() {
	templateFiles, err := os.ReadDir(templateDir)
	if err != nil {
		log.Fatalf("read template_files dir error %+v", err)
	}
	for _, templateFile := range templateFiles {
		if !templateFile.IsDir() {
			initTemplate(templateFile.Name())
		}
	}
}
func initTemplate(fileName string) {
	templateFileContent, err := os.Open(path.Join(templateDir, fileName))
	if err != nil {
		log.Printf("[error] read template_file %s error %+v", fileName, err)
		return
	}

	decoder := json.NewDecoder(templateFileContent)
	deviceTemplate := &DeviceTemplate{}
	if err := decoder.Decode(deviceTemplate); err == nil {
		deviceTemplate.FileName = fileName
		Templates["v"+deviceTemplate.FwVersion.Major+"."+deviceTemplate.FwVersion.Minor] = deviceTemplate
	} else {
		log.Printf("[error] decode template_file %s error %+v", fileName, err)
	}
}
