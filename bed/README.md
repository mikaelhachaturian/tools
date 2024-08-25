# bed - base64 Encrypt/Decrypt

```sh
Encode/Decode base64 string.

Usage:
  bed [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decode      Decode base64 to string.
  encode      Encode string to base64.
  help        Help about any command

Flags:
  -h, --help   help for bed

Use "bed [command] --help" for more information about a command.
```

## Examples

**Encode**:
```sh
❯ bed encode hi
aGk=
```

**Decode**:
```sh
❯ bed decode aGk=
hi
```