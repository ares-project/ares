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
	//"enode://7f2d420d5278b900b3a845dbed836e791373d2dcf39cc2828eaf110e545ec28306a60544937ada095f710d465c6a9f52f4e3abf4e3193b76f28ed65126d8793a@192.168.0.204:13990",    //seed1
	"enode://14216f33156aab3b2b2ece5f0925f25df33978ef10b04acbb3c20b05759e743c38df329f845295ec71817dfc2835ebd5b385089c95e8b820c4c4d930467b53ac@222.107.181.53:13990",
	"enode://bbf800555871d465ad58ec327d9eb200cb60eb38c47726e736fb58fd38185cdb555ffe8e63d9966b4918ea8af1afd7c5b42087d2dbaac16e7c74d0a04ebc10f9@211.251.238.155:13990",
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
