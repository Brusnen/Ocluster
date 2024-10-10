package src

import (
	pb "OCluster/grpc/src/orunner"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
)

func UploadFile(file_name string) {
	chunkSize := 1024 * 1024

	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		fmt.Println("Не удалось подключиться: %v", err)
	}
	file, err := os.OpenFile(file_name, os.O_RDONLY, 0755)
	file_size, err := os.Stat(file_name)
	defer file.Close()
	client := pb.NewOClusterClient(conn)
	stream, err := client.FileUploader(context.Background())
	ctx := stream.Context()
	buffer := make([]byte, chunkSize)
	done := make(chan bool)
	go func() {
		for {
			chunk, err := file.Read(buffer)
			if err != nil && err != io.EOF {
				fmt.Println(err)
			}
			if chunk == 0 {
				break
			}
			stream.Send(&pb.FileChunkRequest{Chunk: buffer, FileName: file_name, Filesize: uint64(file_size.Size())})
		}
		if err := stream.CloseSend(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("can not receive %v", err)
			}
			fmt.Println(resp)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			fmt.Println(err)
		}
		close(done)

	}()
	<-done
	fmt.Println("Finished Loading Data")
}
