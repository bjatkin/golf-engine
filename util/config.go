package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type golfConfig struct {
	name             string
	spriteFile       string
	mapFile          string
	flagFile         string
	outputSpriteFile string
	outputMapFile    string
	outputFlagFile   string
}

func (g *golfConfig) String() string {
	return "name=" + g.name + "\n" +
		"spriteFile=" + g.spriteFile + "\n" +
		"mapFile=" + g.mapFile + "\n" +
		"flagFile=" + g.flagFile + "\n" +
		"outputSpriteFile=" + g.outputSpriteFile + "\n" +
		"outputMapFile=" + g.outputMapFile + "\n" +
		"outputFlagFile=" + g.outputFlagFile
}

func (g *golfConfig) getProp(prop string) (string, error) {
	switch prop {
	case "name":
		return g.name, nil
	case "spriteFile":
		return g.spriteFile, nil
	case "mapFile":
		return g.mapFile, nil
	case "flagFile":
		return g.flagFile, nil
	case "outputSpriteFile":
		return g.outputSpriteFile, nil
	case "outputMapFile":
		return g.outputMapFile, nil
	case "outputFlagFile":
		return g.outputFlagFile, nil
	}
	return "", fmt.Errorf("No property named %s", prop)
}

func (g *golfConfig) setProp(prop, value string) error {
	switch prop {
	case "name":
		g.name = value
		return nil
	case "spriteFile":
		g.spriteFile = value
		return nil
	case "mapFile":
		g.mapFile = value
		return nil
	case "flagFile":
		g.flagFile = value
		return nil
	case "outputSpriteFile":
		g.outputSpriteFile = value
		return nil
	case "outputMapFile":
		g.outputMapFile = value
		return nil
	case "outputFlagFile":
		g.outputFlagFile = value
		return nil
	}
	return fmt.Errorf("No property named %s", prop)
}

func toGolfConfig(data string) golfConfig {
	pairs := strings.Split(data, "\n")
	ret := golfConfig{}
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, "=")
		k, v := kv[0], kv[1]

		switch k {
		case "name":
			ret.name = v
		case "spriteFile":
			ret.spriteFile = v
		case "mapFile":
			ret.mapFile = v
		case "flagFile":
			ret.flagFile = v
		case "outputSpriteFile":
			ret.outputSpriteFile = v
		case "outputMapFile":
			ret.outputMapFile = v
		case "outputFlagFile":
			ret.outputFlagFile = v
		}
	}
	return ret
}

func getGolfProp(args []string) error {
	prop := args[0]
	conf, err := ioutil.ReadFile("golf_config")
	if err != nil {
		return err
	}
	confData := toGolfConfig(string(conf))
	val, err := confData.getProp(prop)
	if err != nil {
		return err
	}
	fmt.Printf("   %v\n", val)
	return nil
}

func setGolfProp(args []string) error {
	prop := args[0]
	val := args[1]
	conf, err := ioutil.ReadFile("golf_config")
	if err != nil {
		return err
	}
	confData := toGolfConfig(string(conf))
	err = confData.setProp(prop, val)
	if err != nil {
		return err
	}

	return addFile("golf_config", []byte(confData.String()), true)
}
