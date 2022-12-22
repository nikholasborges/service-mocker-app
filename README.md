# service-mocker-app
a tool application that loads an local API with dinamically routed endpoints based on a JSON input file

## How to use:

1 - Clone the project

### if modifications in the code is made:

1.1 - Be sure to have go version 1.11+ installed.

1.2 - Run the command bellow to build the project into an executable:
```
go build .
```
2 - Create a `config.json`file into the /configs folder following the example in the `config.json.example` file with the endpoints to be mocked.

3 - Execute the `runServiceMock.sh` file.

4 - The Mock API will load dinamically the endpoints configured in the `config.json` file and expose them locally. Note: if no port is setted, the default port to be used is `:8080`

## About

Project created to help developers in the development of new APIs that communicate with other APIs.

## License

MIT © Copyright (c) 2022 nikholasborges
