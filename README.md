# E-Core test

## Dependencies

[Baloo](https://github.com/h2non/baloo)  - For end-to-end test

[Env](https://github.com/nicholasjackson/env) - To handle environment variables

[Go-Hclog](https://github.com/hashicorp/go-hclog) - Simple logger package, key/value based

[Testify](https://github.com/stretchr/testify)  - For tests :D

## Requirements to run the project inside the Docker
- docker [go to the docs](https://docs.docker.com/engine/install/ubuntu/)
- docker-compose [go to the docs](https://docs.docker.com/compose/install/)

### Executing application

###### It`s necessary to be at the root of the project.

Using docker-compose. 

```bash
docker-compose -f docker/docker-compose.yml up --build
```

Or

If you have [make](https://helpmanual.io/man1/make/) installed 
```bash
make run
```

Or just execute via Go

```
go run .
```

### Executing the operations

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```curl
   curl -F 'file=@matrix.csv' "localhost:8080/echo"
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
   curl -F 'file=@matrix.csv' "localhost:8080/invert"
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
   curl -F 'file=@matrix.csv' "localhost:8080/flatten"
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    curl -F 'file=@matrix.csv' "localhost:8080/sum"
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    curl -F 'file=@matrix.csv' "localhost:8080/multiply"
    ``` 

### Executing the tests

I divided the tests between end-to-end and unit test.

#### Running unit test

###### runs race detector, stop on the first test that fails and generates the coverage file

```shell
go test -mod=vendor -v -race -failfast -cover -coverprofile=coverage.out ./services ./handlers || exit 1
```

or simple

```
make test
```

To see the coverage as HTML

```shell
go tool cover -html=coverage.out
```

or 

```shell
make coverage-html
```

#### Running end-to-end

It`s configured to run inside of the container as the tests need the application to be up. Docker-compose will run the app and will run another container with the source code, via entry point a .sh file will be executed in order to execute the tests

```bash
docker-compose -f docker/e2e/docker-compose.yml down
docker-compose -f docker/e2e/docker-compose.yml build
docker-compose -f docker/e2e/docker-compose.yml up --abort-on-container-exit --remove-orphans
docker-compose -f docker/e2e/docker-compose.yml stop
docker-compose -f docker/e2e/docker-compose.yml rm -f
```

simplifying with make

```bash
make e2e-test
```
 