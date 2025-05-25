
# Go-Config

Simple library designed to load any configuration file and access it by describing the path towards the desired data.
Supported formats are JSON, YAML and TOML.

## Usage

Examples of code:

```go
func main() {
    cfg := config.Load("/some/config/file.json", config.JSON)
    cfg.GetStringOrDefault("environment", "dev")
    cfg.MustGetBool("google.analytics.enabled")
    cfg.MustGetString("scraper.algo.0.name")
}
```

## API Reference

These are the available functions to retrieve a key. They will return an error if type is incorrect or if the key is missing.
- `(Config).MustGetString` (prefix value with ---ENV to load the env variable)
- `(Config).MustGetInt`
- `(Config).MustGetFloat`
- `(Config).MustGetBool`
- `(Config).MustGetMap`
- `(Config).MustGetSlice`
- `MustGetType` (generic type)

These are the available functions to retrieve a key or fallback to the default value if an error occurred.
- `(Config).GetStringOrDefault`
- `(Config).GetIntOrDefault`
- `(Config).GetFloatOrDefault`
- `(Config).GetBoolOrDefault`
- `(Config).GetMapOrDefault`
- `(Config).GetSliceOrDefault`

`(Config).ListKeys` allows to retrieve the keys of a map anywhere in the configuration.