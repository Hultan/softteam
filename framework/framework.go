package framework

const applicationVersion = "1.2.9"

type Framework struct {
	Gtk      *Gtk
	IO       *IO
	Process  *Process
	Resource *Resource
	Slice    *Slice
	Error    *Error
	Crypto   *Crypto
	Mail     *MailSender
}

func NewFramework() *Framework {
	framework := new(Framework)

	framework.Gtk = &Gtk{}
	framework.IO = &IO{}
	framework.Process = &Process{}
	framework.Resource = &Resource{}
	framework.Slice = &Slice{}
	framework.Error = &Error{}
	framework.Mail = &MailSender{}
	framework.Crypto = &Crypto{}

	return framework
}

func (f *Framework) getVersion() string {
	return applicationVersion
}
