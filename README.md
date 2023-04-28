# Open Data Service
_By Elon Society of Computing Project Team_

Open Data Service is a service that provides a REST API for accessing data from Elon University. [Get Started](https://ods.elon.edu)

## Introduction

This repository contains the source code for both the backend and frontend of ODS. Located in [backend](https://github.com/elonsoc/ods/tree/main/backend) is the backend code, which is written in Go. Located in [/frontend](https://github.com/elonsoc/ods/tree/main/frontend) is the frontend code, which is written in Typescript.

This project is to learn and employ production level software engineering practices. This includes:
- [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection)
- [Test Driven Development](https://en.wikipedia.org/wiki/Test-driven_development)
- [Continuous Integration](https://en.wikipedia.org/wiki/Continuous_integration)
- [Continuous Deployment](https://en.wikipedia.org/wiki/Continuous_deployment)
- [Infrastructure as Code](https://en.wikipedia.org/wiki/Infrastructure_as_code)
- [Pull Requests](https://en.wikipedia.org/wiki/Pull_request)
- [Code Reviews](https://en.wikipedia.org/wiki/Code_review)

In modern software engineering, not only is it important to write code that works, but it is also important to write code that is maintainable as a team.

## Getting Started
This repository is a [monorepo](https://en.wikipedia.org/wiki/Monorepo) that contains both the frontend and backend code. To get started, you will need to clone this repository and install the dependencies for both the frontend and backend. Although both folders have their own READMEs, we will go over the steps here.

The first step is to clone the repository. You can do this by running the following command:
```bash
git clone https://github.com/elonsoc/ods.git
```

### The Backend
In the backend we have a file called `Makefile` that contains all the commands you need to get started. The first step is to install the dependencies. You can do this by running the following command:
```bash
make init
```
Which will run `go get .`. This will install all the dependencies for the backend.

From there, you can run the backend by running the following command:
```bash
make run
```
which will run `go run .` with example environment variables that connect directly to the expected Docker containers ran by `docker-compose`.

If the example environment variables do not work for you, you can create a `.env` file in the root of the backend folder and add the following environment variables:
```bash
# The port to run the server on
PORT=8080
...
```
and run `go run .` again.

### The Frontend

The frontend requires [pnpm](https://pnpm.io). You can install pnpm through npm by running the following command:
```bash
npm install -g pnpm
```

From there, you can install the dependencies by running the following command:
```bash
pnpm install
```
and run the frontend by running the following command:
```bash
pnpm run dev
```


## Toplevel Commands
In the root of the project, we have a makefile that leverages Docker Compose to run the entire project. You can run the following command to run the entire project:
```bash
make up
```
which will run `docker-compose up` with the `docker-compose.yml` file in the [infra](https://github.com/elonsoc/ods/tree/main/infra) folder.

## Observability and Monitoring

This project uses the following technologies to enable observability and monitoring:
- [Grafana](https://grafana.com)
- [Loki](https://grafana.com/loki)
    - Logging in the backend provided by [Logrus](https://github.com/sirupsen/logrus)
- [Prometheus](https://prometheus.io/)
    - (Though in the future we might consider [VictoriaMetrics](victoriametrics.com)!)
- [statsd](https://github.com/statsd/statsd)
