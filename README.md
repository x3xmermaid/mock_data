# Netmonk Mock App Data

This service show mock data api response XML.

## Main Features

- Show Mock Data

## Status
Under Development

## How to Start Development

### Clone Repository
```bash
$ cd $GOPATH/src/
$ mkdir ketitik
$ cd ketitik
$ mkdir netmonk
$ cd netmonk
$ git clone https://yourusername@bitbucket.org/ketitik/mock-app-data.git
```

## Running Services

### Netmonk BEC-GA API Service

Create netmonk_config.yaml in _bin/conf
Tamplate config file :
```bash
service_data:
  address: localhost:8080

```

Build Service
```bash
./_scripts/build.sh
```

Run Service Inventory
```bash
./_scripts/run.sh 
```

## License
Ketitik's Proprietary.