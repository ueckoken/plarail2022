package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type LogViewer interface {
	GetLog(ctx context.Context, tagname, namespace string) (string, error)
}

// Client is a kubernetes client
type Client struct {
	k8sclient *kubernetes.Clientset
}

// NewClient creates a kubernetes client
func NewClient() (*Client, error) {
	if err := appsv1.AddToScheme(scheme.Scheme); err != nil {
		return nil, fmt.Errorf("failed to add scheme: %w", err)
	}
	if err := corev1.AddToScheme(scheme.Scheme); err != nil {
		return nil, fmt.Errorf("failed to add scheme: %w", err)
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rest incluster config: %w", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create cientset: %w", err)
	}
	return &Client{k8sclient: clientset}, nil
}

// GetLog returns deployment logs by deployment name in specified namespace
func (c *Client) GetLog(ctx context.Context, tagname, namespace string) (string, error) {
	pods, err := c.k8sclient.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", tagname),
	})
	if err != nil {
		return "", fmt.Errorf("failed to fetch pods: %w", err)
	}
	for _, p := range pods.Items {
		req := c.k8sclient.CoreV1().Pods(namespace).GetLogs(p.Name, &corev1.PodLogOptions{})
		podLogs, err := req.Stream(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to stream logs: %w", err)
		}
		defer podLogs.Close()
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, podLogs); err != nil {
			return "", fmt.Errorf("failed to copy logs to buffer: %w", err)
		}
		return buf.String(), nil
	}
	return "", fmt.Errorf("no such pod")
}
