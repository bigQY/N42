// Copyright 2023 The N42 Authors
// This file is part of the N42 library.
//
// The N42 library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The N42 library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the N42 library. If not, see <http://www.gnu.org/licenses/>.

package rawdb

import (
	"encoding/json"
	"fmt"
	"github.com/n42blockchain/N42/common/types"
	"github.com/n42blockchain/N42/log"
	"github.com/n42blockchain/N42/modules"

	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/n42blockchain/N42/params"
)

// ReadChainConfig retrieves the consensus settings based on the given genesis hash.
func ReadChainConfig(db kv.Getter, hash types.Hash) (*params.ChainConfig, error) {
	data, err := db.GetOne(modules.ChainConfig, modules.ConfigKey(hash))
	if err != nil {
		return nil, fmt.Errorf("fetch ChainConfig from db ,error: %v", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("ChainConfig are empty")
	}
	var config params.ChainConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("invalid chain config JSON err: %v", err)
	}
	return &config, nil
}

// WriteChainConfig writes the chain config settings to the database.
func WriteChainConfig(db kv.RwTx, hash types.Hash, cfg *params.ChainConfig) error {
	if cfg == nil {
		return fmt.Errorf("invalid cfg")
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Error("Failed to JSON encode chain config", "err", err)
		return err
	}
	if err := db.Put(modules.ChainConfig, modules.ConfigKey(hash), data); err != nil {
		log.Error("Failed to store chain config", "err", err)
		return err
	}
	return nil
}
