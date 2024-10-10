package ocluster

import (
	"OCluster/grpc/src"
)

type OCluster struct {
}

func (o *OCluster) Run() {
	src.UploadFile("1.jpg")
	//db.InitDb()
	//go server.RunServer()

	select {}
}
