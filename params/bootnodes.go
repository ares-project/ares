// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// ares Foundation Bootnodes
	//"enode://9a00a7d0c5157ca65491784d86864b4c11146e3fe50a19b1fca8d4c4c594292e0fd371f6d2182c52e4e61df63d5840ea8bb1e4a81ade2cc9f48d3ce99b28b31e@222.107.181.53:13990",
	"enode://285ee87e69efd4b55b1235f71449a60b2902854061ce9e782ac689530df84e29e3a2bad357d96eb7bf0692b9aad11eecb9d473bcd44e87aca6764fe6e5026145@211.233.87.170:13990",
	"enode://811ca81b7594d009badf55fdfd7c7b717d32e355cf3ecb26b2c32f706c8466ddee6ee8f96a76e26271f1acb78cae432d296d803d049dc37f333a2dbc487e62e1@211.233.87.170:13990", //first
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
