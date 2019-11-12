package ai

import "agar-life/math/crd"

type memory struct {
	valid     bool
	priority  uint8
	validTime uint
	reason    string
	crd       crd.Crd
}

func (m *memory) set(pr uint8, vt uint, reason string, crd crd.Crd) {
	m.valid = true
	m.priority = pr
	m.validTime = vt
	m.reason = reason
	m.crd = crd
}

func (m *memory) check(pr uint8, cycle uint) (bool, crd.Crd) {
	if m.valid && m.validTime < cycle && m.priority >= pr {
		return true, m.crd
	}
	//m.reset()
	return false, m.crd
}

func (m *memory) checkByReason(pr uint8, cycle uint, reason string) (bool, crd.Crd) {
	if m.valid && m.validTime > cycle && m.priority >= pr && m.reason == reason {
		return true, m.crd
	}
	m.reset()
	return false, m.crd
}

func (m *memory) reset() {
	m.valid = false
}
