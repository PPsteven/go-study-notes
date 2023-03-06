package main

import "fmt"

// Device -> Abstraction
type Device interface {
	setVolume(int)
	getVolume() int
	enable()
	disable()
	isEnable() bool
}

// XiaoMiTV and HuaweiTV Refined abstraction
type XiaoMiTV struct {
	on bool
	volume int
}

func (tv *XiaoMiTV) setVolume(vol int) {
	tv.volume = vol
	fmt.Println("set to ", vol)
}

func (tv *XiaoMiTV) getVolume() int {
	return tv.volume
}

func (tv *XiaoMiTV) enable() {
	if !tv.on {
		tv.on = true
	}
}

func (tv *XiaoMiTV) disable() {
	if tv.on {
		tv.on = false
	}
}

func (tv *XiaoMiTV) isEnable() bool {
	return tv.on
}

type HuaweiTV struct {
	on bool
	volume int
}

func (tv *HuaweiTV) setVolume(vol int) {
	tv.volume = vol
	fmt.Println("set to ", vol)
}

func (tv *HuaweiTV) getVolume() int {
	return tv.volume
}

func (tv *HuaweiTV) enable() {
	if !tv.on {
		tv.on = true
	}
}

func (tv *HuaweiTV) disable() {
	if tv.on {
		tv.on = false
	}
}

func (tv *HuaweiTV) isEnable() bool {
	return tv.on
}

type Remote struct {
	device Device
}

func NewRemote(device Device) *Remote{
	return &Remote{device}
}

func NewAdvancedRemote(device Device) *AdvancedRemote{
	return &AdvancedRemote{0, Remote{device}}
}

func (m *Remote) Switch() {
	if m.device.isEnable() {
		m.device.disable()
	} else {
		m.device.enable()
	}
}

func (m *Remote) VolumeUp() {
	m.device.setVolume(m.device.getVolume()+1)
}

func (m *Remote) VolumeDown() {
	m.device.setVolume(m.device.getVolume()-1)
}

type AdvancedRemote struct {
	originalVol int
	Remote
}

func (m *AdvancedRemote) Mute() {
	m.originalVol = m.device.getVolume()
	m.device.setVolume(0)
}

func (m *AdvancedRemote) UnMute() {
	m.device.setVolume(m.originalVol)
	m.originalVol = 0
}

func main() {
	xiaomi := &XiaoMiTV{}
	huawei := &HuaweiTV{}
	remote := NewRemote(xiaomi)
	remote.Switch()
	remote.VolumeUp()

	advancedRemote := NewAdvancedRemote(huawei)
	advancedRemote.VolumeUp()
	advancedRemote.VolumeUp() // vol = 2
	advancedRemote.Mute() // vol = 0
	advancedRemote.UnMute() // vol = 2
}