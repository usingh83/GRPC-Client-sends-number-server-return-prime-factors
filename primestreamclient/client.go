package main
import (
	"io"
	"fmt"
	"context"
	"log"
	"google.golang.org/grpc"
	"grpc-practice/primestream/primestreampb"
)
func main(){
	cc,err:=grpc.Dial("localhost:50053",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("failed to connect %v",err)
	}
	defer cc.Close()
	c:=primestreampb.NewNumberServiceClient(cc)
	req:=&primestreampb.NumberRequest{
		Request:120,
	}
	rstream,err:=c.Primestream(context.Background(),req)
	if err!=nil{
		log.Fatalf("Stream error",err)
	}
	for {
		msg,err:=rstream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("Error in msg",err)
		}
		fmt.Println("Stream is ", msg.Response)
	}

}