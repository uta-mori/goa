// Package goa implements a Go framework for writing microservices that promotes
// best practice by providing a single source of truth from which server code,
// client code, and documentation is derived. The code generated by goa follows
// the clean architecture pattern where composable modules are generated for the
// transport, endpoint, and business logic layers. The goa package contains
// middleware, plugins, and other complementary functionality that can be
// leveraged in tandem with the generated code to implement complete
// microservices in an efficient manner. By using goa for developing
// microservices, implementers don’t have to worry with the documentation
// getting out of sync from the implementation as goa takes care of generating
// OpenAPI specifications for HTTP based services and gRPC protocol buffer files
// for gRPC based services (or both if the service supports both transports).
// Reviewers can also be assured that the implementation follows the
// documentation as the code is generated from the same source.
//
// Visit https://goa.design for more information.
package goa