package main

import (
	"io"
	"log"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "grpc-web-service/data"
)

const (
	address = "localhost:50051"
)

// createData calls the RPC method CreateData of DataServer
func createData(client pb.DataClient, data *pb.DataRequest) {
	resp, err := client.CreateData(context.Background(), data)
	if err != nil {
		log.Fatalf("Could not create Data: %v", err)
	}
	if resp.Key < 1 {
		log.Printf("%s already exists:", resp.Data)
	}
	if resp.Success {
		log.Printf("A new Data has been added for: %s", resp.Data)
	}
}

// getData calls the RPC method GetData of DataServer
func getData(client pb.DataClient, filter *pb.DataFilter) {
	// calling the streaming API
	stream, err := client.GetData(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get data: %v", err)
	}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetData(_) = _, %v", client, err)
		}
		log.Printf("Data: %v", data)
	}
}

// deleteData calls the RPC method DeleteData of DataServer
func deleteData(client pb.DataClient, data *pb.DeleteDataRequest) {
	resp, err := client.DeleteData(context.Background(), data)
	if err != nil {
		log.Fatalf("Could not Data: %v", err)
	}
	if resp.Success {
		log.Printf("%s has been deleted", resp.Data)
	}
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new DataClient
	client := pb.NewDataClient(conn)

	data := &pb.DataRequest{
		Data:  "gRPC test2",
	}

	// Create a new data
	createData(client, data)

	data = &pb.DataRequest{
		Data:  "I am data also",
	}

	// Create a new data
	createData(client, data)
	
	// Filter with an empty Keyword
	filter := &pb.DataFilter{Keyword: ""}
	getData(client, filter)
	
	delete := &pb.DeleteDataRequest{Data:  "I am data also"}
	
	// Delete data matching given key 
	deleteData(client, delete)
}