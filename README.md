# Analytics

go server for collecting analytics

## Dev

- Install [go-watcher](https://github.com/canthefason/go-watcher) for automatically restarting server upon file change during dev.
- Use [dep](https://github.com/golang/dep) for dependencies.

Run `make` to see the available tasks.

```
  deps       - Installs dependencies
  dev        - Runs development server   PORT ?= 5001
  lint       - Runs linter
  test       - Runs tests
```
