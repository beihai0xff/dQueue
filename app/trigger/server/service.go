package server

import (
	"fmt"
	"net"
	"net/http"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"

	"github.com/beihai0xff/pudding/api/gen/pudding/broker/v1"
	pb "github.com/beihai0xff/pudding/api/gen/pudding/trigger/v1"
	"github.com/beihai0xff/pudding/app/trigger/domain/cron"
	"github.com/beihai0xff/pudding/app/trigger/domain/webhook"
	"github.com/beihai0xff/pudding/app/trigger/pkg/configs"
	"github.com/beihai0xff/pudding/pkg/db/mysql"
	"github.com/beihai0xff/pudding/pkg/grpc/launcher"
	"github.com/beihai0xff/pudding/pkg/grpc/resolver"
	"github.com/beihai0xff/pudding/pkg/log"
	"github.com/beihai0xff/pudding/pkg/log/logger"
	"github.com/beihai0xff/pudding/pkg/utils"
)

const (
	// custom http endpoint prifix path
	httpPrefix = "/pudding/trigger"
)

var (
	// healthEndpointPath health check http endpoint path.
	healthEndpointPath = utils.GetHealthEndpointPath(httpPrefix)
	// swaggerEndpointPath Swagger ui http endpoint path.
	swaggerEndpointPath = utils.GetSwaggerEndpointPath(httpPrefix)

	// db is the MySQL database connection.
	db = mysql.New(configs.GetMySQLConfig())
	// schedulerClient is the scheduler grpc service client.
	schedulerClient broker.SchedulerServiceClient
)

func init() {
	// init grpc connection
	conn, err := grpc.Dial(
		configs.GetSchedulerConsulURL(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatalf("grpc Dial err: %v", err)
	}

	// TODO: close conn
	// defer conn.Close()

	// create scheduler service client
	schedulerClient = broker.NewSchedulerServiceClient(conn)
}

// RegisterLogger registers the logger to the resolver.
func RegisterLogger() {
	log.RegisterLogger(log.DefaultLoggerName, log.WithCallerSkip(1))
	log.RegisterLogger("gorm_log", log.WithCallerSkip(3))
	logger.GetGRPCLogger()
}

// RegisterResolver registers the service to the resolver.
func RegisterResolver(grpcPort, httpPort int) []*resolver.Pair {
	consulURL := configs.GetConsulURL()

	pairs := []*resolver.Pair{
		resolver.GRPCRegistration(pb.CronTriggerService_ServiceDesc.ServiceName,
			grpcPort, resolver.WithConsulResolver(consulURL)),
		resolver.GRPCRegistration(pb.WebhookTriggerService_ServiceDesc.ServiceName,
			grpcPort, resolver.WithConsulResolver(consulURL)),
		resolver.HTTPRegistration(healthEndpointPath,
			httpPort, resolver.WithConsulResolver(consulURL)),
	}
	return pairs
}

// StartServer starts the server.
func StartServer(grpcPort, httpPort int) (*grpc.Server, *health.Server, *http.Server) {
	grpcLis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	httpLis, err := net.Listen("tcp", fmt.Sprintf(":%d", httpPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer, healthcheck := launcher.StartGRPCService(grpcLis, startCronTriggerService, startWebhookTriggerService)
	httpServer := launcher.StartHTTPService(grpcLis, httpLis, healthEndpointPath, swaggerEndpointPath)
	return grpcServer, healthcheck, httpServer
}

func startCronTriggerService(server *grpc.Server, serviceName *string) error {
	// set serviceName
	// use string point to store serviceName, so that we can return it to the caller
	// this is not a good way to do this, but it works
	*serviceName = pb.CronTriggerService_ServiceDesc.ServiceName

	cronHandler := cron.NewHandler(cron.NewTrigger(db, schedulerClient))
	pb.RegisterCronTriggerServiceServer(server, cronHandler)

	return nil
}
func startWebhookTriggerService(server *grpc.Server, serviceName *string) error {
	// set serviceName
	// use string point to store serviceName, so that we can return it to the caller
	// this is not a good way to do this, but it works
	*serviceName = pb.WebhookTriggerService_ServiceDesc.ServiceName

	// register Trigger server
	webhookHandler := webhook.NewHandler(webhook.NewTrigger(db, schedulerClient))
	pb.RegisterWebhookTriggerServiceServer(server, webhookHandler)

	return nil
}
