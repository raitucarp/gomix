package components

import (
	"github.com/raitucarp/gomix/element"
)

type forEachComponent[T any] struct {
	items      []T
	fallback   IsComponent
	component  *component
	components []IsComponent
}

func ForEach[T any](items ...T) *forEachComponent[T] {
	f := &forEachComponent[T]{
		items:    items,
		fallback: element.Fragment().Element(),
		component: &component{
			el: element.Fragment().Element(),
		},
	}

	return f
}

func (f *forEachComponent[T]) Fallback(fallback IsComponent) *forEachComponent[T] {
	f.fallback = fallback
	return f
}

func (f *forEachComponent[T]) Map(mapFn func(item T, index int) IsComponent) *forEachComponent[T] {
	for index, item := range f.items {
		f.components = append(f.components, mapFn(item, index))
	}
	return f
}

func (f *forEachComponent[T]) Element() *element.HtmlElement {
	if len(f.items) <= 0 {
		return f.fallback.Element()
	}

	el := element.Fragment().Element()
	for _, comp := range f.components {
		el.Children(comp.Element())
	}

	return el
}
