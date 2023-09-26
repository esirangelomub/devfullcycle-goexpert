# My API

## Guidelines for API

Consulting [github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Struct of API

### cmd/

The cmd directory contains the main application for this project.

### internal/

The internal directory contains private application and library code.

### pkg/

The pkg directory contains library code that's ok to use by external applications.

### configs/

The configs directory contains scripts do configure the project.

### api/

The api directory contains the OpenAPI/Swagger specs, JSON schema files, protocol definition files.

### test/

The test directory contains the tests.
