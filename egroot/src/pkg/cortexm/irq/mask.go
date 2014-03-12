package irq

// Disable disables all exceptions other than NMI and faults. Internally it
// sets Cortex-M PRIMASK to 1.
func Disable()

// Enable reverts Disable. If you modified any data that can be used by enabled
// interrupt handlers you probably need to call sync.Memory() before.
func Enable()

// Disabled returns true if excepions are disabled (PRIMASK != 0).
func Disabled() bool

// DisableFaults disables all exceptions other than NMI. Internally it sets
// Cortex-M FAULTMASK to 1.
func DisableFaults()

// EnableFaults reverts DisableFaults. If you modified any data that can be used
// by enabled interrupt handlers you probably need to call sync.Memory() before.
func EnableFaults()

// FaultsDisabled returns true if excepions and faults are disabled
// (FAULTMASK != 0).
func FaultsDisabled() bool

// DisablePri disables the same or lover priority exceptions than p.
func DisablePri(p Prio)

// TODO
func DisablePriMax(p Prio)
func DisabledPri() Prio
func EnablePri()