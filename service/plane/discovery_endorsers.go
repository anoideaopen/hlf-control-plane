package plane

import (
	"context"

	pb "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) DiscoveryEndorsers(ctx context.Context, req *pb.DiscoveryEndorsersRequest) (*pb.DiscoveryEndorsersResponse, error) {
	peers, err := s.discCli.GetEndorsers(ctx, req.ChannelName, req.ChaincodeName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get peers: %v", err)
	}
	return &pb.DiscoveryEndorsersResponse{Result: peers}, nil
}
