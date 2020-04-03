package main

import (
    "plugin"
    "path/filepath"
    "fmt"
)

type middleware interface {
    Sanitize(string) string
}

type PluginManager struct{

}

// an industry solution should search plugins at launch, creating a sort of registry with names of plugins and 
// the exported fields of those plugins. Benefits:
// 1) Unused plugins do not need to be left in memory, therefore saving memory
// 2) Easy identification if plugin does not exist
// 
// For debugging, would be nice to be able to load a plugin at run time

func (p PluginManager) LoadPlugins(name string) middleware {
    plugins, err := filepath.Glob("plugins/*.so")
    if err != nil {
        panic(err)
    }

    for _, fn := range plugins {
        file, err := plugin.Open(fn)
        if err != nil {
            panic(err)
        }

        symbol, err := file.Lookup(name)

        if err != nil {
            continue
        }

        midPlugin, ok := symbol.(middleware)
        if !ok {
            fmt.Println("Unexpected type from module symbol. Does your plugin implement the 'middleware' interface?")
        }
        return midPlugin
    }
    return nil
}

