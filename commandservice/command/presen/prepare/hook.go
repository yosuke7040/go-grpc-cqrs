package prepare

import (
	"context"
	"fmt"
	"github.com/yosuke7040/commandservice/infra/sqlboiler/handler"
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func CommandServiceLifecycle(lifecycle fx.Lifecycle, server *CommandServer) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if err := handler.DBConnect(); err != nil {
				panic(err)
			}

			port := 8082
			listerner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				return err
			}

			reflection.Register(server.Server)
			go func() {
				log.Printf("Command Server 開始 ポート番号: %v", port)
				server.Server.Serve(listerner)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Server.GracefulStop()
			log.Printf("Command Server 停止")
			return nil
		},
	})
}
