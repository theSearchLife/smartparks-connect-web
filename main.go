package main

import (
	"github.com/SmartParksOrg/smartparks-connect-web/device_template"
	"github.com/SmartParksOrg/smartparks-connect-web/grpc_client"
	"github.com/SmartParksOrg/smartparks-connect-web/web"
)

func main() {
	templateMananger := device_template.NewTemplateManager()
	grpcClient := grpc_client.NewGrpcClient()
	web.StartHttpServer(grpcClient, templateMananger)

}
