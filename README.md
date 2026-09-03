# GoFormsDemo

[![CI](https://github.com/Go-Forms/GoFormsDemo/actions/workflows/ci.yml/badge.svg)](https://github.com/Go-Forms/GoFormsDemo/actions/workflows/ci.yml)
![License](https://img.shields.io/badge/license-MIT-blue)

A working example app for [GoForms](https://github.com/Go-Forms/GoForms), showing the
intended project layout and a good chunk of the control catalog wired
together with real event handlers.

## Run it

```
go build -o demo.exe .
./demo.exe
```

(or `go run .`)

## Layout

```
GoFormsDemo/
  Forms/
    MainForm/
      MainForm.go            <- event handlers (btnGreet_Click, clock_Tick, ...)
      MainForm-designer.go   <- control declarations + initializeComponent()
    AboutForm/
      AboutForm.go
      AboutForm-designer.go
  main.go                     <- goforms.NewApplication + goforms.Run(mainForm)
  go.mod                      <- require goforms v0.0.0 + replace goforms => ../GoForms
```

## What it demonstrates

- **Basic input round-trip**: type a name in the `TextBox`, click "Greet" —
  the `Button.Click` handler reads `TextBox.Text()` and updates a `Label`,
  and nudges a `ProgressBar`.
- **GroupBox with relative coordinates**: the "Options" box's `CheckBox`,
  three `RadioButton`s (grouped via `RadioButtonGroup`) and `ComboBox` are all
  positioned relative to the GroupBox itself, not the form.
- **A second form + `ShowDialog`**: "About..." opens `AboutForm` modally from
  a background goroutine (per `Form.ShowDialog`'s threading rule) and reacts
  to its `DialogResult` once closed.
- **`ListBox` + `ContextMenu`**: right-click the fruit list for "Remove
  selected"/"Clear all" (uses the `listWidget`/`contextMenuHost` fix described
  in the GoForms README — this is the one control where the context menu is
  verified to actually work).
- **`Timer`**: a clock label ticks every second, started in `Form.Load` and
  stopped in `Form.Closing`.
- **`PictureBox`**: an in-memory generated placeholder image (no external
  asset files needed).
- **`MenuStrip`**: File → Exit, Help → About.

## Extending the demo

Add a new form the same way `AboutForm` was added: a folder under `Forms/`
with `<Name>.go` + `<Name>-designer.go`, then construct and `Show()`/`ShowDialog()`
it from wherever makes sense (typically a menu item or button handler in
`MainForm.go`).
