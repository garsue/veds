# Veds: Viewer for Emulated Datastore

Veds is browser-based datastore emulator viewer.

# Install

```
go get github.com/garsue/veds
```

# Usage

Start our viewer server.

```
veds localhost:8080
```

`localhost:8080` is datastore emulator's host.

Then you can viewer with http://localhost:8090 in your web browser.

# Requirements

Go 1.8.0 or later.

# Features

Can not do anything!

# Next

* View entities
* Create entites

# Build

Use [dep](https://github.com/golang/dep) for vendoring.

```
dep ensure -update
```

Then you can build the project.
