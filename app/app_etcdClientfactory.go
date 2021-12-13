package app

import (
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"
	"time"
)

type etcdCientFactory struct {
	etcdClient *clientv3.Client
}

func NewEtcdClientFactory(addr []string) (*etcdCientFactory, error) {
	eTfc := &etcdCientFactory{}
	tlsInfo := transport.TLSInfo{
		CertFile:      `./config/etcd.pem`,
		KeyFile:       `./config/etcd-key.pem`,
		TrustedCAFile: `./config/ca.pem`,
	}
	config, err := tlsInfo.ClientConfig()
	if err != nil {
		panic(err.Error())
	}
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
		TLS:         config,
	}
	if clientTem, err := clientv3.New(conf); err == nil {
		eTfc.etcdClient = clientTem
	} else {
		return nil, err
	}
	return eTfc, nil
}
func (fact *etcdCientFactory) EtcdClient() *clientv3.Client {
	return fact.etcdClient
}
