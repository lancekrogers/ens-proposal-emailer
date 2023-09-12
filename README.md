# Tally backend take-home

## Overview

This project monitors the ENS Governance contract and sends an email notification when a new proposal is submitted to the contract.

## Prerequisites

- Go 1.21+
- SMTP Email Account (Gmail recommended)
- Make (for running Makefile commands)

## Installation

1. Clone the repository:
```zsh
git clone git@github.com:lancekrogers/tally-take-home.git
```
2. Navigate to the project directory

## Configuration

Copy exp.config.yaml to config.yaml and fill in the below empty string fields.  If you are using gmail for your smtp account you will need to enable the gmail api and create an app password

```yaml
TallyApi:
  Key: ""
  Url: "https://api.tally.xyz/query"

EmailSettings: 
  DestinationAddress: ""
  SenderUsername: ""
  SenderPassword: ""
  Host: "smtp.gmail.com"
  Port: "587"

ENSGovernanceContract: 
  Address: "0x323A76393544d5ecca80cd6ef2A560C6a395b7E3"
```

## Running the service

1. Build the project:
```zsh
make build
```
2. Run the project:
```zsh
make run
```

On your first run the latest proposal should be added to your database and an email alert will be sent out.  Since proposals are sparse, there will not be many logs. Run `make clean` to clear the database and start from scratch.


## Design Overview

```zsh
├── cmd
│   └── main.go
├── config.yaml
├── internal
   ├── email
   │   ├── email.go
   │   ├── email_test.go
   │   └── smtp.go
   ├── monitoring
   │   └── monitor.go
   ├── setup.go
   ├── store
   │   └── boltdb_store.go
   ├── tally
   │   └── tally.go
   └── utils
       └── config.go
```

### Package Structure

- `email:` Email notification utility
    - Utilizes the net/smtp package
- `monitoring:` Responsible for monitoring the ENS governance contract
    - Checks for new proposals at regular intervals and triggers email notifications
- `store:` Manages the database interactions
    - Uses BoltDB to store the block number of the latest processed proposal
- `tally:` Wrapper interface for the tally api
- `utils:` Utility functions for configuration loading and other shared functionalies
