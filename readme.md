# Servicar API

This is a [PocketBase](https://pocketbase.io/) API for [Servicar](https://servicar.com.br)

## How to build

for linux:
```bash
env GOOS=linux GOARCH=amd64 go build
```

for windows:
```powershell
env GOOS=windows GOARCH=amd64 go build
```

## How to use

```bash
./servicar(.exe) serve
```

navigate to [http://localhost:8090/_/](http://localhost:8090/_/) and enter the admin credentials
from there you can find instructions to the API, just click on the collection
and see the API preview.

### The API

The API is served at [http://localhost:8090/api/](http://localhost:8090/api/). But it can also be used with and SDK. 
See [docs](https://pocketbase.io/docs/client-side-sdks/).

### Static Content

For public files inside `pb_public` directory, navigate to `http://localhost:8090/<path>`. Ex:
`http://localhost:8090/carros/focushatch_<size>.webp`
`http://localhost:8090/montadoras/byd_<size>.webp`

images are served in three sizes:

- carros: `90`, `420`, `1280`
- montadoras: `70`, `150`, `320`


