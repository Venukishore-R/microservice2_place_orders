package middlewares

import (
	"context"
	"fmt"
	slog "log"

	"github.com/Venukishore-R/microservice1_auth/protos"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var logger log.Logger

func Mid() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			slog.Print("context 111", ctx)
			var conn *grpc.ClientConn
			conn, err = grpc.NewClient(":5005", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return nil, err
			}

			defer conn.Close()

			client := protos.NewUserServiceClient(conn)
			tok := (ctx.Value(jwt.JWTContextKey)).(string)

			md := metadata.New(map[string]string{"authorization": "Bearer " + tok})
			ctx = metadata.NewOutgoingContext(ctx, md)
			slog.Print("context ", ctx)

			var header metadata.MD

			response, err = client.Authenticate(ctx, &protos.Empty{}, grpc.Header(&header))

			if err != nil {
				slog.Printf("err1 %v", err)
				return nil, err
			}

			resp := response.(*protos.AuthenticateUserResp)
			if resp.Status != 200 {
				slog.Printf("err2 %v", fmt.Errorf("not 200 code"))

				return nil, err
			}

			ctx = context.WithValue(ctx, "Name", resp.Name)
			ctx = context.WithValue(ctx, "Email", resp.Email)
			ctx = context.WithValue(ctx, "Phone", resp.Phone)

			return next(ctx, request)
		}
	}
}
