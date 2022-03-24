# Octopus Report

 Octopus Report runs audit reports against an Octopus Deploy server.

## Requirements

### Octopus Deploy 3.x
The Octopus REST API has breaking changes starting with 4.x, so this program only works with 3.x.

### Go 1.17.3
[Download Go](https://golang.org/dl/)

## How to Build
```bash
go build -o octoreport.exe main.go
octoreport.exe
```

## How to Use

You need to specify the Octopus server to query, either in a config.yml file, in environment variables, or in command line parameters.

config.yml

```yaml
octopus:
   address: https://samples.octopus.app
   apikey: API-GUEST
```
```bash
octoreport
```

environment variables

```bash
export OCTOPUS_ADDRESS=https://samples.octopus.app
export OCTOPUS_APIKEY=API-GUEST
octoreport
```

command line parameters

```bash
octoreport octopus.address=https://samples.octopus.app octopus.apikey=API-GUEST
```

### Setting Log Level

```bash
export LOG_LEVEL=Debug
```

### Command Line Options
```bash
octoreport -help
```

| Option         | Description |
| -------------- | ----------- |
| -help          | display usage of this program |
| -info          | display information about this program |
| -config        | display program configuration that was loaded |
| -server        | display the Octopus server version information |
| -report        | display a report of all teams and their users/roles/environments/project groups/projects |
| -teams         | display all teams |
| -users         | display all users |
| -roles         | display all user roles |
| -environments  | display all environments |
| -projects      | display all projects |
| -projectgroups | display all project groups |
| -tenants       | display all tenants |
| -channels      | display all channels |
| -feeds         | display all feeds |
| -lifecycles    | display all lifecycles |
| -machines      | display all machines |
| -machineroles  | display all machineroles |

## API Documentation

[Swagger UI](https://samples.octopus.app/swaggerui/index.html)

[Octopus REST API](https://octopus.com/docs/api-and-integration/api)

[Octopus Deploy API](https://github.com/OctopusDeploy/OctopusDeploy-Api/wiki)

## Public Sample Octopus Deploy Server

[Demo Octopus Deploy](https://samples.octopus.app)

The following credentials are needed
| Name     | Value |
| -------- | ----- |
| username | guest |
| password | guest |
| API key  | API-GUEST |

## Example API Calls

Hit these URLs in your browser to test access to the Octopus Deploy server.

[Get Server Information](https://samples.octopus.app/api?apikey=API-GUEST)  
[Get Teams](https://samples.octopus.app/api/teams?apikey=API-GUEST)  
[Get Users](https://samples.octopus.app/api/users?apikey=API-GUEST)  
[Get UserRoles](https://samples.octopus.app/api/userroles?apikey=API-GUEST)  
[Get Environments](https://samples.octopus.app/api/environments?apikey=API-GUEST)  
[Get Projects](https://samples.octopus.app/api/projects?apikey=API-GUEST)  
[Get Project Groups](https://samples.octopus.app/api/projectgroups?apikey=API-GUEST)  
[Get Tenants](https://samples.octopus.app/api/tenants?apikey=API-GUEST)  
[Get Channels](https://samples.octopus.app/api/channels?apikey=API-GUEST)  
[Get Feeds](https://samples.octopus.app/api/feeds?apikey=API-GUEST)  
[Get Lifecycles](https://samples.octopus.app/api/lifecycles?apikey=API-GUEST)  
[Get Machines](https://samples.octopus.app/api/machines?apikey=API-GUEST)  
[Get MachineRoles](https://samples.octopus.app/api/machineroles/all?apikey=API-GUEST)  
