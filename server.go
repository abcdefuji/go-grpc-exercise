package main  
  
import (  
    "context"  
    "errors"
    "google.golang.org/grpc"
    "log"
    "net"
    cat "study-grpc/pb"  
)  
  
type myCatService struct{}  
  
func (s *myCatService) GetMyCat(ctx context.Context, message *cat.GetMyCatMessage) (*cat.MyCatResponse, error) {  
    switch message.TargetCat {  
    case "tama":  
        return &cat.MyCatResponse{  
            Name: "tama",  
            Kind: "Maine Coon",  
        }, nil  
    case "mike":  
        return &cat.MyCatResponse{  
            Name: "mike",  
            Kind: "Norwegian Forest Cat",  
        }, nil  
    default:  
        return nil, errors.New("Not Found YourCat..")  
    }  
}  
  
func main() {  
    port, err := net.Listen("tcp", ":1234")  
    if err != nil {  
        log.Println(err.Error())  
        return  
    }
    s := grpc.NewServer()  
    cat.RegisterCatServer(s, &myCatService{})  
    s.Serve(port)  
}