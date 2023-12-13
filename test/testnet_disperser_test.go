package integration_test

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDialAndDisperseBlob(t *testing.T) {
	// Using insecure credentials for testing purposes
	conn, err := grpc.Dial("disperser-goerli.eigenda.xyz:443", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := disperser.NewDisperserClient(conn)
	_, err = client.DisperseBlob(
		context.Background(),
		&disperser.DisperseBlobRequest{
			Data: []byte("0"),
			SecurityParams: []*disperser.SecurityParams{
				{QuorumId: 0, AdversaryThreshold: 25, QuorumThreshold: 50},
			},
		},
	)
	if err != nil {
		t.Fatalf("DisperseBlob failed: %v", err)
	}
	// You can add more assertions here if needed
}
