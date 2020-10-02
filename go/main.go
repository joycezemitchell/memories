package main

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	memoriespb "allyapps.com/memories/proto"
	server "allyapps.com/memories/server"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	memoriespb.RegisterMemoriesServiceServer(s, &server.Server{})

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// This is where GRPC Gateway comes into play - START ________________________
	conn, _ := grpc.Dial("localhost:10000", grpc.WithInsecure())

	jsonpb := &runtime.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
	)

	memoriespb.RegisterMemoriesServiceHandler(context.Background(), gwmux, conn)
	// This is where GRPC Gateway comes into play - END ________________________

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		gwmux.ServeHTTP(w, r)
		return
	})

	http.ListenAndServe(":13000", nil)

}
