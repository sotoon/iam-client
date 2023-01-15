# bepa-client

# Project Name
> A simple yet powerful library that empowers you using bepa2 APIs.

## Table of Contents
* [General Info](#general-information)
* [Quick Start🏎️](#quick-start)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Architecture](#architecture)
* [Setup](#setup)
* [Project Status](#project-status)
* [Room for Improvement](#room-for-improvement)
* [Acknowledgements](#acknowledgements)
* [Support Notes](#support-notes)
* [External Links](#external-links)
* [Contact](#contact)


## General Information
- There are so many products that need to use Sotoon IAM Service Aka Bepa as their **Identity and Access Management** engine.
- Bepa-Client is a Golang Library empowering you can use it to control **the risk of API changes** and other support issues.
- It is under active development and support of Sotoon Integration Tribe.

## Technologies Used
- Golang :)

## Quick Start🏎 

Simply add bepa-client library latest stable version to your `go.mod` file:
> **Note!** Please check latest version tag [hear](https://git.cafebazaar.ir/infrastructure/bepa-client/-/tags).

### Installation

In to `go.mod` file add:
```mod
module git.cafebazaar.ir/infrastructure/kraken/commander

go 1.19

require (
	git.cafebazaar.ir/infrastructure/bepa-client v1.0.14
)
```
Resolve Golang library from private repository.

```bash
# you should have ssh access to gitlab repo of bepa-client
export GOPRIVATE=git.cafebazaar.ir 
go mod tidy
go mod vendor
```
### Initialization

Then simply use the client in your code:

```golang
import (
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
)

func SimpleBepaClientExample() {
    // ... initialize BEPA_URL, accessToken, defaultWorkspaceId, userId
    client, err := client.NewClient(accessToken, BEPA_URL, defaultWorkspaceId, userId)
    if err != nil {
        logger.Error("Cannot make a Bepa client.", zap.Error(err))
        // handle error or kill the process
        return
    }
}

func ReliableBepaClientExample() {
    // ... initialize serverUrls, accessToken, defaultWorkspaceId, userId, bepaTimeout
    client, err := return client.NewReliableClient(accessToken, serverUrls, defaultWorkspace, userUUID, bepaTimeout)
    if err != nil {
        logger.Error("Cannot make a Bepa client.", zap.Error(err))
        // handle error or kill the process
        return
    }
}

```

### Usage

See the [client.go](pkg/client/client.go) file to see the full list of API functions. The usage is so simple, just call the function with your intended parameters:

```golang
// Get Workspace Data by name
workspace, err := client.GetWorkspaceByName(workspaceName)

// authorize user with Sotoon IAM System
err := client.Authorize(identity, userType, action, object string)

// identify token with Sotoon IAM System
subject, err := client.Identify(token)
```

## Features

1. Almost all services of Sotoon IAM service!
2. Client-Side Fail-over.

## Architecture
Brief overview of projects deployment architecture.

![Bepa Client Failover](./docs/bepa_client_failover.png)


## Setup
Please refer to (Makefile)[Makefile].

## Project Status

Actively developing and supported.

## Room for Improvement
- Cache health-check result
- Developer API Guide
- Internal Mock Object in the library

## Support Notes
- Please don't hesitate to give any type of feedback to maintainers!

## External Links

- Commander: a successful example usage of the library with *Mocking and Testing*
- Link to other projects that use this project or are used by this project.
- Link to visualization tools and panels (e.g A reference to Turnilo would be beneficial in Account Manager project.)

## Acknowledgements
Please check the contributors section.

## Contact
Created by [@sib](https://www.sib.sotoon.ir/) team of Sotoon Integration Tribe!

