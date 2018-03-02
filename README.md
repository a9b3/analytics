# Analytics

go server for collecting analytics

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
