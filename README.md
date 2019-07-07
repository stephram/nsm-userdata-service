
# National Sports Museum - User Data Service


Assumes a working go1.11.x or later environment.

Build and run
```
make install
make generate
make build
make run
```

Modifiy the API in `spec/userdata-swagger.yaml` then repeat the `generate, build, run` steps above.

### The API
The API is embedded as an open-api 2.0 specification (Swagger). It can be viewed a couple of ways.

- ```http://localhost:8080/swagger.json```
- ```http://localhost:8080/swagger-ui```

