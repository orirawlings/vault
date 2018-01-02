package command

import (
	"testing"

	"github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/meta"
	"github.com/hashicorp/vault/vault"
	"github.com/mitchellh/cli"
)

func Test_Seal(t *testing.T) {
	core, _, token := vault.TestCoreUnsealed(t)
	ln, addr := http.TestServer(t.Fatalf, core)
	defer ln.Close()

	ui := new(cli.MockUi)
	c := &SealCommand{
		Meta: meta.Meta{
			ClientToken: token,
			Ui:          ui,
		},
	}

	args := []string{"-address", addr}
	if code := c.Run(args); code != 0 {
		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
	}

	sealed, err := core.Sealed()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if !sealed {
		t.Fatal("should be sealed")
	}
}
