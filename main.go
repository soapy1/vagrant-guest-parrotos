package main

import (
	"os"

	sdk "github.com/hashicorp/vagrant-plugin-sdk"
	"github.com/soapy1/vagrant-guest-parrotos/internal/guest"
)

// Options are the SDK options to use for instantiation.
var ComponentOptions = []sdk.Option{
	sdk.WithComponents(
		&guest.ParrotOS{},
	),
	sdk.WithName("parrotos"),
}

func main() {
	sdk.Main(ComponentOptions...)
	os.Exit(0)
}
