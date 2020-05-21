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
	"syscall"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os/exec"
	"os"
	"path/filepath"
)

const (
	port = ":9090"
)

type Tag struct {
    Table   string    `json:"table"`
    CreateTable string `json:"CreateTable"`
}

type Tag2 struct {
    TableRow   string    `json:"TableRow"`
}

type server struct {
	pb.UnimplementedGhostServer
}

func (s *server) Diskcheck(ctx context.Context, in *pb.DiskRequest) (*pb.APIResponse, error) {
	log.Printf("Diskcheck Received: "+ in.Dir)
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(in.Dir, &fs)
	if err != nil {
		return &pb.APIResponse{Responsemessage: "Diskcheck error", Responsecode: 1}, nil
	}
	return &pb.APIResponse{Responsemessage: strconv.FormatUint(fs.Bfree * uint64(fs.Bsize) / 1024 / 1024 / 1024, 10)+"G", Responsecode: 0}, nil
}

func (s *server) Checkdefinition(ctx context.Context, in *pb.DefinitionRequest) (*pb.APIResponse, error) {
	log.Printf("Checkdefinition Received: "+ in.Schemaname)
	log.Printf("Checkdefinition Received: "+ in.Tablename)
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/"+in.Schemaname)
	if err != nil {
		//panic(err.Error())
		return &pb.APIResponse{Responsemessage: "N/A", Responsecode: 1}, nil
	}
	defer db.Close()
	results, err := db.Query("show create table "+in.Schemaname+"."+in.Tablename)
	if err != nil {
		//panic(err.Error())
		return &pb.APIResponse{Responsemessage: "N/A", Responsecode: 1}, nil
	}
	var defintion string = "";
	for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.Table, &tag.CreateTable)
        if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			return &pb.APIResponse{Responsemessage: "N/A", Responsecode: 1}, nil
        }
                // and then print out the tag's Name attribute
		log.Printf(tag.Table)
		log.Printf(tag.CreateTable)
		defintion = tag.CreateTable
    }
	defer results.Close()
	return &pb.APIResponse{Responsemessage: defintion, Responsecode: 0}, nil
}

func (s *server) Cutover(ctx context.Context, in *pb.Empty) (*pb.APIResponse, error) {
	err := os.Remove("/tmp/ghost.postpone.flag")
	if err != nil {
		return &pb.APIResponse{Responsemessage: "Cutover ", Responsecode: 1}, nil
	}
	return &pb.APIResponse{Responsemessage: "Cutover ", Responsecode: 0}, nil
}

func (s *server) Putpanicflag(ctx context.Context, in *pb.Empty) (*pb.APIResponse, error) {
	cmd := exec.Command("touch", "/tmp/ghost.panic.flag")
	log.Printf("Running command touch /tmp/ghost.panic.flag ...")
	err := cmd.Run()
	log.Printf("Command finished with code: %v", err)
	return &pb.APIResponse{Responsemessage: "Putpanicflag ", Responsecode: 0}, nil
}

func (s *server) Cleanup(ctx context.Context, in *pb.Empty) (*pb.APIResponse, error) {
	array := [2]string{"/tmp/ghost.panic.flag","/tmp/ghost.postpone.flag"}
	for i := 0; i < len(array); i++ {
		err := os.Remove(array[i])
		log.Printf(array[i])
		if err != nil {
			return &pb.APIResponse{Responsemessage: "Cleanup ", Responsecode: 1}, nil
		}
	}
	files, err := filepath.Glob("/tmp/gh-ost.*.*.sock")
	if err != nil {
		//panic(err)
		return &pb.APIResponse{Responsemessage: "Clean up failed", Responsecode: 1}, nil
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			//panic(err)
			return &pb.APIResponse{Responsemessage: "Clean up failed", Responsecode: 1}, nil
		}
	}
	return &pb.APIResponse{Responsemessage: "Cleanup ", Responsecode: 0}, nil
}

func (s *server) Dryrun(ctx context.Context, in *pb.GhostRequest) (*pb.APIResponse, error) {
	log.Printf("Dryrun Received: "+ in.Schemaname)
	log.Printf("Dryrun Received: "+ in.Tablename)
	log.Printf("Dryrun Received: "+ in.Statement)
	info, err := os.Stat("/tmp/gh-ost."+in.Schemaname+"."+in.Tablename+".sock")
	if os.IsNotExist(err) {
		var binary string     = "/usr/bin/gh-ost"
		var conf string       = "/etc/mysql/debian.cnf"
		out, err := exec.Command(binary,
			"--max-load=Threads_running=50",             // 1
			"--critical-load=Threads_running=1500",      // 2
			"--chunk-size=500",                          // 3
			"--max-lag-millis=1500",                     // 4
			"--conf="+conf,                              // 5
			"--host=127.0.0.1",                          // 6
			"--allow-on-master",                         // 7  *
			"--approve-renamed-columns",                 // 8  *
			"--database="+in.Schemaname,                 // 9
			"--table="+in.Tablename,                     // 10
			"--alter="+in.Statement,                     // 11
			"--switch-to-rbr",                           // 12
			"--cut-over=default",                        // 13
			"--exact-rowcount",                          // 14
			"--concurrent-rowcount",                     // 15
			"--default-retries=120",                     // 16
			"--timestamp-old-table",                     // 17
			"--panic-flag-file=/tmp/ghost.panic.flag",   // 18
			"--postpone-cut-over-flag-file=/tmp/ghost.postpone.flag", // 19
			"--initially-drop-ghost-table",              // 20 *
			"--verbose").CombinedOutput()
		if err != nil {
			//log.Fatal(err)
		}
		log.Printf("Dryrun finished with code: %v", err)
        return &pb.APIResponse{Responsemessage: string(out), Responsecode: 0}, nil
    }else {
		if (info.IsDir()){
			log.Printf("Directory")
		}
		return &pb.APIResponse{Responsemessage: "Ghost socket file already exists.", Responsecode: 1}, nil
	}

}

func (s *server) execute(ctx context.Context, in *pb.GhostRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	log.Printf("Received: "+ in.Statement)
	return &pb.APIResponse{Responsemessage: "ExecuteNohup ", Responsecode: 0}, nil
}

func (s *server) ExecuteNohup(ctx context.Context, in *pb.GhostRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	log.Printf("Received: "+ in.Statement)
	info, err := os.Stat("/tmp/gh-ost."+in.Schemaname+"."+in.Tablename+".sock")
	if os.IsNotExist(err) {
		var binary string     = "/usr/bin/gh-ost"
		var conf string       = "/etc/mysql/debian.cnf"
		out, err := exec.Command(binary,
			"--max-load=Threads_running=50",             // 1
			"--critical-load=Threads_running=1500",      // 2
			"--chunk-size=500",                          // 3
			"--max-lag-millis=1500",                     // 4
			"--conf="+conf,                              // 5
			"--host=127.0.0.1",                          // 6
			"--allow-on-master",                         // 7  *
			"--approve-renamed-columns",                 // 8  *
			"--database="+in.Schemaname,                 // 9
			"--table="+in.Tablename,                     // 10
			"--alter="+in.Statement,                     // 11
			"--switch-to-rbr",                           // 12
			"--cut-over=default",                        // 13
			"--exact-rowcount",                          // 14
			"--concurrent-rowcount",                     // 15
			"--default-retries=120",                     // 16
			"--timestamp-old-table",                     // 17
			"--panic-flag-file=/tmp/ghost.panic.flag",   // 18
			"--postpone-cut-over-flag-file=/tmp/ghost.postpone.flag", // 19
			"--initially-drop-ghost-table",              // 20 *
			"--verbose",                                 // 21
			"--execute").CombinedOutput()
		if err != nil {
			//log.Fatal(err)
		}
		log.Printf("Execution finished with code: %v", err)
        return &pb.APIResponse{Responsemessage: string(out), Responsecode: 0}, nil
    }else {
		if (info.IsDir()){
			log.Printf("Directory")
		}
		return &pb.APIResponse{Responsemessage: "Ghost socket file already exists.", Responsecode: 1}, nil
	}
}

func (s *server) Interactive(ctx context.Context, in *pb.InteractiveRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	log.Printf("Received: "+ in.Ghostcommand)
	cmd := "echo "+in.Ghostcommand+" | nc -U /tmp/gh-ost."+in.Schemaname+"."+in.Tablename+".sock"
	out, err := exec.Command("bash","-c",cmd).CombinedOutput()
	if err != nil {
		return &pb.APIResponse{Responsemessage: string(out), Responsecode: 1}, nil
	}
	return &pb.APIResponse{Responsemessage: string(out), Responsecode: 0}, nil
}

func (s *server) Rowcount(ctx context.Context, in *pb.DefinitionRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/"+in.Schemaname)
	if err != nil {
		//panic(err.Error())
		return &pb.APIResponse{Responsemessage: "N/A", Responsecode: 1}, nil
	}
	defer db.Close()
	results, err := db.Query("SELECT TABLE_ROWS FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '"+in.Schemaname+"' AND TABLE_NAME = '"+in.Tablename+"'")
	if err != nil {
		//panic(err.Error())
		return &pb.APIResponse{Responsemessage: "N/A", Responsecode: 1}, nil
	}

	var tableRow string = "";
	for results.Next() {
        var tag Tag2
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.TableRow)
        if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			return &pb.APIResponse{Responsemessage: "N/A", Responsecode: 1}, nil
        }
                // and then print out the tag's Name attribute
		log.Printf(tag.TableRow)
		tableRow = tag.TableRow
    }
	defer results.Close()
	return &pb.APIResponse{Responsemessage: tableRow, Responsecode: 0}, nil
}

func (s *server) Ibdsize(ctx context.Context, in *pb.IbdRequest) (*pb.APIResponse, error) {
	log.Printf("Received: "+ in.Dir)
	log.Printf("Received: "+ in.Schemaname)
	log.Printf("Received: "+ in.Tablename)
	fi, err := os.Stat(in.Dir+"/"+in.Schemaname+"/"+in.Tablename+".ibd")
	if err != nil {
		//panic(err.Error())
		return &pb.APIResponse{Responsemessage: "N/A", Responsecode: 1}, nil
	}
	// get the size
	size := fi.Size()
	log.Printf(strconv.FormatInt(size,10))
	return &pb.APIResponse{Responsemessage: strconv.FormatInt(size / 1024 / 1024 / 1024,10)+"G", Responsecode: 0}, nil
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