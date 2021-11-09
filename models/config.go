/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/11/9 17:24
 */

package models

type Config struct {
	App []struct {
		Name  string   `yaml:"name"`
		TLS   string   `yaml:"tls,omitempty"`
		Ports []string `yaml:"ports"`
		Hosts []string `yaml:"hosts"`
		Path  []struct {
			Path    string `yaml:"path"`
			Service string `yaml:"service"`
		} `yaml:"path"`
		Block []string `yaml:"block,omitempty"`
	} `yaml:"app"`
	Services []struct {
		Name   string `yaml:"name"`
		Type   string `yaml:"type"`
		Addr   string `yaml:"addr,omitempty"`
		Header []struct {
			Key   string `yaml:"key"`
			Value string `yaml:"value"`
		} `yaml:"header,omitempty"`
		Root string `yaml:"root,omitempty"`
	} `yaml:"services"`
	TLS []struct {
		Name string `yaml:"name"`
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	} `yaml:"tls"`
}
