# Analytics

go server for collecting analytics

## API

### Application

Each user can have multiple applications. Each application will allow you to
send data to it.

|URL|Method|Body|Header|Description|Response|
|---|---|---|---|---|---|
|`/app`|GET||`authorization:jwt`|Get list of apps|`{ results: [app...] }`|
|`/app`|POST|`{name:string}`|`authorization:jwt`|Create an app|`{name,...}`|
|`/app`|PATCH|`{name:string,...}`|`authorization:jwt`|Patches an app|`{name,...}`|
|`/app/{id}/track`|POST|`{name:string, data:object}`|`key:string`|Send metric to track to the app|null|

## Dev

- Install [go-watcher](https://github.com/canthefason/go-watcher) for automatically restarting server upon file change during dev.
- Use [dep](https://github.com/golang/dep) for dependencies.
- Make a `.env` file, look at `config/config.go` for the required variables.

Run `make` to see the available tasks.

```

  deps       - Installs dependencies
  dev        - Runs development server     PORT ?= 9090
  test       - Runs tests
  db.start   - Starts the development dbs
  db.stop    - Stops the development dbs

```
