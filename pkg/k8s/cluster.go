package k8s

import (
	"context"
	"encoding/json"
	"github.com/goudai-projects/gd-ops/log"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
)

type Cluster struct {
	mutex *sync.Mutex
	cache *sync.Map
}

var c *Cluster
var once sync.Once

type Getter func() (*kubernetes.Clientset, error)

func GetInstance() *Cluster {
	return c
}

func init() {
	log.Info("init Cluster ...")
	once.Do(func() {
		c = &Cluster{
			mutex: &sync.Mutex{},
			cache: &sync.Map{},
		}
	})
}

func (c *Cluster) CreateDaemonSet(ctx context.Context, clusterId string, daemonSetJson string) (*appsv1.DaemonSet, error) {
	client, err := c.get(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	var daemonSet appsv1.DaemonSet
	err = json.Unmarshal([]byte(daemonSetJson), &daemonSet)
	if err != nil {
		return nil, err
	}
	return client.AppsV1().DaemonSets(daemonSet.Namespace).Create(ctx, &daemonSet, metav1.CreateOptions{})
}

func (c *Cluster) CreateDeployment(ctx context.Context, clusterId string, deploymentJson string) (*appsv1.Deployment, error) {
	client, err := c.get(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	var deployment appsv1.Deployment
	err = json.Unmarshal([]byte(deploymentJson), &deployment)
	if err != nil {
		return nil, err
	}
	return client.AppsV1().Deployments(deployment.Namespace).Create(ctx, &deployment, metav1.CreateOptions{})
}

func (c *Cluster) GetDaemonSetList(ctx context.Context, clusterId string, ns string) (*appsv1.DaemonSetList, error) {
	client, err := c.get(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	return client.AppsV1().DaemonSets(ns).List(ctx, metav1.ListOptions{})
}

func (c *Cluster) GetDeploymentList(ctx context.Context, clusterId string, ns string) (*appsv1.DeploymentList, error) {
	client, err := c.get(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	return client.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})
}

func (c *Cluster) GetNodeList(ctx context.Context, clusterId string) (*v1.NodeList, error) {
	client, err := c.get(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	list, err := client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Cluster) get(ctx context.Context, clusterId string) (*kubernetes.Clientset, error) {
	return c.getK8sClientFromMap(ctx, clusterId)
}

func (c *Cluster) getK8sClientFromMap(ctx context.Context, clusterId string) (*kubernetes.Clientset, error) {
	if f, ok := c.cache.Load(clusterId); ok {
		return f.(Getter)()
	}
	var client *kubernetes.Clientset
	var once sync.Once
	var err error
	wrapGetter := Getter(func() (*kubernetes.Clientset, error) {
		once.Do(func() {
			client, err = c.createK8sClient(clusterId)
		})
		if err != nil {
			log.Warn("init k8s client failed," + err.Error())
			return nil, err
		}
		return client, err
	})

	f, loaded := c.cache.LoadOrStore(clusterId, wrapGetter)
	if loaded {
		return f.(Getter)()
	}
	return wrapGetter()
}

func (c *Cluster) createK8sClient(clusterId string) (*kubernetes.Clientset, error) {
	// TODO
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(""))
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
