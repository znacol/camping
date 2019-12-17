package api

import (
	"github.com/spf13/cobra"
	pb "github.com/znacol/camping/go/proto"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

var CampingManager pb.CampingServiceClient

// InitClientCMD connects to the grpc service
func InitClientCMD(cmd *cobra.Command, args []string) error {
	if CampingManager != nil {
		return nil
	}
	svrAddr := cmd.Flag("server").Value.String()
	opts := grpc.WithInsecure()
	var err error
	conn, err = grpc.Dial(svrAddr, opts)
	if err != nil {
		return err
	}

	CampingManager = pb.NewCampingServiceClient(conn)
	return nil
}

// CloseCMD closes the connection
func CloseCMD(cmd *cobra.Command, args []string) error {
	if conn == nil {
		return nil
	}
	return conn.Close()
}

