# GO Client for productive.io - Time Entry

This is an interactive CLI (Command Line Interface) application designed for entering your time into the productive.io platform. 

## Getting Started

To get your API token and Company ID follow the instructions here: https://developer.productive.io/index.html#header-api-token

### Build and Run

- Clone directory and run `go run .\cmd\productive-go-client\` from the directory root


1. On your first run, the application will prompt you for your details, including your API key and organization ID.

2. Follow the prompts to enter the required information.

3. Once configured, you can use the application to enter your time data into productive.io from your terminal.

4. If configuration is present, you can run the client non-interactively using flags
```
  -date string
        Date for the operation in YYYY-MM-DD format (default "2024-05-10")
  -hours int
        Hours component of the time (default 8)
  -minutes int
        Minutes component of the time
  -notes string
        Date for the operation in YYYY-MM-DD format
  -serviceId string
        ID of the service
```

## Other
- To edit or reset the config, use the config.json file created in the app root directory on first run.
