package utilfx

import (
	"slices"
	"strings"
	"testing"

	"go.uber.org/fx"
)

type testRunner interface {
	Run() string
}

type testRunnerFunc func() string

func (fn testRunnerFunc) Run() string { return fn() }

type testHandler interface {
	Handle() string
}

type testService struct{}

func (*testService) Handle() string { return "handled" }

type incompatibleService struct{}

func TestTags(t *testing.T) {
	t.Parallel()

	if got := NameTag("analytics"); got != `name:"analytics"` {
		t.Fatalf("NameTag() = %q", got)
	}
	if got := OptionalNameTag("analytics"); got != `name:"analytics" optional:"true"` {
		t.Fatalf("OptionalNameTag() = %q", got)
	}
	if got := GroupTag("handlers"); got != `group:"handlers"` {
		t.Fatalf("GroupTag() = %q", got)
	}
	if got := GroupName("options", "analytics"); got != "options:analytics" {
		t.Fatalf("GroupName() = %q", got)
	}
	if got := GroupName("options", ""); got != "options" {
		t.Fatalf("GroupName() = %q", got)
	}
}

func TestModuleFor(t *testing.T) {
	t.Parallel()

	var called bool
	defaultModule := fx.Provide(func() string { return "default" })
	makeNamed := func(name string) fx.Option {
		called = true
		return fx.Provide(func() string { return name })
	}

	var defaultValue string
	defaultApp := fx.New(
		fx.NopLogger,
		ModuleFor("", defaultModule, makeNamed),
		fx.Invoke(func(value string) { defaultValue = value }),
	)
	requireNoFxErr(t, defaultApp.Err())
	if called {
		t.Fatal("makeNamed was called for an empty name")
	}
	if defaultValue != "default" {
		t.Fatalf("default value = %q", defaultValue)
	}

	var namedValue string
	namedApp := fx.New(
		fx.NopLogger,
		ModuleFor("analytics", defaultModule, makeNamed),
		fx.Invoke(func(value string) { namedValue = value }),
	)
	requireNoFxErr(t, namedApp.Err())
	if !called {
		t.Fatal("makeNamed was not called")
	}
	if namedValue != "analytics" {
		t.Fatalf("named value = %q", namedValue)
	}
}

func TestBind(t *testing.T) {
	t.Parallel()

	var got string
	app := fx.New(
		fx.NopLogger,
		fx.Provide(makeTestService),
		fx.Provide(Bind[testHandler, *testService]),
		fx.Invoke(func(handler testHandler) { got = handler.Handle() }),
	)
	requireNoFxErr(t, app.Err())
	if got != "handled" {
		t.Fatalf("handler result = %q", got)
	}
}

func TestBind_IncompatibleImplementation(t *testing.T) {
	t.Parallel()

	app := fx.New(
		fx.NopLogger,
		fx.Provide(func() *incompatibleService { return &incompatibleService{} }),
		fx.Provide(Bind[testHandler, *incompatibleService]),
		fx.Invoke(func(testHandler) {}),
	)
	if err := app.Err(); err == nil {
		t.Fatal("expected an Fx error")
	} else if !strings.Contains(err.Error(), "does not implement") {
		t.Fatalf("unexpected Fx error: %v", err)
	}
}

func TestGrouped(t *testing.T) {
	t.Parallel()

	var got []string
	app := fx.New(
		fx.NopLogger,
		Grouped("handlers", func() string { return "one" }, nil, func() string { return "two" }),
		fx.Invoke(func(in struct {
			fx.In

			Values []string `group:"handlers"`
		}) {
			got = in.Values
		}),
	)
	requireNoFxErr(t, app.Err())
	requireStrings(t, got, []string{"one", "two"})
}

func TestGroupedFor(t *testing.T) {
	t.Parallel()

	var got []string
	app := fx.New(
		fx.NopLogger,
		GroupedFor("analytics", "options", func() string { return "value" }),
		fx.Invoke(func(in struct {
			fx.In

			Values []string `group:"options:analytics"`
		}) {
			got = in.Values
		}),
	)
	requireNoFxErr(t, app.Err())
	requireStrings(t, got, []string{"value"})
}

func TestGroupedAs(t *testing.T) {
	t.Parallel()

	var got []string
	app := fx.New(
		fx.NopLogger,
		GroupedAs[testRunner](
			"runners",
			func() testRunnerFunc { return func() string { return "first" } },
			func() testRunnerFunc { return func() string { return "second" } },
		),
		fx.Invoke(func(in struct {
			fx.In

			Runners []testRunner `group:"runners"`
		}) {
			got = make([]string, 0, len(in.Runners))
			for _, runner := range in.Runners {
				got = append(got, runner.Run())
			}
		}),
	)
	requireNoFxErr(t, app.Err())
	requireStrings(t, got, []string{"first", "second"})
}

func makeTestService() *testService {
	return &testService{}
}

func requireNoFxErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected Fx error: %v", err)
	}
}

func requireStrings(t *testing.T, got, want []string) {
	t.Helper()

	got = slices.Clone(got)
	want = slices.Clone(want)
	slices.Sort(got)
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Fatalf("strings = %v, want %v", got, want)
	}
}
