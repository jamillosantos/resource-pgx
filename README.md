# rscpgx

This package implements a Resource from the [go-services](https://github.com/jamillosantos/go-services) library.

## Configuration

```go
type PlatformConfig struct {
	DSN string `config:"dsn,secret"`
}
```

| Name  | Type       | Description                                |
|-------|------------|--------------------------------------------|
| `DSN` | `string` | The DSN use to connect with the database.  |

Internally, `rscpgx` uses a the `pgxpool.New` to intiialize the database connection. So, the `DSN` is the same as the
`pgxpool.New` function, which implies that the `DSN` con configure all the parameters for a `pgxpool.New` 
([documentation](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool#hdr-Establishing_a_Connection)).
