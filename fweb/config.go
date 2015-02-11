package fweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Config instance. It's a two-level map.
type Config map[string]map[string]interface{}

// String returns config string by given key string.
func (cfg *Config) String(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return ""
	}
	str, ok := (*cfg)[keys[0]][keys[1]]
	if !ok {
		return ""
	}
	return fmt.Sprint(str)
}

// StringOr returns config string by given key string.
// It returns def string if empty string.
func (cfg *Config) StringOr(key string, def string) string {
	value := cfg.String(key)
	if value == "" {
		cfg.Set(key, def)
		return def
	}
	return value
}

// Int returns config int by given string.
func (cfg *Config) Int(key string) int {
	str := cfg.String(key)
	i, _ := strconv.Atoi(str)
	return i
}

// IntOr returns config int by given string.
// It returns def int if 0.
func (cfg *Config) IntOr(key string, def int) int {
	i := cfg.Int(key)
	if i == 0 {
		cfg.Set(key, def)
		return def
	}
	return i
}

// Float returns config float64 by given string.
func (cfg *Config) Float(key string) float64 {
	str := cfg.String(key)
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

// FloatOr returns config float64 by given string.
// It returns def float64 if float 0.
func (cfg *Config) FloatOr(key string, def float64) float64 {
	f := cfg.Float(key)
	if f == 0.0 {
		cfg.Set(key, def)
		return def
	}
	return f
}

// Bool returns config bool.
func (cfg *Config) Bool(key string) bool {
	str := cfg.String(key)
	b, _ := strconv.ParseBool(str)
	return b
}

// Set value with given key string.
// The key need as "section.name".
func (cfg *Config) Set(key string, value interface{}) {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return
	}
	if (*cfg) == nil {
		(*cfg) = make(map[string]map[string]interface{})
	}
	if _, ok := (*cfg)[keys[0]]; !ok {
		(*cfg)[keys[0]] = make(map[string]interface{})
	}
	(*cfg)[keys[0]][keys[1]] = value
}

// NewConfig creates new Config instance with json file.
func NewConfig(fileAbsPath string) (*Config, error) {
	cfg := new(Config)
	bytes, e := ioutil.ReadFile(fileAbsPath)
	if e != nil {
		return cfg, e
	}
	e = json.Unmarshal(bytes, cfg)
	return cfg, e
}
