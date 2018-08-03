package main

import (
	"time"
	"log"
	"net"
	"google.golang.org/grpc"
	"grpc-practice/primestream/primestreampb"
)
type server struct {}

func (*server) Primestream(req *primestreampb.NumberRequest, stream primestreampb.NumberService_PrimestreamServer) (error){
	N:=req.GetRequest()
	k:=int64(2)
	for N>1{
		var val int64;
		if (N%k)==0{
			val=k
			N=N/k
		}else{
			k=k+1
			continue
		}
		res:=&primestreampb.NumbermanyResponse{
			Response:val,
		}
		stream.Send(res)
		time.Sleep(1000*time.Millisecond)
	}
	return nil
}

func main(){
	lis,err:=net.Listen("tcp","0.0.0.0:50053")
	if err!=nil{
		log.Fatalf("Cannot creat listener %v",err)
	}
	s:=grpc.NewServer()
	primestreampb.RegisterNumberServiceServer(s, &server{})
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to creat server %v",err)
	}
}