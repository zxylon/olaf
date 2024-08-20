## Documentation
* [User Guide](https://github.com/zxylon/olaf/blob/main/docs/en/guide.md)
* [Architecture](https://github.com/zxylon/olaf/blob/main/docs/en/architecture.md)
* [Getting Started Tutorial](https://github.com/zxylon/olaf/blob/main/docs/en/tutorial.md)
* [Unit Testing](https://github.com/zxylon/olaf/blob/main/docs/en/unit_testing.md)


[Switch to Chinese version](https://github.com/zxylon/olaf/blob/main/docs/zh/guide.md)

# Olaf User Guide

Olaf is an application scaffolding based on Golang that helps you quickly build efficient and reliable applications. This guide will show you how to use Olaf to create and develop your applications.

## Installation

You can install Olaf using the following command:

```bash
go install github.com/zxylon/olaf@latest
```

> Tip: If `go install` is successful but the `olaf` command is not found, it is because the environment variable is not configured. You can configure the `GOBIN` directory in the environment variable.

## Creating a New Project

You can use the following command to create a new Golang project:

```bash
olaf new projectName
```

This command will create a directory named `projectName` and generate an elegant Golang project structure within it.


> Olaf provides 3 types of layouts:

* **Basic Layout**

The Basic Layout contains a minimalistic directory structure and is suitable for developers who are already familiar with Olaf projects.

* **Advanced Layout**

**Recommendation: We recommend beginners to choose the Advanced Layout first.**

The Advanced Layout includes many examples of using Olaf (e.g., db, redis, jwt, cron, migration, etc.), which is suitable for developers to quickly learn and understand the architectural ideas of Olaf.

* **Chat Layout**

Chat Layout is a Olaf based long link project template suitable for scenarios such as instant messaging and game development.

## Quick Start with Docker

If you want to quickly try out the Olaf advanced layout, we recommend using the following commands to start the project quickly:

```
cd ./deploy/docker-compose && docker compose up -d && cd ../../

go run ./cmd/migration

olaf run ./cmd/server
```

Alternatively, you can use the `make` command directly:

```
make bootstrap
```

## Creating Components

You can use the following commands to create components such as handler, service, repository, and model for your project:

```bash
olaf create handler user
olaf create service user
olaf create repository user
olaf create model user
```

These commands will create components named `UserHandler`, `UserService`, `UserRepository`, and `UserModel`, and place them in the correct directories.

If you want to create the corresponding components in a custom directory, you can do so as follows:
```bash
olaf create handler internal/handler/user/center
olaf create service internal/service/user/center
olaf create repository internal/repository/user/center
olaf create model internal/model/user/center
```

You can also use the following command to create all components (handler, service, repository, and model) at once:

```bash
olaf create all user
```

## Starting the Project

You can quickly start the project using the following command:

```bash
// Equivalent to go run ./cmd/server

olaf run
```

This command will start your Golang project and support hot-reload when files are updated.

## Compiling wire.go

You can quickly compile `wire.go` using the following command:

```bash
// Equivalent to cd cmd/server && wire
olaf wire
```

This command will automatically search for the `wire.go` file in your project and compile the required dependencies.

## Configuration File

### Starting with a Specific Configuration File
Olaf uses the Viper library to manage configuration files.

By default, it loads `config/local.yml`, but you can specify the configuration file path using environment variables or parameters.

```
// Linux or MacOS
APP_CONF=config/prod.yml olaf run

// Windows
set APP_CONF=config\prod.yml && olaf run
```
Alternatively, you can use the parameter approach: `go run ./cmd/server -conf=config/prod.yml`

### Reading Configuration Items

You can create a file named `local.yaml` in the `config` directory to store your configuration information. For example:

```yaml
data:
  mysql:
    user: root:123456@tcp(127.0.0.1:3380)/user?charset=utf8mb4&parseTime=True&loc=Local
  redis:
    addr: 127.0.0.1:6350
    password: ""
    db: 0
    read_timeout: 0.2s
    write_timeout: 0.2s
```

You can use dependency injection `conf *viper.Viper` to read the configuration information in your code:

```go
package repository

import (
	"context"
	"fmt"
	"github.com/zxylon/olaf-layout-advanced/pkg/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

func NewDB(conf *viper.Viper) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}


```
Tip: After performing dependency injection through parameters, don't forget to execute the `olaf wire` command to generate the dependency file.

## Logging

Olaf uses the Zap library to manage logs. You can configure the log in the `config` directory. For example:

```yaml
log:
  log_level: info
  encoding: json           # json or console
  log_file_name: "./storage/logs/server.log"
  max_backups: 30              # Maximum number of log file backups
  max_age: 7                   # Maximum number of days to keep files
  max_size: 1024               # Maximum size of each log file in MB
  compress: true               # Whether to compress the log files
```

You can use the following method to record logs in your code:

```go
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zxylon/olaf-layout-basic/internal/service"
	"github.com/zxylon/olaf-layout-basic/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

// ...

func (h *userHandler) GetUserById(ctx *gin.Context) {
	h.logger.Info("GetUserByID", zap.Any("user", user))
	// ...
}

// ...

```

## Database

Olaf uses the GORM library to manage databases. You can configure the database in the `config` directory. For example:

```yaml
data:
  mysql:
    user: root:123456@tcp(127.0.0.1:3380)/user?charset=utf8mb4&parseTime=True&loc=Local
  redis:
    addr: 127.0.0.1:6350
    password: ""
    db: 0
    read_timeout: 0.2s
    write_timeout: 0.2s
```

You can connect to the database using the following code:

```go
package repository

import (
	"github.com/zxylon/olaf-layout-advanced/internal/model"
)


type UserRepository interface {
	FirstById(id int64) (*model.User, error)
}
type userRepository struct {
	*Repository
}

func NewUserRepository(repository *Repository) *UserRepository {
	return &UserRepository{
		Repository: repository,
	}
}

func (r *userRepository) FirstById(id int64) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

```

It is important to note that `xxxRepository`, `xxxService`, `xxxHandler`, etc. in Olaf are implemented based on interfaces. This is known as **interface-oriented programming**, which can improve code flexibility, scalability, testability, and maintainability. It is a programming style highly recommended in the Go language.

In the above code, we wrote:
```
type UserRepository interface {
	FirstById(id int64) (*model.User, error)
}
type userRepository struct {
	*Repository
}
```
instead of directly writing:
```
type UserRepository struct {
	*Repository
}
```
> Tip: The unit tests in the Olaf advanced layout are based on the characteristics of `interface` for mock operations.

## Testing

Olaf uses libraries such as testify, redismock, gomock, and go-sqlmock to write tests.

You can run the test using the following command:

```bash
go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./.olaf/coverage.out ./test/server/...
go tool cover -html=./.olaf/coverage.out -o coverage.html

```

The above command will generate an HTML file named `coverage.html`. You can open it directly in a browser to view detailed unit test coverage.

## Conclusion

Olaf is a practical Golang application scaffolding that helps you quickly build efficient and reliable applications. We hope this guide can help you make better use of Olaf.