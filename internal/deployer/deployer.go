package deployer

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"kafka-connect-deployer/internal/config"
	"log"
	"os"
	"time"
)

const (
	contentType = "application/json"
	uri         = "/connectors"
)

var (
	retryWaitMin = 1 * time.Second
	retryWaitMax = 10 * time.Second
)

type Deployer interface {
	Deploy()
}

type KafkaConnectDeployer struct {
	config config.Config
	client *retryablehttp.Client
}

func New(config config.Config) Deployer {
	c := retryablehttp.NewClient()
	c.RetryMax = config.MaxRetry
	c.RetryWaitMin = retryWaitMin
	c.RetryWaitMax = retryWaitMax
	return &KafkaConnectDeployer{
		config: config,
		client: c,
	}
}

func (k *KafkaConnectDeployer) Deploy() {
	files, err := ioutil.ReadDir(k.config.ConnectorsDir)
	if err != nil {
		log.Panic("invalid dir: "+k.config.ConnectorsDir, err)
	}

	for _, file := range files {
		k.deploy(file.Name())
	}
}

func (k *KafkaConnectDeployer) deploy(fileName string) {
	log.Print("start deploy " + fileName)

	file, err := os.Open(k.config.ConnectorsDir + "/" + fileName)
	if err != nil {
		log.Fatal("fail to open file: "+fileName, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("fail on close file %v \n", err)
		}
	}()

	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("fail to read file: "+fileName, err)
	}

	res, err := k.client.Post(k.config.ConnectUrl+uri, contentType, fileBody)
	if err != nil {
		log.Fatal("fail to post "+fileName, err)
	}

	if res.StatusCode > 299 {
		bodyBuf := new(bytes.Buffer)
		if _, err := bodyBuf.ReadFrom(res.Body); err != nil {
			log.Fatal(fmt.Sprintf("response not ok and fail to open body: %s - code: %d: - error: %v", fileName, res.StatusCode, err))
		}

		log.Fatal(fmt.Sprintf("response not ok: %s - code: %d - body: %s", fileName, res.StatusCode, bodyBuf.String()))
	}

	log.Printf("%s - status: %s \n", fileName, res.Status)
}
