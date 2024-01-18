package handler

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/google/uuid"
	"github.com/raanfefu/basajaun/gui/cmd/policy"
	"github.com/raanfefu/basajaun/gui/cmd/settings"
	"github.com/raanfefu/basajaun/gui/pkg"
)

const TEMP_DIRECTORY = "/tmp/"
const ROOT_SCRIPT = "/app/scripts/opa-build.sh"

func PostPublish(w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)

	var body pkg.PublishRequest
	json.NewDecoder(r.Body).Decode(&body)
	response := policy.PostPolicy(*body.Rule)
	settings := settings.GetSettings()

	ROOT := fmt.Sprintf("%s%s/", TEMP_DIRECTORY, uuid.New().String())

	err := os.Mkdir(ROOT, os.ModePerm)
	if err != nil {
		panic(err)
	}
	src := fmt.Sprintf("%s%s", ROOT, "src")
	err = os.Mkdir(src, os.ModePerm)
	if err != nil {
		panic(err)
	}

	rule_string, _ := json.MarshalIndent(response.Rule, "", "  ")
	rego_string := response.Render

	manifest := &pkg.Manifest{
		Name: body.Name,
		Type: "REST",
	}
	manifest_string, _ := json.MarshalIndent(manifest, "", "  ")

	manifestPath := fmt.Sprintf("%s%s", ROOT, "manifest.json")
	err = os.WriteFile(manifestPath, []byte(manifest_string), 0644)
	if err != nil {
		panic("Error")
	}
	restfile := fmt.Sprintf("%s%s", ROOT, "rest.json")
	err = os.WriteFile(restfile, []byte(rule_string), 0644)
	if err != nil {
		panic("Error")
	}

	regofile := fmt.Sprintf("%s%s", ROOT, "src/policy.rego")
	err = os.WriteFile(regofile, []byte(*rego_string), 0644)
	if err != nil {
		panic("Error")
	}

	zipfile := fmt.Sprintf("%s%s", manifest.Name, ".tar.gz")
	cmd := exec.Command(ROOT_SCRIPT, ROOT, zipfile)

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// UPLOAD FILE
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://%s.blob.core.windows.net/", settings.AzureCrend.StorageAccountName)
	client, err := azblob.NewClient(url, credential, nil)

	if err != nil {
		panic(err)
	}

	filetoupload := fmt.Sprintf("%s%s", ROOT, zipfile)
	file, err := os.Open(filetoupload)
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

	options := &azblob.UploadFileOptions{
		BlockSize:   int64(100),
		Concurrency: uint16(10),
		Progress: func(bytesTransferred int64) {
			fmt.Println(bytesTransferred)
		},
	}
	_, err = client.UploadBuffer(context.TODO(), settings.AzureCrend.ContainerName, zipfile, bs, options)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic(err)
	}
	//UPLOAD

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(""))
}
