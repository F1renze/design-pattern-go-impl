package adapter

import (
	"fmt"
)

type IPlatformAdapter interface {
	CreateServer(cpu, mem float64) error
}

// AWS
type AWSClient struct {
}

func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("create aws instance, cpu: %f, mem: %f\n", cpu, mem)
	return nil
}

func NewAWSClientAdapter() IPlatformAdapter {
	return &AWSClientAdapter{client: &AWSClient{}}
}

type AWSClientAdapter struct {
	client *AWSClient
}

func (a *AWSClientAdapter) CreateServer(cpu, mem float64) error {
	return a.client.RunInstance(cpu, mem)
}

// AliYun
type AliYunClient struct {
}

func (c *AliYunClient) CreateInstance(cpu, mem int) error {
	fmt.Printf("create aliyun Instance, cpu: %d, mem: %d\n", cpu, mem)
	return nil
}

func NewAliYunClientAdapter() IPlatformAdapter {
	return &AliYunClientAdapter{client: &AliYunClient{}}
}

type AliYunClientAdapter struct {
	client *AliYunClient
}

func (a *AliYunClientAdapter) CreateServer(cpu, mem float64) error {
	return a.client.CreateInstance(int(cpu), int(mem))
}
