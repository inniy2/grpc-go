# go

# download go
# https://golang.org/dl/

wget https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz

tar xvfP go1.14.3.linux-amd64.tar.gz

sudo mv go /usr/local/

sudo chown -R root.root /usr/local/go

vim ~/.bashrc

export PATH=$PATH:/usr/local/go/bin

source ~/.bashrc

# Get protoc-gen-go
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/go-sql-driver/mysql
go get gopkg.in/yaml.v2

vim ~/.bashrc

export PATH=$PATH:/home/sangsun/go/bin

source ~/.bashrc


# protoc
# protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld


protoc -I proto/ proto/ghost.proto --go_out=plugins=grpc:proto


# build
go build server.go


# For local test
sudo lxc file push server u2/root/server

# config.yaml
# Server configurations
server:
  port: 9090

# database configurations
database:
  user: root
  pass: 12345678
