package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"net"
	"time"
	"wxprotocol"
	pb "wxprotocol/HelloMsg"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/*
Country Name (2 letter code) []:cn
State or Province Name (full name) []:zj
Locality Name (eg, city) []:hz
Organization Name (eg, company) []:wechat.com
Organizational Unit Name (eg, section) []:wechat.com
Common Name (eg, fully qualified host name) []:wechat.com
Email Address []:wechat@wechat.com
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build
*/

type server struct{}

func (s *server) HelloWechat(ctx context.Context, in *pb.HelloMsg) (*pb.HelloMsg, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "v1_LCEO_Code" && appid != "v1_xiami_code" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效")
	}
	if appkey != "v2_4d5d55d6fbb6ce87b0f53019fefebb89" && appkey != "v2_40ca8199d6aaa75d5468b875cd2b294f" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效")
	}

	if in.Cmd == 10000 {
		in.User = wxprotocol.NewUser()
		return in, nil
	}
	if in.Cmd > 0 {
		wxprotocol.Pack(in)
	}
	if in.Cmd < 0 {
		wxprotocol.UnPack(in)
	}
	return in, nil
}

func main() {

	flag.Parse()
	listenPort := flag.Arg(0)
	if listenPort == "" {
		listenPort = "12580"
	}

	l, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		log.Println("could not listen err", err)
		return
	}
	cert, err := tls.X509KeyPair(clientcrt, serverkey)
	if err != nil {
		log.Println("failed to parse certificate", err)
		return
	}
	log.Println("listen on port", listenPort)
	creds := credentials.NewServerTLSFromCert(&cert)
	srv := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServer(srv, &server{})
	if err := srv.Serve(l); err != nil {
		log.Println("grpc serve error", err)
	}
	time.Sleep(time.Second)
}
