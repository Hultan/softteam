package framework

type Framework struct {
	Gtk *Gtk
	IO *IO
	Process *Process
	Resource *Resource
	Slice *Slice
	Error *Error
}

func NewFramework() *Framework {
	framework := new(Framework)
	
	framework.Gtk = &Gtk{}
	framework.IO = &IO{}
	framework.Process = &Process{}
	framework.Resource = &Resource{}
	framework.Slice = &Slice{}
	framework.Error = &Error{}

	return framework
}