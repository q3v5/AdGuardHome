package dnsforward

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateDNSRequestDevice(t *testing.T) {
	testCases := []struct {
		name    string
		device  *DNSRequestDevice
		wantErr bool
	}{{
		name:    "nil_device",
		device:  nil,
		wantErr: false,
	}, {
		name: "disabled_no_user_agent",
		device: &DNSRequestDevice{
			Enabled:   false,
			UserAgent: "",
		},
		wantErr: false,
	}, {
		name: "disabled_with_user_agent",
		device: &DNSRequestDevice{
			Enabled:   false,
			UserAgent: "TestAgent/1.0",
		},
		wantErr: false,
	}, {
		name: "enabled_with_user_agent",
		device: &DNSRequestDevice{
			Enabled:   true,
			UserAgent: "TestAgent/1.0",
		},
		wantErr: false,
	}, {
		name: "enabled_without_user_agent",
		device: &DNSRequestDevice{
			Enabled:   true,
			UserAgent: "",
		},
		wantErr: true,
	}, {
		name: "enabled_user_agent_too_long",
		device: &DNSRequestDevice{
			Enabled:   true,
			UserAgent: strings.Repeat("a", maxUserAgentLen+1),
		},
		wantErr: true,
	}, {
		name: "enabled_user_agent_with_newline",
		device: &DNSRequestDevice{
			Enabled:   true,
			UserAgent: "Test\nAgent",
		},
		wantErr: true,
	}, {
		name: "enabled_user_agent_with_carriage_return",
		device: &DNSRequestDevice{
			Enabled:   true,
			UserAgent: "Test\rAgent",
		},
		wantErr: true,
	}, {
		name: "enabled_user_agent_with_tab",
		device: &DNSRequestDevice{
			Enabled:   true,
			UserAgent: "Test\tAgent",
		},
		wantErr: true,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateDNSRequestDevice(tc.device)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsValidUserAgent(t *testing.T) {
	testCases := []struct {
		name string
		ua   string
		want bool
	}{{
		name: "empty",
		ua:   "",
		want: true,
	}, {
		name: "valid_simple",
		ua:   "TestAgent/1.0",
		want: true,
	}, {
		name: "valid_with_spaces",
		ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
		want: true,
	}, {
		name: "invalid_newline",
		ua:   "Test\nAgent",
		want: false,
	}, {
		name: "invalid_carriage_return",
		ua:   "Test\rAgent",
		want: false,
	}, {
		name: "invalid_tab",
		ua:   "Test\tAgent",
		want: false,
	}, {
		name: "invalid_control_char",
		ua:   "Test\x01Agent",
		want: false,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, isValidUserAgent(tc.ua))
		})
	}
}
