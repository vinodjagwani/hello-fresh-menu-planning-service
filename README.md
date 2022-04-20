# HelloFresh Australia, Assignment (Software Engineer) by Vinod Jagwani

# Getting Started

### System requirements

1. Windows 10 or Mac or Linux or Ubuntu, minimum memory requirement is 8 gb

### Software requirements

1. Go 1.16 or higher
3. git
4. Intellij Idea (Optional) or any other IDE for source code

### Installation Instructions

1. Download Go from (version depends on os)  https://go.dev/dl/
3. Set Go home (GO_HOME)
4. Verify by typing on console go version
8. Download intelij idea https://www.jetbrains.com/go/

### Running locally instructions

There are different ways to run it locally:

1. Import project into intelij idea for go and run from idea itself by right click on main.go
2. Using command from console "go build" and "./hello-fresh-menu-planning-service" or "go run main.go"
3. Using command from console "make docker.start"

In order to check application whether application is up or not you can hit the following url from local
http://localhost:8080/swagger/index.html

### Running Integration Test

Using command from console "ENV_PROFILE=test make test.integration"

Note: Integration tests are in inside package internal/adoptors/e2e_test, This is 1 integration test for few APIs, the
idea is to demonstrate how testing can be done. so for others apis tests are ignored. (just for assignment purpose,
since real production required with 80% coverage)

### Generating swagger file

Using command from console "make generate.swagger"

### Postman collection is also included in the folder

### Service is using profile, so in order to run with different profile you can set profile

currently, there are 3 profiles: default, docker, azure, you can add more based on env by keeping file on root folder
such application-{profile_name}.yml

## Bonus (Deployed on Azure Public Cloud)

### Deployment Instructions on azure cloud

Service is deployed on Azure public cloud using docker, docker registry and postgres db (This is my public cloud)

### Azure cloud resources

1. resource-group: jagwani-azure-resource-group
2. location: centralus
3. docker-registry: jagwaniazuredockerregistry
4. docker-container: hellofreshaucontainer
5. postgres-db-instance: azure-dev-postgres

### Github workflow is attached with steps at .github/workflows/cicd-config.yml

so when you push the code on github it will automatically deploy on azure cloud which configured with above resources.

You can check with following swagger url.

http://azure-hellofresh-au-demo.eastus.azurecontainer.io:8080/swagger/index.html

Note: All apis are listed in swagger

### Setup Azure cloud from locally (Optional)

1. install azure cli using command: 'brew install azure-cli'

az login: 'login azure from command line'

create resource group by using following command
'az group create --name=jagwani-azure-resource-group --location=centralus'

create docker registry by using following command
'az acr create --resource-group jagwani-azure-resource-group --location centralus --name jagwaniazuredockerregistry
--sku Basic'

create docker container on azure by using following command az container create --resource-group
jagwani-azure-resource-group --name hellofreshaucontainer --image
jagwaniazuredockerregistry.azurecr.io/hello-fresh-menu-planning-service --dns-name-label azure-hellofresh-au-demo
--ports 8080

Get azure credentials by using following command 
az ad sp create-for-rbac --name "myApp" --role contributor \
--scopes /subscriptions/1cefad64-45a1-431d-af83-1478eb03a754/resourceGroups/jagwani-azure-resource-group \
--sdk-auth

You can use above credentials to set up GitHub action or any pipeline where applicable.

## Design Decision for apis and structures for apis

The project structure is following the ddd hexagonal concept where packages are structured in adoptors, ports, infra and
application For understanding hexagonal architecture please refer to the link or below link for getting bit of idea

http://tindaloscode.blogspot.com/2013/11/ddd-and-hexagonal-architecture.html

Here are the considerations while designing the apis:

1. All orm mapping is inside in the infra package (infra are secondary adaptor whos goal is to communicate with external
   systems, such database or message system)
2. Application folder contains ports, services and usecase
3. Use cases are Individual usecase or flows which will used by FE and these use cases will further serve to the rest
   api via ports
4. Ports can be in/out, but main goal here is to give access to the primary adoptors via ports such as rest apis
5. Every use case will be places in separate go file. so different developers can work on different use case indepandly
6. Service package will contain the business logic or any logic which can be used by use case or adoptors
7. There is rest package which contains inside adaptor package, rest apis are primary adaptor
8. There are dtos and mapper to convert request response objects into domain entities

## Language/Libraries are being used for developing this service in go

1. go (lang)
2. gin (web framework for go)
3. gorm (orm)
4. docker
5. docker-compose
6. postgres
7. swag
8. jwt-go
9. profiles

## Other Considerations

I have chosen gin over bee as web framework using following comparison

https://go.libhunt.com/compare-gin-vs-beego-astaxie

and GORM 2.0 over sqlx since it's rewritten from scratch, they improved many things (just for exploring purpose)
