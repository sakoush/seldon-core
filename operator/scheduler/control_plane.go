/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed by
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

package scheduler

import (
	"context"
	"io"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	"github.com/seldonio/seldon-core/apis/go/v2/mlops/scheduler"
)

func (s *SchedulerClient) SubscribeControlPlaneEvents(ctx context.Context, grpcClient scheduler.SchedulerClient, namespace string) error {
	logger := s.logger.WithName("SubscribeControlPlaneEvents")

	stream, err := grpcClient.SubscribeControlPlane(
		ctx,
		&scheduler.ControlPlaneSubscriptionRequest{SubscriberName: "seldon manager"},
		grpc_retry.WithMax(SchedulerConnectMaxRetries),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(SchedulerConnectBackoffScalar)),
	)
	if err != nil {
		return err
	}

	for {
		event, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Error(err, "event recv failed")
			return err
		}
		logger.Info("Received event", "event", event)

		if err := s.handleStateOnReconnect(ctx, grpcClient, namespace); err != nil {
			logger.Error(err, "failed to handle state reconstruction")
			return err
		}

	}
	return nil
}
