# GeneratedPersonApp

## What is this?

This is a generated jiffy application based on the jiffy [simpleSingleEntityModel.json model.](https://github.com/1414C/jiffy/blob/master/support/testing_models/simpleSingleEntityModel.json)

## What to do

Open a terminal window, switch to a folder under $GOPATH/src/ and perform the following:
```bash
    go get -u github.com/1414C/GeneratedPersonApp
```

Edit the *.dev.config.json* file and verify the following:

1. Set the "external_address" value to an ipv4 address that jiffy can bind the service end-points to.
1. Edit the "database" sub-keys to aim at the database of your choosing as per the [Jiffy configuration documentation.](https://1414c.github.io/jiffy/usage/us-content-a/)
1. Edit the "logging" sub-keys as per your requirements.  The initial settings are probably fine to start with.
1. Everything else should be fine.

# Running the Application

Review the content of the Jiffy docs starting at the [Let's Run Something](https://1414c.github.io/jiffy/getting-started/gs-content-d/) section.  Switch to the folder you installed the application source code in and run the following:

```bash
    go run main.go -dev -rs
```

Once the application is running, continue to follow the tutorial in the [Getting Started](https://1414c.github.io/jiffy/getting-started/level-two/) section to start testing the application via Postman or your favorite test tool.

Subsequent starts of the application should omit the *-rs* flag as follows:

```bash
    go run main.go -dev
```