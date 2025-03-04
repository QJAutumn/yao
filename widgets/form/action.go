package form

import (
	"fmt"

	"github.com/yaoapp/gou"
	"github.com/yaoapp/yao/widgets/action"
	"github.com/yaoapp/yao/widgets/hook"
	"github.com/yaoapp/yao/widgets/table"
)

var processActionDefaults = map[string]*action.Process{

	"Setting": {
		Name:    "yao.form.Setting",
		Guard:   "bearer-jwt",
		Process: "yao.form.Xgen",
		Default: []interface{}{nil},
	},
	"Component": {
		Name:    "yao.form.Component",
		Guard:   "bearer-jwt",
		Default: []interface{}{nil, nil, nil},
	},
	"Find": {
		Name:    "yao.form.Find",
		Guard:   "bearer-jwt",
		Default: []interface{}{nil, nil},
	},
	"Save": {
		Name:    "yao.form.Save",
		Guard:   "bearer-jwt",
		Default: []interface{}{nil},
	},
	"Create": {
		Name:    "yao.form.Create",
		Guard:   "bearer-jwt",
		Default: []interface{}{nil},
	},
	"Update": {
		Name:    "yao.form.Update",
		Guard:   "bearer-jwt",
		Default: []interface{}{nil, nil},
	},
	"Delete": {
		Name:    "yao.table.Delete",
		Guard:   "bearer-jwt",
		Default: []interface{}{nil},
	},
}

// SetDefaultProcess set the default value of action
func (act *ActionDSL) SetDefaultProcess() {

	act.Setting = action.ProcessOf(act.Setting).
		Merge(processActionDefaults["Setting"]).
		SetHandler(processHandler)

	act.Component = action.ProcessOf(act.Component).
		Merge(processActionDefaults["Component"]).
		SetHandler(processHandler)

	act.Find = action.ProcessOf(act.Find).
		WithBefore(act.BeforeFind).WithAfter(act.AfterFind).
		Merge(processActionDefaults["Find"]).
		SetHandler(processHandler)

	act.Save = action.ProcessOf(act.Save).
		WithBefore(act.BeforeSave).WithAfter(act.AfterSave).
		Merge(processActionDefaults["Save"]).
		SetHandler(processHandler)

	act.Create = action.ProcessOf(act.Create).
		WithBefore(act.BeforeCreate).WithAfter(act.AfterCreate).
		Merge(processActionDefaults["Create"]).
		SetHandler(processHandler)

	act.Update = action.ProcessOf(act.Update).
		WithBefore(act.BeforeUpdate).WithAfter(act.AfterUpdate).
		Merge(processActionDefaults["Update"]).
		SetHandler(processHandler)

	act.Delete = action.ProcessOf(act.Delete).
		WithBefore(act.BeforeDelete).WithAfter(act.AfterDelete).
		Merge(processActionDefaults["Delete"]).
		SetHandler(processHandler)

}

// BindModel bind model
func (act *ActionDSL) BindModel(m *gou.Model) {

	name := m.ID
	act.Find.Bind(fmt.Sprintf("models.%s.Find", name))
	act.Save.Bind(fmt.Sprintf("models.%s.Save", name))
	act.Create.Bind(fmt.Sprintf("models.%s.Create", name))
	act.Update.Bind(fmt.Sprintf("models.%s.Update", name))
	act.Delete.Bind(fmt.Sprintf("models.%s.Delete", name))

	// bind options
	if act.Bind.Option != nil {
		act.Find.Default[1] = act.Bind.Option
	}
}

// BindForm bind form
func (act *ActionDSL) BindForm(form *DSL) error {

	// Copy Hooks
	hook.CopyBefore(act.BeforeFind, form.Action.BeforeFind)
	hook.CopyBefore(act.BeforeSave, form.Action.BeforeSave)
	hook.CopyBefore(act.BeforeCreate, form.Action.BeforeCreate)
	hook.CopyBefore(act.BeforeUpdate, form.Action.BeforeUpdate)
	hook.CopyBefore(act.BeforeDelete, form.Action.BeforeDelete)
	hook.CopyAfter(act.AfterFind, form.Action.AfterFind)
	hook.CopyAfter(act.AfterSave, form.Action.AfterSave)
	hook.CopyAfter(act.AfterCreate, form.Action.AfterCreate)
	hook.CopyAfter(act.AfterUpdate, form.Action.AfterUpdate)
	hook.CopyAfter(act.AfterDelete, form.Action.AfterDelete)

	// Merge Actions
	act.Find.Merge(form.Action.Find)
	act.Save.Merge(form.Action.Save)
	act.Create.Merge(form.Action.Create)
	act.Update.Merge(form.Action.Update)
	act.Delete.Merge(form.Action.Delete)

	return nil
}

// BindTable bind table
func (act *ActionDSL) BindTable(tab *table.DSL) error {

	// Copy Hooks
	hook.CopyBefore(act.BeforeFind, tab.Action.BeforeFind)
	hook.CopyBefore(act.BeforeSave, tab.Action.BeforeSave)
	hook.CopyBefore(act.BeforeCreate, tab.Action.BeforeCreate)
	hook.CopyBefore(act.BeforeUpdate, tab.Action.BeforeUpdate)
	hook.CopyBefore(act.BeforeDelete, tab.Action.BeforeDelete)
	hook.CopyAfter(act.AfterFind, tab.Action.AfterFind)
	hook.CopyAfter(act.AfterSave, tab.Action.AfterSave)
	hook.CopyAfter(act.AfterCreate, tab.Action.AfterCreate)
	hook.CopyAfter(act.AfterUpdate, tab.Action.AfterUpdate)
	hook.CopyAfter(act.AfterDelete, tab.Action.AfterDelete)

	// Merge Actions
	act.Find.Merge(tab.Action.Find)
	act.Save.Merge(tab.Action.Save)
	act.Create.Merge(tab.Action.Create)
	act.Update.Merge(tab.Action.Update)
	act.Delete.Merge(tab.Action.Delete)

	return nil
}
