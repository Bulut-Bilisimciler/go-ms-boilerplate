## Bulut-Bilisimciler-Go-Gin-Kafka-Nats-Microservice-Template

This is the simple boilerplate repository inspired by [THIS_REPO](https://github.com/Bulut-Bilisimciler/go-ms-boilerplate/tree/master)

### To Use This Template
- You need to clone repository to your local machine by using `git clone` command.
- You need to replaceAll [go.mod](go.mod) module name and all imported class in project yours (You can use vscode "SHIFT + CTRL + F" and replace all.). Old name is "github.com/Bulut-Bilisimciler/go-ms-boilerplate" and new name is your new project name like "github.com/MYNAME/MYSUPERAPP".
- You need to update all service-specific global variables like "APP_NAME, API_PREFIX, RN_PREFIX, FN_PREFIX" in [config.go file](./internal/config/config.go).
- Last, you need to replace HTTP Swagger documentation variables with your new name in [main.go Swagger comments](./main.go).

### Core Utilities
- [Uber/Fx](https://github.com/uber-go/fx) - for module based dependency injection. This project includes HTTP Server, Kafka Consumer and NATS Consumer services with Fx Injection.
- [Gin HTTP Server](https://gin-gonic.com/) - used for handling all HTTP requests.
- [Kafka Consumer](https://kafka.apache.org/) - implemented as a service for consuming messages from Kafka. You can use it for common Kafka operations like producing, consuming topics for your microservice.
- [Nats Req-Reply](https://nats.io) - used for intra microservices instant endpointless [request-reply](https://docs.nats.io/nats-concepts/core-nats/reqreply) communication.
- [Cron Scheduler](https://github.com/go-co-op/gocron/tree/v2) - used for application scheduled jobs. You can use it for any scheduled jobs like cron jobs. E.g. usage: push application metrics to push-gateway every 5 minutes, or nitfy admins for error queues or retry dead letter queue items.
- [Prometheus Metric Exporter](https://gin-gonic.com/) - used for export in-app-metrics from application like example files [1](), [2](), [3]() .
- [In-App-Event-Emitter](./pkg/infrastructure/event_emitter.go) you can use go built-in "sync" library as global-level and inject it to every service individually. In final, this emitter acts like a in-app event emitter. You can [emit any message from any in-app-service](./pkg/application/kafka_handlers/kafka_service.go) and [listen from its listener](./pkg/application/kafka_handlers/in_app_kafka_events.go).
- [Mage](https://magefile.org/) - used for [Makefile](https://www.gnu.org/software/make/manual/make.html) alternative in go. You can run your in-app-level operations or any level operations in go and you can run it with `mage` command like `make`. You can look our examples in [THIS_FILE](./tools/mage/mage.go). You can check available commands with `mage -l` command on directory (e.g. `cd ./tools/mage && mage -l`).
- [Taskfile](https://taskfile.dev/#/) - used to run out-of-code commands like "task gen-swag-docs" for documentation and any out-of-code level operations. For all taskfile commands, run `task -a`. ([task installation required](https://taskfile.dev/installation/))
- [DevContainer](https://code.visualstudio.com/docs/remote/containers) - used for auto ready VSCode development environment. You can run your development environment in a container with VSCode. You can check our [devcontainer.json](./.devcontainer/devcontainer.json) file for configuration.
- [PlantUML](https://plantuml.com/) - used for every service handlers documentation. You can create same name with ".go" files like ".plantuml" and explain your endpoints as sequence diagrams. You can copy-paste it to [PlantUML Server](https://www.plantuml.com/plantuml/uml/) and see your diagram. You can also use [VSCode PlantUML Extension](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml) for previewing your diagrams. (e.g. create_comment.go -> create_comment.plantuml docs file.)

### Directories

#### Core Directories
- [iac](./iac/) Infrastucture as Code (Terraform, Cloudformation etc.) or applciation definition yamls for kubernetes etc. or any other implementation specific yamls. (e.g. prometheus-metric)
- [internal/config](./internal/config/) for app-level .env.yaml file and configurations
- [pkg/infrastructure](./pkg/infrastructure/) any third-party infrastructure implementations (e.g. kafka, nats, redis, db etc.)
- [pkg/domain](./pkg/domain/) for domain application classes specific implementations (e.g. user, comment, post class and interface definitions etc.)
- [pkg/api](./pkg/api/) for handling all endpoint requests client facing implementation. This handlers is bridge for application user and [pkg/application](./pkg/application) app service core handlers. It is better to write middlewares and call service handlers here. (e.g. [http routes](./pkg/api/http_routes.go), [kafka consumer topic listeners](pkg/api/kafka_consumer_routes.go), [nats subject listener](pkg/api/nats_responder_routes.go) **definitions** etc.)
- [pkg/application](./pkg/application/) for handling all business logic and application core handlers. Core handlers of services should be here. For instance:
	- [Example HTTP Route](./pkg/api/http_routes.go) -> [Example Route Handler](pkg/application/http_handlers/get_user_info.go) - [Dont forget to add Gin Swagger Documentation for further documentation.](https://github.com/swaggo/gin-swagger)
	- [Example Kafka Route](./pkg/api/kafka_consumer_routes.go) -> [Example Route Handler](./pkg/application/kafka_handlers/consume_echo_incoming_text.go)
	- [Example Nats Route](./pkg/api/nats_responder_routes.go) -> [Example Route Handler](./pkg/application/nats_handlers/respond_echo_incoming_text.go)
	- [Example Job Schedule Route](./pkg/api/scheduled_jobs_routes.go) -> [Example Scheduled Job Handler](./pkg/application/scheduler_handlers/handle_echo_hello.go)

#### Utility Directories
- [.devcontainer](./.devcontainer/) - VSCode development container definition
- [.github](./.github/) - GitHub Actions workflows
- [.vscode](./.vscode/) - VSCode settings
- [.husky](./.husky/) - [Husky](https://typicode.github.io/husky/) Git Hooks
- [docs](./docs/) - documentation (for swagger outputs. individual service docs are in their own directories -plantuml-)
- [tools](./tools/) - tools for development (mage etc.)


### Libraries
- [Testing - testify](https://github.com/stretchr/testify) - for testing
- [Concurrency - conc](https://github.com/sourcegraph/conc) - for manage concurrency multi-threading
- [Config - viper](https://github.com/spf13/viper) - for configuration
- [Json - jsoniter](https://github.com/json-iterator/go) - for json encode/decode
- [Log - zerolog](https://github.com/rs/zerolog) - for logging
- [Gin-Log - gin-zerolog](https://github.com/gin-contrib/logger) - for gin http logging
- [Makefile - Mage](https://github.com/magefile/mage/tree/master) - go-level alternative for make and makefile (our use-case is in-app-level operations)
- [Taskfile - Task](https://taskfile.dev/#/) - for out-of-code makefile alternative (our use-case is out-of-code-level operations)
- [Git Hooks - Husky](https://dev.to/devnull03/get-started-with-husky-for-go-31pa) - for manage git hooks
- [Cache - go-redis/v9](https://redis.uptrace.dev/guide/go-redis.html) - for redis connection
- [Kafka - segmentio](https://github.com/segmentio/kafka-go) - for manage kafka connection
- [ReqReply - NATS](https://github.com/nats-io/nats.go) - for intra-communication nats connection
- [Cron - gocron](https://github.com/go-co-op/gocron/v2) - for in-app scheduled jobs


### Running locally

```bash
git clone https://github.com/Bulut-Bilisimciler/go-ms-boilerplate.git .
docker-compose up -d # waits all required services to be ready
cp ./internal/config/sample.env.yaml ./internal/config/.env.yaml # copy env file and update it with your own values
go run .
```

For generating swagger documentation, run:
```bash
task gen-swag-docs
```

for sync local compose postgres db with remote db, [there are several scripts in mage](tools/mage/mage_database.go), run:
```bash
# mage database:syncDB LOCAL_DB_DSN REMOTE_DB_DSN TO_LOCAL_SCHEMA_NAME TABLES_WITH_COMMA_STRING
# example usage below
mage database:syncDB postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=public postgres://remoteuser:remotepass@remoteurl:5432/remotedb?sslmode=disable&search_path=public remotesyncrev1 users,profiles,comments
```
