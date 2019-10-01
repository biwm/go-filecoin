package fixtures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	cid "github.com/ipfs/go-cid"

	th "github.com/filecoin-project/go-filecoin/testhelpers"
	"github.com/filecoin-project/go-filecoin/types"
)

// The file used to build these addresses can be found in:
// $GOPATH/src/github.com/filecoin-project/go-filecoin/fixtures/setup.json
//
// If said file is modified these addresses will need to change as well
// rebuild using
// TODO: move to build script
// https://github.com/filecoin-project/go-filecoin/issues/921
// cat ./fixtures/setup.json | ./gengen/gengen --json --keypath fixtures > fixtures/genesis.car 2> fixtures/gen.json

// TestAddresses is a list of pregenerated addresses.
var TestAddresses []string

// testKeys is a list of filenames, which contain the private keys of the pregenerated addresses.
var testKeys []string

// TestMiners is a list of pregenerated miner acccounts. They are owned by the matching TestAddress.
var TestMiners []string

type detailsStruct struct {
	Keys   []*types.KeyInfo
	Miners []struct {
		Owner               int
		Address             string
		NumCommittedSectors uint64
	}
	GenesisCid cid.Cid `refmt:",omitempty"`
}

func init() {
	gopath, err := th.GetGoPath()
	if err != nil {
		panic(err)
	}

	detailspath := filepath.Join(gopath, "/src/github.com/filecoin-project/go-filecoin/fixtures/test/gen.json")
	detailsFile, err := os.Open(detailspath)
	if err != nil {
		// fmt.Printf("Fixture data not found. Skipping fixture initialization: %s\n", err)
		return
	}
	defer func() {
		if err := detailsFile.Close(); err != nil {
			panic(err)
		}
	}()
	detailsFileBytes, err := ioutil.ReadAll(detailsFile)
	if err != nil {
		panic(err)
	}
	var details detailsStruct
	if err := json.Unmarshal(detailsFileBytes, &details); err != nil {
		panic(err)
	}

	var keys []int
	for key := range details.Keys {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	miners := details.Miners

	for _, key := range keys {
		info := details.Keys[key]
		addr, err := info.Address()
		if err != nil {
			panic(err)
		}
		TestAddresses = append(TestAddresses, addr.String())
		testKeys = append(testKeys, fmt.Sprintf("%d.key", key))
	}

	for _, miner := range miners {
		TestMiners = append(TestMiners, miner.Address)
	}
}

// KeyFilePaths returns the paths to the wallets of the testaddresses
func KeyFilePaths() []string {
	gopath, err := th.GetGoPath()
	if err != nil {
		panic(err)
	}
	folder := "/src/github.com/filecoin-project/go-filecoin/fixtures/test/"

	res := make([]string, len(testKeys))
	for i, k := range testKeys {
		res[i] = filepath.Join(gopath, folder, k)
	}

	return res
}

// staging devnet addrs
const (
	stagingFilecoinBootstrap0 string = "/dns4/staging.kittyhawk.wtf/tcp/9000/ipfs/Qmd6xrWYHsxivfakYRy6MszTpuAiEoFbgE1LWw4EvwBpp4"
	stagingFilecoinBootstrap1 string = "/dns4/staging.kittyhawk.wtf/tcp/9001/ipfs/QmXq6XEYeEmUzBFuuKbVEGgxEpVD4xbSkG2Rhek6zkFMp4"
	stagingFilecoinBootstrap2 string = "/dns4/staging.kittyhawk.wtf/tcp/9002/ipfs/QmXhxqTKzBKHA5FcMuiKZv8YaMPwpbKGXHRVZcFB2DX9XY"
	stagingFilecoinBootstrap3 string = "/dns4/staging.kittyhawk.wtf/tcp/9003/ipfs/QmZGDLdQLUTi7uYTNavKwCd7SBc5KMfxzWxAyvqRQvwuiV"
	stagingFilecoinBootstrap4 string = "/dns4/staging.kittyhawk.wtf/tcp/9004/ipfs/QmZRnwmCjyNHgeNDiyT8mXRtGhP6uSzgHtrozc42crmVbg"
)

// nightly devnet addrs
const (
	nightlyFilecoinBootstrap0 string = "/dns4/nightly.kittyhawk.wtf/tcp/9000/ipfs/Qmd6xrWYHsxivfakYRy6MszTpuAiEoFbgE1LWw4EvwBpp4"
	nightlyFilecoinBootstrap1 string = "/dns4/nightly.kittyhawk.wtf/tcp/9001/ipfs/QmXq6XEYeEmUzBFuuKbVEGgxEpVD4xbSkG2Rhek6zkFMp4"
	nightlyFilecoinBootstrap2 string = "/dns4/nightly.kittyhawk.wtf/tcp/9002/ipfs/QmXhxqTKzBKHA5FcMuiKZv8YaMPwpbKGXHRVZcFB2DX9XY"
	nightlyFilecoinBootstrap3 string = "/dns4/nightly.kittyhawk.wtf/tcp/9003/ipfs/QmZGDLdQLUTi7uYTNavKwCd7SBc5KMfxzWxAyvqRQvwuiV"
	nightlyFilecoinBootstrap4 string = "/dns4/nightly.kittyhawk.wtf/tcp/9004/ipfs/QmZRnwmCjyNHgeNDiyT8mXRtGhP6uSzgHtrozc42crmVbg"
)

// user devnet addrs
const (
	userFilecoinBootstrap0 string = "/dns4/user.kittyhawk.wtf/tcp/9000/ipfs/Qmd6xrWYHsxivfakYRy6MszTpuAiEoFbgE1LWw4EvwBpp4"
	userFilecoinBootstrap1 string = "/dns4/user.kittyhawk.wtf/tcp/9001/ipfs/QmXq6XEYeEmUzBFuuKbVEGgxEpVD4xbSkG2Rhek6zkFMp4"
	userFilecoinBootstrap2 string = "/dns4/user.kittyhawk.wtf/tcp/9002/ipfs/QmXhxqTKzBKHA5FcMuiKZv8YaMPwpbKGXHRVZcFB2DX9XY"
	userFilecoinBootstrap3 string = "/dns4/user.kittyhawk.wtf/tcp/9003/ipfs/QmZGDLdQLUTi7uYTNavKwCd7SBc5KMfxzWxAyvqRQvwuiV"
	userFilecoinBootstrap4 string = "/dns4/user.kittyhawk.wtf/tcp/9004/ipfs/QmZRnwmCjyNHgeNDiyT8mXRtGhP6uSzgHtrozc42crmVbg"
)

// user devnet addrs
const (
	alphaStagingFilecoinBootstrap0 string = "/dns4/bootstrappers-0.alpha-staging.kittyhawk.wtf/tcp/30600/ipfs/QmcnXbqvMEwqG4RmWVgRggbLM4aNtu1BZ123qJDRpK6AqA"
	alphaStagingFilecoinBootstrap1 string = "/dns4/bootstrappers-1.alpha-staging.kittyhawk.wtf/tcp/30601/ipfs/QmXeMDTYcUm5qgASyRYU52QNc8rvjMY5eHgTGVmF6C2rn5"
	alphaStagingFilecoinBootstrap2 string = "/dns4/bootstrappers-2.alpha-staging.kittyhawk.wtf/tcp/30602/ipfs/QmbKWFbYgV7QRUNMQsQoUDJDVgaV7nm67LmrxkUTTKJDec"
)

// person devnet addrs
const (
	personFilecoinBootstrap0 string = "/dns4/bootstrappers-0.person.kittyhawk.wtf/tcp/30800/ipfs/QmVsQTkx1aTTEJPCkfrLPiqyLjGWDYLxhk9uTx7Y9vpsAN"
	personFilecoinBootstrap1 string = "/dns4/bootstrappers-1.person.kittyhawk.wtf/tcp/30801/ipfs/QmXCuoJNpRSwopXLTytLkqNkn121gfYakYgUEyzSKi4aSA"
	personFilecoinBootstrap2 string = "/dns4/bootstrappers-2.person.kittyhawk.wtf/tcp/30802/ipfs/QmdJ6hqcsLHwY6yRF84h3KqtMedRcagDn9AJbjWHFQ6DXJ"
)

// DevnetStagingBootstrapAddrs are the dns multiaddrs for the nodes of the filecoin
// staging devnet.
var DevnetStagingBootstrapAddrs = []string{
	stagingFilecoinBootstrap0,
	stagingFilecoinBootstrap1,
	stagingFilecoinBootstrap2,
	stagingFilecoinBootstrap3,
	stagingFilecoinBootstrap4,
}

// DevnetNightlyBootstrapAddrs are the dns multiaddrs for the nodes of the filecoin
// nightly devnet
var DevnetNightlyBootstrapAddrs = []string{
	nightlyFilecoinBootstrap0,
	nightlyFilecoinBootstrap1,
	nightlyFilecoinBootstrap2,
	nightlyFilecoinBootstrap3,
	nightlyFilecoinBootstrap4,
}

// DevnetUserBootstrapAddrs are the dns multiaddrs for the nodes of the filecoin
// user devnet
var DevnetUserBootstrapAddrs = []string{
	userFilecoinBootstrap0,
	userFilecoinBootstrap1,
	userFilecoinBootstrap2,
	userFilecoinBootstrap3,
	userFilecoinBootstrap4,
}

// DevnetAlphaStagingBootstrapAddrs are the dns multiaddrs for the nodes of the filecoin
// alpha staging devnet
var DevnetAlphaStagingBootstrapAddrs = []string{
	alphaStagingFilecoinBootstrap0,
	alphaStagingFilecoinBootstrap1,
	alphaStagingFilecoinBootstrap2,
}

// DevnetAlphaStagingBootstrapAddrs are the dns multiaddrs for the nodes of the filecoin

// alpha staging devnet
var DevnetPersonBootstrapAddrs = []string{
	personFilecoinBootstrap0,
	personFilecoinBootstrap1,
	personFilecoinBootstrap2,
}
