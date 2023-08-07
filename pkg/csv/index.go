package csv

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dezhishen/file-data-pipeline/pkg/config"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

type csvDecoder struct {
	Headers []config.Header
}

func (i *csvDecoder) Decode(data []byte, v interface{}) error {
	reader := csv.NewReader(bytes.NewReader(data))
	rr, err := reader.ReadAll()
	if err != nil {
		return errors.New("gojsonq: " + err.Error())
	}
	if len(rr) < 1 {
		return errors.New("gojsonq: csv data can't be empty! At least contain the header row")
	}
	var arr = make([]map[string]interface{}, 0)
	headers := i.Headers
	for i := 0; i <= len(rr)-1; i++ {
		if rr[i] == nil { // if a row is empty, skip it
			continue
		}
		mp := map[string]interface{}{}
		for j, header := range headers {
			typ := header.Type
			hdr := header.Name
			switch typ {
			default:
				mp[hdr] = rr[i][j]
			case "STRING":
				mp[hdr] = rr[i][j]
			case "NUMBER":
				if fv, err := strconv.ParseFloat(rr[i][j], 64); err == nil {
					mp[hdr] = fv
				} else {
					mp[hdr] = 0.0
				}
			case "BOOLEAN":
				if strings.ToLower(rr[i][j]) == "true" ||
					rr[i][j] == "1" {
					mp[hdr] = true
				} else {
					mp[hdr] = false
				}
			}
		}
		arr = append(arr, mp)
	}
	bb, err := json.Marshal(arr)
	if err != nil {
		return fmt.Errorf("gojsonq: %v", err)
	}
	return json.Unmarshal(bb, &v)
}

func Read(path string, dataConfig config.PipeLine) (*gojsonq.JSONQ, error) {
	var jq *gojsonq.JSONQ
	if dataConfig.FileFormat == "csv" {
		jq = gojsonq.New(gojsonq.SetDecoder(&csvDecoder{})).File(concatPath(path, dataConfig.Name, dataConfig.FileFormat))
	} else if dataConfig.FileFormat == "json" {
		jq = gojsonq.New().File(concatPath(path, dataConfig.Name, dataConfig.FileFormat))
	} else {
		return nil, errors.New("gojsonq: unsupported file format")
	}
	return jq, nil
}

func Write(path string, dataConfig config.PipeLine, jq *gojsonq.JSONQ) error {
	return nil
}

func concatPath(path string, fileName string, fileType string) string {
	if path == "" {
		path = "." // current directory
	}
	if strings.HasSuffix(path, "/") {
		return path + fileName
	}
	return path + "/" + fileName + "." + fileType
}
