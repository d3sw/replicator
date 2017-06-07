package structs

// FailsafeMode is the configuration struct for administratively interacting
// with the global failsafe lock.
type FailsafeMode struct {
	// Config stores partial configuration required to interact with Consul.
	Config *Config

	// Disable instructs the failsafe CLI command to disable failsafe mode.
	Disable bool

	// Enable instructs the failsafe CLI command to enable failsafe mode.
	Enable bool

	// Force supresses confirmation prompts when enabling/disabling failsafe.
	Force bool

	// Verb represents the action to be displayed during confirmation prompts.
	Verb string
}
