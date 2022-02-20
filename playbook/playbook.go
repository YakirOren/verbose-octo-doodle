package playbook

import (
	"bytes"
	"encoding/json"
	"hellowasm/files"
	"log"
	"sync"
)

const (
	PlaybookTmpl = `{{$Name := .Name | Title}}name: {{$Name}} Demo Playbook
type: Flow
active: false
version: "1.0"
desc: "{{$Name}} Playbook"
runner: ""
steps:
  - desc: example description delete afterwards
    text: "# {{$Name}} Actions"{{range $index, $Action := .Actions}}
  - id: S{{$index}}
    desc: Description
    name: "{{$Action.Name | camelcaseSplit}}"
    action: {{$Name}}.{{$Action.Name}}{{end}}`
)

type DescribeFile struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Actions     []struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	} `json:"actions"`
}

func ToPlaybook(input string) (string, error) {
	// marshal into struct
	var file DescribeFile
	if err := json.Unmarshal([]byte(input), &file); err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	// generate playbook from the struct
	if err := files.RunTemplate(buf, PlaybookTmpl, file); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ToPlaybookFromURL(url string) (string, error) {
	var wg sync.WaitGroup
	var result string
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		// download file
		result, err = files.DownloadFileFromUrl(url)
	}()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return ToPlaybook(result)
}
