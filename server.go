/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "./proto"
)

const (
	port = ":9090"
)


type server struct {
	pb.UnimplementedGhostServer
}

func (s *server) Diskcheck(ctx context.Context, in *pb.DiskRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Dir)
	return &pb.APIResponse{Responsemessage: "Diskcheck ", Responsecode: 0}, nil
}

func (s *server) Checkdefinition(ctx context.Context, in *pb.DefinitionRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	return &pb.APIResponse{Responsemessage: "Checkdefinition ", Responsecode: 0}, nil
}

func (s *server) Cutover(ctx context.Context, in *pb.Empty) (*pb.APIResponse, error) {
	return &pb.APIResponse{Responsemessage: "Cutover ", Responsecode: 0}, nil
}

func (s *server) Putpanicflag(ctx context.Context, in *pb.Empty) (*pb.APIResponse, error) {
	return &pb.APIResponse{Responsemessage: "Putpanicflag ", Responsecode: 0}, nil
}

func (s *server) Cleanup(ctx context.Context, in *pb.Empty) (*pb.APIResponse, error) {
	return &pb.APIResponse{Responsemessage: "Cleanup ", Responsecode: 0}, nil
}

func (s *server) Dryrun(ctx context.Context, in *pb.GhostRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	log.Printf("Received: "+ in.Statement)
	return &pb.APIResponse{Responsemessage: "Dryrun ", Responsecode: 0}, nil
}

func (s *server) ExecuteNohup(ctx context.Context, in *pb.GhostRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	log.Printf("Received: "+ in.Statement)
	return &pb.APIResponse{Responsemessage: "ExecuteNohup ", Responsecode: 0}, nil
}

func (s *server) Interactive(ctx context.Context, in *pb.InteractiveRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	log.Printf("Received: "+ in.Ghostcommand)
	return &pb.APIResponse{Responsemessage: "Interactive ", Responsecode: 0}, nil
}

func (s *server) Rowcount(ctx context.Context, in *pb.DefinitionRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	return &pb.APIResponse{Responsemessage: "Rowcount ", Responsecode: 0}, nil
}

func (s *server) Ibdsize(ctx context.Context, in *pb.IbdRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Dir)
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	return &pb.APIResponse{Responsemessage: "Ibdsize ", Responsecode: 0}, nil
}



func main() {
	log.Printf("port"+port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGhostServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}