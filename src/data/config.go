package data

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	BaseProbability float64 `json:"BaseProbability"` //基础成功概率
	BaseFee         float64 `json:"BaseFee"`         //基础费用
	MaxCount        int     `json:"MaxCount"`        //最大强化次数，保底

	ExtraProbability float64 `json:"ExtraProbability"` //额外成功概率
	ExtraFee         float64 `json:"ExtraFee"`         //额外精炼材料费用（加概率）
	ExtraMaxCount    int     `json:"ExtraMaxCount"`    //最大强化次数，保底（保护强化时）

	FailAddProbability float64 `json:"FailAddProbability"` //失败增加概率
}

var defaultConfig = &Config{
	BaseProbability:    0.8,
	BaseFee:            100,
	FailAddProbability: 0.1,
	ExtraProbability:   0.2,
	ExtraFee:           100,
	MaxCount:           3,
	ExtraMaxCount:      1,
}

var CachePath = "./cache"
var CfgFileName = "config.json"

func ReadConfig() (*Config, error) {
	f, er := os.Open(CachePath + "/" + CfgFileName)
	if er != nil {
		return defaultConfig, nil //默认文件
	}

	defer f.Close()

	data, er := io.ReadAll(f)
	if er != nil {
		return nil, er
	}

	cfg := &Config{}
	er = json.Unmarshal(data, cfg)
	if er != nil {
		return nil, er
	}

	return cfg, nil
}

func SaveConfig(cfg *Config) error {
	if cfg == nil {
		return nil
	}

	er := os.MkdirAll(CachePath, os.ModePerm)
	if er != nil {
		return er
	}

	f, er := os.OpenFile(CachePath+"/"+CfgFileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if er != nil {
		return er
	}

	defer f.Close()

	data, er := json.MarshalIndent(cfg, "", "    ")
	if er != nil {
		return er
	}

	f.Write(data)
	return nil
}
