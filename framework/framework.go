package framework

type Framework struct {
	Gtk *GTK
	IO *IO
	Process *Process
	Resource *Resource
	Slice *Slice
}

func NewFramework() *Framework {
	framework := new(Framework)
	framework.Gtk = &GTK{}
	framework.IO = &IO{}
	framework.Process = &Process{}
	framework.Resource = &Resource{}
	framework.Slice = &Slice{}
	return framework
}