package database

import "runtime"

type RawConfig struct {
	Path      string
	Cache     int
	Handles   int
	Ancient   string
	Namespace string
	ReadOnly  bool
}

func DefaultRawConfig() *RawConfig {
	if runtime.GOOS == "darwin" { // MacOS
		return &RawConfig{
			//Path: "../ethereumdata_1400w_small/copychain",
			Path: "/Users/zsj/Desktop/BCTS小组任务/智能合约应用逻辑课题/ethereumdata/copychain",
			//Path:      "/Volumes/ETH_DATA/ethereum/geth/chaindata",
			Cache:   2048,
			Handles: 5120,
			//Ancient: "../ethereumdata_1400w_small/copychain/ancient",
			Ancient: "/Users/zsj/Desktop/BCTS小组任务/智能合约应用逻辑课题/ethereumdata/copychain/ancient",
			//Ancient:   "/Volumes/ETH_DATA/ethereum/geth/chaindata/ancient",
			Namespace: "eth/db/chaindata/",
			ReadOnly:  false,
		}
	} else {
		return &RawConfig{
			Path: "/home/ubuntu/ethereumdata/copychain",
			//Path: "/home/fuzh/ethereumdata/copychain",
			//Path:    "../ethereumdata_1465w_small/copychain",
			Cache:   2048,
			Handles: 5120,
			Ancient: "/home/ubuntu/ethereumdata/copychain/ancient",
			//Ancient: "/home/fuzh/ethereumdata/copychain/ancient",
			//Ancient:   "../ethereumdata_1465w_small/copychain/ancient",
			Namespace: "eth/db/chaindata/",
			ReadOnly:  false,
		}
	}
}

type StateDBConfig struct {
	Cache     int
	Journal   string
	Preimages bool
}

func defaultStateDBConfig() *StateDBConfig {
	if runtime.GOOS == "darwin" { // MacOS
		return &StateDBConfig{
			Cache:     614,
			Journal:   "/Volumes/ETH_DATA/ethereum/geth/triecache",
			Preimages: false,
		}
	} else {
		return &StateDBConfig{
			Cache:     614,
			Journal:   "/experiment/ethereum/geth/triecache",
			Preimages: false,
		}
	}
}
