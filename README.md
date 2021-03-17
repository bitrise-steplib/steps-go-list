# Go list [![Bitrise Build Status](https://app.bitrise.io/app/e783d140de7df9d9/status.svg?token=RsL0h68Nd4D8wA_CtODchQ&branch=master)](https://app.bitrise.io/app/e783d140de7df9d9) [![Bitrise Step Version](https://img.shields.io/badge/version-0.10.1-blue)](https://www.bitrise.io/integrations/steps/go-list) [![GitHub License](https://img.shields.io/badge/license-MIT-lightgrey.svg)](https://raw.githubusercontent.com/bitrise-steplib/steps-go-list/master/LICENSE) [![Bitrise Community](https://img.shields.io/badge/community-Bitrise%20Discuss-lightgrey)](https://discuss.bitrise.io/)

This step runs the `go list ./...` command for you to list the go packages named by the import paths, starting from the current working directory.  
It can return a filtered package list in line with the exclude patterns.

## Examples

### List packages in the working directory excluding vendor/*

```yml
---
format_version: '8'
default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
project_type: other
workflows:
  release:
    steps:
    - git-clone: {}
    - go-list@0:
        inputs:
        - exclude: vendor/*
```

## Configuration

### Inputs

| Parameter | Description | Required | Default |
| --- | --- | --- | --- |
| exclude | Exclude patterns | - | "*/vendor/*" |

### Outputs

| Environment Variable | Description |
| --- | --- |
| BITRISE_GO_PACKAGES | List of go packages |

## Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/steps-go-list/pulls) and [issues](https://github.com/bitrise-steplib/steps-go-list/issues) against this repository. 

For pull requests, work on your changes in a forked repository and use the bitrise cli to [run your tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/)

### Creating your own steps

Follow [this guide](https://devcenter.bitrise.io/contributors/create-your-own-step/) if you would like to create your own step
