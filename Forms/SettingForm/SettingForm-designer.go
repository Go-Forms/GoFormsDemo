package settingform

import "goforms"

// SettingForm-designer.go is the generated-looking half of the WinForms-style
// partial-class split: field declarations and initializeComponent() live
// here. Hand-written logic and event handler bodies belong in SettingForm.go.
// The GoForms Designer regenerates only this file.
type SettingForm struct {
	*goforms.Form
	button1            *goforms.Button
	colorPickerButton1 *goforms.ColorPickerButton
}

// NewSettingForm mirrors "new SettingForm()" - allocate, then initializeComponent.
func NewSettingForm() *SettingForm {
	f := &SettingForm{Form: goforms.NewForm("SettingForm", 500, 350)}
	f.initializeComponent()
	return f
}

// initializeComponent mirrors the WinForms designer's InitializeComponent():
// pure layout, no business logic.
func (f *SettingForm) initializeComponent() {
	f.SetClientSize(500, 350)
	f.CenterOnScreen()
	f.button1 = goforms.NewButton("Button")
	f.button1.SetBounds(20, 20, 90, 30)
	f.AddControl(f.button1)
	f.colorPickerButton1 = goforms.NewColorPickerButton("ColorPickerButton", f.Form)
	f.colorPickerButton1.SetBounds(72, 92, 110, 30)
	f.AddControl(f.colorPickerButton1)
}
