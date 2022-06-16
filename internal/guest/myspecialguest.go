package guest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/vagrant-plugin-sdk/component"
	plugincore "github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

// ParrotOS is a Guest implementation.
type ParrotOS struct {
}

// DetectFunc implements component.Guest
func (h *ParrotOS) GuestDetectFunc() interface{} {
	return h.Detect
}

func (h *ParrotOS) Detect(t plugincore.Target, ui terminal.UI) (bool, error) {
	m, err := t.Specialize((*plugincore.Machine)(nil))
	if err != nil {
		return false, fmt.Errorf("oh no, could not specialize target")
	}
	machine := m.(plugincore.Machine)
	comm, err := machine.Communicate()
	if err != nil {
		return false, fmt.Errorf("oh no, could not get communicator")
	}

	cmd := `cat /etc/os-release | grep ID=parrot`
	found, err := comm.Test(machine, strings.Split(cmd, " "))
	ui.Output(fmt.Sprintf("Detected parrot: %t", found))
	return found, err
}

// ParentsFunc implements component.Guest
func (h *ParrotOS) ParentFunc() interface{} {
	return h.Parent
}

func (h *ParrotOS) Parent() string {
	return "debian"
}

// HasCapabilityFunc implements component.Guest
func (h *ParrotOS) HasCapabilityFunc() interface{} {
	return h.CheckCapability
}

func (h *ParrotOS) CheckCapability(n *component.NamedCapability) bool {
	return false
}

// CapabilityFunc implements component.Guest
func (h *ParrotOS) CapabilityFunc(name string) interface{} {
	return fmt.Errorf("requested capability %s not found", name)
}

var (
	_ component.Guest = (*ParrotOS)(nil)
)
