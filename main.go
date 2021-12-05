package main

import (
	"log"
	"strconv"
	"os"
)

func main() {

	log.Printf("Starting the connector and reading properties in the properties.ini file")
	var prop map[string]string

	if len(os.Args) <= 1 { 
		/* Reading properties from ./properties.ini */
		prop, _ = ReadPropertiesFile("./properties.ini")
	} else  {
		prop, _ = ReadPropertiesFile(os.Args[1])
	}
	
	port, _ := strconv.Atoi(prop["GreenplumPort"])
	batch, _ := strconv.Atoi(prop["Batch"])

	log.Printf("Properties read: Connecting to the Grpc server specified")

	/* Connect to the grpc server specified */
	gpssClient := MakeGpssClient(prop["GpssAddress"], prop["GreenplumAddress"], int32(port), prop["GreenplumUser"], prop["GreenplumPassword"], prop["Database"], prop["SchemaName"], prop["TableName"])
	gpssClient.ConnectToGrpcServer()

	log.Printf("Connected to the grpc server")

	log.Printf("delegating to pipe client")

	pipe := makePipeClient(prop["PipePath"], gpssClient, batch, prop["Delim"])
	pipe.readPipe()

}
