# compose-cli

A command line tool to compose files. 

Composes output from an input template file interspersed with template variables and a values file with values for those variables. The output is sent to stdout, but can be overridden with a flag.

## Usage

The `compose` sub-command can compose arbitrary files. In the example below, it composes from `in.txt` with values interpolated from `overrides.yaml`. The output can be sent to a file with the `--out` flag.

```sh
$ cat in.txt
Hello {{ .Values.who }}
$ cat overrides.yaml
who: Mario

$ composer compose --in in.txt --values overrides.yaml
Hello Mario

$ composer compose --in in.txt --values overrides.yaml --out out.txt
```

The `docker` sub-command composes a `docker-compose.yaml` file from
`template.yaml` and `values.yaml` files in the same directory. 
Any of the values in `values.yaml` can be overridden with the `--set`
flag. 

```sh
$ compose docker
$ compose docker --set container_name=whoami
```
