package settings

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/raanfefu/basajaun/gui/pkg"
)

const DATAPATH = "/app/data/data.json"

func GetSettings() *pkg.Settings {

	file, err := os.Open(DATAPATH)
	if err != nil {
		panic(err)
	}

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		panic(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(bs), &result)

	return &pkg.Settings{
		Repository:  "https://github.com/raanfefu/basajaun-demo",
		GithubToken: "B41BD5F462719C6D6118E673A2389",
		AzureCrend: &pkg.ACredentials{
			TenantID:           os.Getenv("AZURE_TENANT_ID"),
			ClientID:           os.Getenv("AZURE_CLIENT_ID"),
			ClientSecret:       os.Getenv("AZURE_CLIENT_SECRET"),
			StorageAccountName: os.Getenv("AZURE_STORAGE_ACCOUNT"),
			ContainerName:      os.Getenv("AZURE_BLON_CONTAINER"),
		},
		TypeBundle: "REST",
		Data:       result,
	}
}
