package k8s

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"reflect"
	"sync"
	"testing"
)

func TestCluster_createK8sClient(t *testing.T) {
	type fields struct {
		mutex *sync.Mutex
		cache *sync.Map
	}
	type args struct {
		clusterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *kubernetes.Clientset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cluster{
				mutex: tt.fields.mutex,
				cache: tt.fields.cache,
			}
			got, err := c.createK8sClient(tt.args.clusterId)
			if (err != nil) != tt.wantErr {
				t.Errorf("createK8sClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createK8sClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCluster_get(t *testing.T) {
	type fields struct {
		mutex *sync.Mutex
		cache *sync.Map
	}
	type args struct {
		ctx       context.Context
		clusterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *kubernetes.Clientset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cluster{
				mutex: tt.fields.mutex,
				cache: tt.fields.cache,
			}
			got, err := c.get(tt.args.ctx, tt.args.clusterId)
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCluster_getK8sClientFromMap(t *testing.T) {
	type fields struct {
		mutex *sync.Mutex
		cache *sync.Map
	}
	type args struct {
		ctx       context.Context
		clusterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *kubernetes.Clientset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cluster{
				mutex: tt.fields.mutex,
				cache: tt.fields.cache,
			}
			got, err := c.getK8sClientFromMap(tt.args.ctx, tt.args.clusterId)
			if (err != nil) != tt.wantErr {
				t.Errorf("getK8sClientFromMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getK8sClientFromMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInstance1(t *testing.T) {
	tests := []struct {
		name string
		want *Cluster
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
