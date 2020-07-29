package main

import (
	"gomicrogrpcstudy/serviceimpl"
	"gomicrogrpcstudy/services"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("192.168.31.82:12379"),
	)

	service := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":9071"),
		micro.Registry(reg),
	)

	service.Init()
	prods := &serviceimpl.ProdService{}
	services.RegisterProdServiceHandler(service.Server(), prods)
	service.Run()
}
