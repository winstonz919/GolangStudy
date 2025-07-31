/*
 * @Author: WinstonRD
 * @Date: 2025-07-18 15:24:43
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-20 00:02:59
 * @FilePath: /34-go-redis/internal/config/config.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type CtxKey string

type Config struct {
	RedisConf *RedisConf `yaml:"redis"`
}

type RedisConf struct {
	Addr  string `yaml:"addr"`
	Passd string `yaml:"passd"`
	Db    int    `yaml:"db"`
}

func NewConfig(confPath string) *Config {
	file, err := os.ReadFile(confPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var data Config

	if err = yaml.Unmarshal(file, &data); err != nil {
		fmt.Println(err)
		return nil
	}
	return &data
}
