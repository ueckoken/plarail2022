package servo

import (
	"errors"
	"testing"
	pb "ueckoken/plarail2022-external/spec"
)

func TestSend_trapResponseGrpcErr(t *testing.T) {
	// status 0 UNKNOWN; 1 SUCCESS; 2 FAILED
	testPatterns := []struct {
		grpcErr        error
		inResponseSync *pb.ResponseSync
		expectErrMsg   string
	}{
		{
			grpcErr:        errors.New("TEST ERROR"),
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(0)},
			expectErrMsg:   "gRPC Err: `TEST ERROR`; gRPC Response status is `UNKNOWN`",
		}, {
			grpcErr:        errors.New("TEST ERROR"),
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(1)},
			expectErrMsg:   "gRPC Err: `TEST ERROR`; gRPC Response status is `SUCCESS`",
		}, {
			grpcErr:        errors.New("TEST ERROR"),
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(2)},
			expectErrMsg:   "gRPC Err: `TEST ERROR`; gRPC Response status is `FAILED`",
		}, {
			grpcErr:        nil,
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(0)},
			expectErrMsg:   "gRPC Err: `%!w(<nil>)`; gRPC Response status is `UNKNOWN`",
		}, {
			grpcErr:        nil,
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(2)},
			expectErrMsg:   "gRPC Err: `%!w(<nil>)`; gRPC Response status is `FAILED`",
		},
	}

	for _, tp := range testPatterns {
		t.Run(tp.expectErrMsg,
			func(t *testing.T) {
				err := trapResponseGrpcErr(tp.inResponseSync, tp.grpcErr)
				if err == nil {
					t.Errorf("Expect err occured, but not occured.")
				} else if err.Error() != tp.expectErrMsg {
					t.Errorf("err format failed. Actual err is: %e", err)
				}
			},
		)
	}

	// Normal
	var grpcErr error
	rs := &pb.ResponseSync{Response: pb.ResponseSync_Response(1)}
	err := trapResponseGrpcErr(rs, grpcErr)
	if err != nil {
		t.Errorf("Expect err is NOT occured, but occured. : %e", err)
	}
}
