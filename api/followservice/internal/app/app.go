package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"social-network/followservice/pkg/followservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *ServiceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	followservice.RegisterFollowServer(a.grpcServer, a.serviceProvider.FollowImpl(ctx))
	return nil
}

func (a *App) Run() error {
	// wg := sync.WaitGroup{}
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	if err := a.runGRPCServer(); err != nil {
		log.Fatal(err)
	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	if err := a.runHTTPServer(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	return nil
}

func (a *App) runGRPCServer() error {
	list, err := net.Listen("tcp", a.serviceProvider.grpcConfig.Address())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GRPC follow service server is running on %s", a.serviceProvider.grpcConfig.Address())
	if err := a.grpcServer.Serve(list); err != nil {
		log.Fatal(err)
	}
	return nil
}

// func (a *App) runHTTPServer() error {
// 	return a.httpServer.ListenAndServe()
// }
