# Chaos Core Protobuf

Shared protobuf contracts for Chaos projects. The repository publishes two
independent Buf modules and their Go runtime packages.

| Module | Purpose | Go package |
| --- | --- | --- |
| `buf.build/chaos/core` | Common value types: time, duration, error, resource, file, URL, value, and version. | `github.com/chaos-io/core/go/chaos/core` |
| `buf.build/chaos/fino2` | Custom protobuf options used by fino2-generated projects. | `github.com/chaos-io/core/go/fino2` |

## Use in a protobuf module

Add only the modules your schema imports:

```yaml
deps:
  - buf.build/chaos/core
  - buf.build/chaos/fino2
```

```proto
import "chaos/core/time.proto";
import "fino2/options.proto";

message Book {
  option (fino2.model) = {generate: true services: "BookService"};

  chaos.core.Timestamp create_time = 1;
}
```

Generated Go code imports the corresponding shared package. Do not vendor or
regenerate these proto files in individual services: a process must link each
protobuf descriptor only once.

## Maintain

Run Buf commands from the repository root:

```sh
buf lint
buf build
buf push
```

`buf push` publishes `chaos/core` and `chaos/fino2` independently. Changing a
module's protobuf package name, field numbers, extension names, or extension
field numbers is a public compatibility change and must be treated as breaking.
