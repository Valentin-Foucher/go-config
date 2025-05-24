
# Go-Config

Simple library designed to load any configuration file and access it by describing the path towards the desired data.
Supported formats are JSON, YAML and TOML.

## Usage

Examples of code:

```go
func main() {
    cfg := config.Load("/some/config/file.json", config.JSON)
    cfg.GetString("environment")
    cfg.GetBool("google.analytics.enabled")
    cfg.GetString("scraper.algo.0.name")
}
```
