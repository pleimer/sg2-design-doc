package main

import (
    "strings"
    "os"
    "fmt"
)

func Handle(raw string, p middleware) string {
   return p.Sanitize(raw)
}

func main() {
    pluginName := os.Args[1]
    input := strings.Join(os.Args[2:], " ")
    pm := PluginManager{}
    plugin := pm.LoadPlugins(pluginName)
    if plugin != nil {
        println(Handle(input, plugin))
        return
    }

    fmt.Printf("Plugin %s does not exist\n", pluginName)
}
