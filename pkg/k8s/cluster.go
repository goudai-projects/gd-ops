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
	clrclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sync"
)

type Cluster struct {
	mutex     *sync.Mutex
	cache     *sync.Map
	clientMap *sync.Map
}

var c Cluster
var once sync.Once

type ClientsetGetter func() (*kubernetes.Clientset, error)
type ClientGetter func() (clrclient.Client, error)

func GetInstance() *Cluster {
	return &c
}

func init() {
	log.Info("init Cluster ...")
	once.Do(func() {
		c = Cluster{}
	})
}

func (c *Cluster) DeploymentByJson(ctx context.Context, clusterId string, daemonSetJson string) (string, error) {
	var daemonSet appsv1.DaemonSet
	err := json.Unmarshal([]byte(daemonSetJson), &daemonSet)
	if err != nil {
		return "", err
	}
	if client, err := c.getClient(ctx, clusterId); err != nil {
		return "", err
	} else {
		update, err := controllerutil.CreateOrUpdate(ctx, client, &daemonSet, func() error {
			err := json.Unmarshal([]byte(daemonSetJson), &daemonSet)
			if err != nil {
				return err
			}
			return nil
		})
		return string(update), err
	}
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
	return c.getK8sClientsetFromMap(ctx, clusterId)
}

func (c *Cluster) getK8sClientsetFromMap(ctx context.Context, clusterId string) (*kubernetes.Clientset, error) {
	if f, ok := c.cache.Load(clusterId); ok {
		return f.(ClientsetGetter)()
	}
	var client *kubernetes.Clientset
	var once sync.Once
	var err error
	wrapGetter := ClientsetGetter(func() (*kubernetes.Clientset, error) {
		once.Do(func() {
			client, err = c.createK8sClientset(clusterId)
		})
		if err != nil {
			log.Warn("init k8s client failed," + err.Error())
			return nil, err
		}
		return client, err
	})

	f, loaded := c.cache.LoadOrStore(clusterId, wrapGetter)
	if loaded {
		return f.(ClientsetGetter)()
	}
	return wrapGetter()
}

func (c *Cluster) getClient(ctx context.Context, clusterId string) (clrclient.Client, error) {
	if f, ok := c.cache.Load(clusterId); ok {
		return f.(ClientGetter)()
	}
	var client clrclient.Client
	var once sync.Once
	var err error
	wrapGetter := ClientGetter(func() (clrclient.Client, error) {
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
		return f.(ClientGetter)()
	}
	return wrapGetter()
}

func (c *Cluster) createK8sClientset(clusterId string) (*kubernetes.Clientset, error) {
	// TODO
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(""))
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func (c *Cluster) createK8sClient(clusterId string) (clrclient.Client, error) {
	// TODO
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(""))
	if err != nil {
		return nil, err
	}
	return clrclient.New(config, clrclient.Options{})
}
