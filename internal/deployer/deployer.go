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
	Deploy() error
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

func (k *KafkaConnectDeployer) Deploy() error {
	if k.config.WaitStartTime > time.Nanosecond {
		log.Printf("waiting %s before start", k.config.WaitStartTime)
		time.Sleep(k.config.WaitStartTime)
	}

	files, err := ioutil.ReadDir(k.config.ConnectorsDir)
	if err != nil {
		return fmt.Errorf("fail to read dir %s. error: %v", k.config.ConnectorsDir, err)
	}

	for _, file := range files {
		if err := k.deploy(file.Name()); err != nil {
			return err
		}
	}

	return nil
}

func (k *KafkaConnectDeployer) deploy(fileName string) error {
	log.Print("start deploy " + fileName)

	file, err := os.Open(k.config.ConnectorsDir + "/" + fileName)
	if err != nil {
		return fmt.Errorf("fail to open file %s. error: %v", fileName, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("fail on close file %v \n", err)
		}
	}()

	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("fail to read file %s: %v", fileName, err)
	}

	res, err := k.client.Post(k.config.ConnectUrl+uri, contentType, fileBody)
	if err != nil {
		return fmt.Errorf("fail to post connector %s. error: %v", fileName, err)
	}

	if res.StatusCode > 299 {
		bodyBuf := new(bytes.Buffer)
		if _, err := bodyBuf.ReadFrom(res.Body); err != nil {
			return fmt.Errorf("response not ok and fail to open body: %s - code: %d: - error: %v", fileName, res.StatusCode, err)
		}

		return fmt.Errorf("response not ok: %s - code: %d - body: %s", fileName, res.StatusCode, bodyBuf.String())
	}

	log.Printf("%s - status: %s \n", fileName, res.Status)
	return nil
}
