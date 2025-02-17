package hscontrol

import (
	"net/netip"
	"os"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

var _ = check.Suite(&Suite{})

type Suite struct{}

var (
	tmpDir string
	app    *Headscale
)

func (s *Suite) SetUpTest(c *check.C) {
	s.ResetDB(c)
}

func (s *Suite) TearDownTest(c *check.C) {
	os.RemoveAll(tmpDir)
}

func (s *Suite) ResetDB(c *check.C) {
	if len(tmpDir) != 0 {
		os.RemoveAll(tmpDir)
	}
	var err error
	tmpDir, err = os.MkdirTemp("", "autoygg-client-test2")
	if err != nil {
		c.Fatal(err)
	}
	cfg := Config{
		PrivateKeyPath:      tmpDir + "/private.key",
		NoisePrivateKeyPath: tmpDir + "/noise_private.key",
		DBtype:              "sqlite3",
		DBpath:              tmpDir + "/headscale_test.db",
		IPPrefixes: []netip.Prefix{
			netip.MustParsePrefix("10.27.0.0/23"),
		},
		OIDC: OIDCConfig{
			StripEmaildomain: false,
		},
	}

	app, err = NewHeadscale(&cfg)
	if err != nil {
		c.Fatal(err)
	}
}
