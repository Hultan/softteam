package framework

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Builder struct {
	Builder        *gtk.Builder
}

func NewBuilder(fileNameOrPath string) (*Builder, error) {
	builder := new(Builder)
	err := builder.createBuilder(fileNameOrPath)
	if err != nil {
		return nil, err
	}
	return builder, nil
}

func (s *Builder) createBuilder(fileNameOrPath string) error {
	io := NewIO()
	if !io.FileExists(fileNameOrPath) {
		resource := NewResources()
		fileNameOrPath = resource.GetResourcePath(fileNameOrPath)
	}

	builder, err := gtk.BuilderNewFromFile(fileNameOrPath)
	if err != nil {
		return err
	}

	s.Builder = builder
	return nil
}

func (s *Builder) GetObject(name string) glib.IObject {
	obj, err := s.Builder.GetObject(name)
	if err != nil {
		panic(err)
	}

	return obj
}
