# go-api-boilerplate

This boilerplate project is based on @corylanou's "tns-restful-json-api" project. Contains middleware mechanism with "Authentication Middleware" example and JWT encoding and decoding examples.

### Usage

- You can define new routes or modify existing ones in "routes.go" file. You need to sepecify name, method (GET, POST, PUT...), pattern, handler and middleware (if you will use otherwise nil) of a route.
- You can write your handler functions in "handler" folder, middlewares in "middleware" folder and models in "model" folder and use them after import these packages to relevant go files.
