# Overview

Projects consists of gRPC service called `pb`
and it's client API called `client-api`.

## Client-API
Code for Client-API is located [here](client-api).
This service parses json file and creates entities which
then will be sent to `pds` over gRPC protocol.

### Parsing
In current implementation we create `io.Reader` from 
opening the file and calling `bufio.NewReader` onto the 
file handler.
Then we use [Go's Decode Stream API](https://golang.org/pkg/encoding/json/#example_Decoder_Decode_stream) 
to unmarshal our entities by chunks.
We assume that `io.Reader` points to a valid JSON file
and read first `json.Delim` token. Then we 
read token once more to get entity's `CODE` from it.
We wouldn't have to do that if JSON file was in format like that:
```json
[
  {"code": "XXXX", "alias": "XX"},
  {"code": "XXXX", "alias": "XX"}
]

```
Then we execute callback if unmarshalling into an entity
was successfull. In our case, we send it to PDB.

### Client
[Client package](client-api/api/client.go) is responsible for
communication between client-api and pdb.
This client just implements methods of PDServiceClient.

### Jaeger
It's also possible to bootstrap [jaeger](https://www.jaegertracing.io/) for full overview
of client-api performance! To do so, just define these env
variable which will point to JAEGER host.
```.env
JAEGER_AGENT_PORT=6831
JAEGER_SAMPLER_TYPE=const
JAEGER_SAMPLER_PARAM=1
JAEGER_AGENT_HOST=localhost
```

### Server
Client-API also has HTTP interface to communicate
with PDS:

#### Routes
URL | Description
------------ | -------------
**/ping** | Health check route. Could be used as a mark that application has successfully started.
**/ports** | List all ports which have been successfully parsed and inserted into datalayer of PDS.
**/ports/{CODE}** | Find and show one Port by given Code (Code should be in XXXXX format).
**/ports/{CODE}/delete** | Deletes Port with given Code from datalayer in PDS.

# PDS
Code for PDS is located [here](pds).
This service receives data from [Client-API](client-api) and uses
datalayer to manage it. 
As datalayer we use [InMemoryStorage](pds/datalayer/impl.go).

## Development
To regenerate protobuf code please run `make generate` (assuming you have [protoc-gen-go](https://github.com/golang/protobuf) installed).
 