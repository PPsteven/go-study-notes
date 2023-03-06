package main

// USBADisk Type-A U盘
type USBADisk interface {
	PluggedIntoUSBAPort()
}

// USBCDisk Type-C U盘
type USBCDisk interface {
	PluggedIntoUSBCPort()
}

type USBADiskImpl struct {}

func (d *USBADiskImpl) PluggedIntoUSBAPort() {}

type USBCDiskImpl struct {}

func (d *USBCDiskImpl) PluggedIntoUSBCPort() {}

type Mac struct {}

func (m *Mac) LoadDisk(d USBCDisk){
	d.PluggedIntoUSBCPort()
}

type Win struct {}

func (w *Win) LoadDisk(d USBADisk) {
	d.PluggedIntoUSBAPort()
}

type A2CAdapter struct {
	USBADisk
}

func (a *A2CAdapter) PluggedIntoUSBCPort() {}

type C2AAdapter struct {
	USBCDisk
}

func (a *C2AAdapter) PluggedIntoUSBAPort() {}

func A2C(a USBADisk) USBCDisk {
	return &A2CAdapter{a}
}

func C2A(c USBCDisk) USBADisk {
	return &C2AAdapter{c}
}

func main() {
	mac := &Mac{}
	win := &Win{}
	usbA := &USBADiskImpl{}
	usbC := &USBCDiskImpl{}
	// win 支持USB口
	win.LoadDisk(usbA)
	// mac 支持TypeC口
	mac.LoadDisk(usbC)
	// 如果要让USBA口的U盘在mac上，需要借助适配器，即接头转换器
	mac.LoadDisk(A2C(usbA))
	// 如果要让TypeC口的U盘在win上，也借助适配器
	win.LoadDisk(C2A(usbC))
}

