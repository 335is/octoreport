# Octopus Report

 Octopus Report runs audit reports against an Octopus Deploy server.

## Requirements

### Go 1.17.3

[Download Go](https://golang.org/dl/)

## How to Use

Requires two environment variables pointing to the Octopus Deploy server.

```bash
export OCTOMON_OCTOPUS_ADDRESS=https://samples.octopus.app
export OCTOMON_OCTOPUS_APIKEY=API-GUEST
go run main.go
```

### Setting Log Level

```bash
export LOG_LEVEL=Debug
```

## Example Octopus Deploy Server

[Demo Octopus Deploy](https://samples.octopus.app)

username: guest

password: guest

API key: API-GUEST

## API Documentation

[Swagger UI](https://samples.octopus.app/swaggerui/index.html)

[Octopus REST API](https://octopus.com/docs/api-and-integration/api)

[Octopus Deploy API](https://github.com/OctopusDeploy/OctopusDeploy-Api/wiki)

## Example API Calls

Hit these URLs in your browser to test access to the Octopus Deploy server.

[Server Information](https://samples.octopus.app/api?apikey=API-GUEST)

[Get Users](https://samples.octopus.app/api/users?apikey=API-GUEST)

[Get Teams](https://samples.octopus.app/api/teams?apikey=API-GUEST)

[Get ScopedUserRoles](https://samples.octopus.app/api/teams/teams-everyone/scopeduserroles?apikey=API-GUEST)

[Get Environments](https://samples.octopus.app/api/environments?apikey=API-GUEST)

[Get Projects](https://samples.octopus.app/api/projects?apikey=API-GUEST)
