# Web apps

## 1. quick start

- running `./scripts/local_setup.sh` to setup the tools required by the project. After setup, please update `.envrc` and `.envrc.test` based on your own needed.
- running `make openapi_to_proto` to **generate** the api required proto file.
- running `make regen_proto` to **regenerate** all the golang code from `proto` files.
- running `go mod tidy` to get all the dependency.
- Check `configs` folder and `.envrc` file, especially `.envrc` file, it contains several env var, you will need modify it such as `DATABASE_URL` to your own config.
- running `make migration_create migration_name=[MigrationName]` to create migration
- running `make migration_up` to do all the migrations for the database.
- running `make regen_gorm` to **generate** all the database models and query.
- running `make wire` to **generate** all the required dependency injection files.
- running `make mockery` to **generate** all the testing required mockery package.
- All above generated command can be combined with `make generate`.
- For testing, can running `make generate ENVRC_FILE=.envrc.test` to specify we use `.envrc.test` instead of `.envrc`.
- running `make` to start the project in dev env. 🚀

## 2. Project structure

The generated project is based on golang best practice file structure.

- `pkg` contains different packages that you can share with other projects.
- `internal` contains packages that only for this project, but without core business logic. You can put the http/grpc handler in this folder, but the core business service should in the root `app` folder.
- `app` folder contains core business service. This is the core service, it should be contains only this project's business logic.
- `cmd` folder contains several starting command. eg, start the http server.
- `openapi` folder contains the `openapi/swagger` documentation. And it also contains the file that generated by documentation. We are using a different flow here, please check the next section.
- `gorm` contains all the gorm generated packages. It also contains a generate go file which used to generate the gorm model and query.
- `configs` contains the the yaml config file that needed by starting the app.
- `scripts` folder contains some helpful scripts, but typically you don't use the script directly, instead, `Make file` define several command to use these scripts.
- `.envrc` is required as it contains several **env** var that required for the app.

## 3. Some details

- For http server, this project is based on `Openapi` first approach. When you add a new API or update an existing API, you should first update `./openapi/docs/api.yaml`, and you generate all the proto files and golang related code to have your request and response model.
- For grpc server, you just follow the standard way. You create proto file by yourself, and then generated all the go files to use.
- For `orm`, we use `gorm` here. But we are only using the [`gen`](https://gorm.io/gen/) approach instead of the traditional orm approach. When you need do some db migration, you should first do the migration, you can use the `make migration-create migration_name=[migrationName]` to generate migration first. And then, write your sql in the migration file. After that, run `make migration up` to do the migration.(The underlying is using [`goose`](https://github.com/pressly/goose) to do all these migration work.) After this, you can run `make regen_gorm` to **generate** the gorm model and query.
