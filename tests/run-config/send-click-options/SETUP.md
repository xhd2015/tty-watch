# Scenario

**Feature**: `ParseArgs` maps send click flags to Config.Send without registry

```
cli.ParseArgs(["send", "sid", "--click", "--row", "10", "--col", "67"])
  -> Config{Command:"send", Send:{Session:"sid", Mode:Click, Row:10, Col:67, …}}
# pure flag validation; no Run / no inject
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Mode = "parse-args"
	req.Args = []string{"send", "sid", "--click", "--row", "10", "--col", "67"}
	return nil
}
```
