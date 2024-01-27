// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/dept"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/member"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/menu"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/post"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/role"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/schema"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/tenant"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deptMixin := schema.Dept{}.Mixin()
	deptMixinFields0 := deptMixin[0].Fields()
	_ = deptMixinFields0
	deptFields := schema.Dept{}.Fields()
	_ = deptFields
	// deptDescRemark is the schema descriptor for remark field.
	deptDescRemark := deptMixinFields0[4].Descriptor()
	// dept.DefaultRemark holds the default value on creation for the remark field.
	dept.DefaultRemark = deptDescRemark.Default.(string)
	// deptDescSort is the schema descriptor for sort field.
	deptDescSort := deptMixinFields0[5].Descriptor()
	// dept.DefaultSort holds the default value on creation for the sort field.
	dept.DefaultSort = deptDescSort.Default.(int32)
	// dept.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	dept.SortValidator = deptDescSort.Validators[0].(func(int32) error)
	// deptDescState is the schema descriptor for state field.
	deptDescState := deptMixinFields0[6].Descriptor()
	// dept.DefaultState holds the default value on creation for the state field.
	dept.DefaultState = deptDescState.Default.(int32)
	// dept.StateValidator is a validator for the "state" field. It is called by the builders before save.
	dept.StateValidator = deptDescState.Validators[0].(func(int32) error)
	// deptDescName is the schema descriptor for name field.
	deptDescName := deptFields[0].Descriptor()
	// dept.NameValidator is a validator for the "name" field. It is called by the builders before save.
	dept.NameValidator = deptDescName.Validators[0].(func(string) error)
	// deptDescParentID is the schema descriptor for parent_id field.
	deptDescParentID := deptFields[1].Descriptor()
	// dept.DefaultParentID holds the default value on creation for the parent_id field.
	dept.DefaultParentID = deptDescParentID.Default.(int32)
	// deptDescAncestors is the schema descriptor for ancestors field.
	deptDescAncestors := deptFields[2].Descriptor()
	// dept.DefaultAncestors holds the default value on creation for the ancestors field.
	dept.DefaultAncestors = deptDescAncestors.Default.([]int)
	// deptDescID is the schema descriptor for id field.
	deptDescID := deptMixinFields0[0].Descriptor()
	// dept.IDValidator is a validator for the "id" field. It is called by the builders before save.
	dept.IDValidator = deptDescID.Validators[0].(func(uint32) error)
	memberMixin := schema.Member{}.Mixin()
	memberMixinFields0 := memberMixin[0].Fields()
	_ = memberMixinFields0
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescRemark is the schema descriptor for remark field.
	memberDescRemark := memberMixinFields0[4].Descriptor()
	// member.DefaultRemark holds the default value on creation for the remark field.
	member.DefaultRemark = memberDescRemark.Default.(string)
	// memberDescSort is the schema descriptor for sort field.
	memberDescSort := memberMixinFields0[5].Descriptor()
	// member.DefaultSort holds the default value on creation for the sort field.
	member.DefaultSort = memberDescSort.Default.(int32)
	// member.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	member.SortValidator = memberDescSort.Validators[0].(func(int32) error)
	// memberDescState is the schema descriptor for state field.
	memberDescState := memberMixinFields0[6].Descriptor()
	// member.DefaultState holds the default value on creation for the state field.
	member.DefaultState = memberDescState.Default.(int32)
	// member.StateValidator is a validator for the "state" field. It is called by the builders before save.
	member.StateValidator = memberDescState.Validators[0].(func(int32) error)
	// memberDescUsername is the schema descriptor for username field.
	memberDescUsername := memberFields[0].Descriptor()
	// member.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	member.UsernameValidator = func() func(string) error {
		validators := memberDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescPassword is the schema descriptor for password field.
	memberDescPassword := memberFields[1].Descriptor()
	// member.DefaultPassword holds the default value on creation for the password field.
	member.DefaultPassword = memberDescPassword.Default.(string)
	// member.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	member.PasswordValidator = func() func(string) error {
		validators := memberDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescNickname is the schema descriptor for nickname field.
	memberDescNickname := memberFields[2].Descriptor()
	// member.DefaultNickname holds the default value on creation for the nickname field.
	member.DefaultNickname = memberDescNickname.Default.(string)
	// member.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	member.NicknameValidator = func() func(string) error {
		validators := memberDescNickname.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(nickname string) error {
			for _, fn := range fns {
				if err := fn(nickname); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescPhone is the schema descriptor for phone field.
	memberDescPhone := memberFields[3].Descriptor()
	// member.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	member.PhoneValidator = func() func(string) error {
		validators := memberDescPhone.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(phone string) error {
			for _, fn := range fns {
				if err := fn(phone); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescEmail is the schema descriptor for email field.
	memberDescEmail := memberFields[4].Descriptor()
	// member.DefaultEmail holds the default value on creation for the email field.
	member.DefaultEmail = memberDescEmail.Default.(string)
	// member.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	member.EmailValidator = func() func(string) error {
		validators := memberDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescAvatar is the schema descriptor for avatar field.
	memberDescAvatar := memberFields[5].Descriptor()
	// member.DefaultAvatar holds the default value on creation for the avatar field.
	member.DefaultAvatar = memberDescAvatar.Default.(string)
	// member.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	member.AvatarValidator = func() func(string) error {
		validators := memberDescAvatar.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(avatar string) error {
			for _, fn := range fns {
				if err := fn(avatar); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescDescription is the schema descriptor for description field.
	memberDescDescription := memberFields[6].Descriptor()
	// member.DefaultDescription holds the default value on creation for the description field.
	member.DefaultDescription = memberDescDescription.Default.(string)
	// member.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	member.DescriptionValidator = func() func(string) error {
		validators := memberDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescID is the schema descriptor for id field.
	memberDescID := memberMixinFields0[0].Descriptor()
	// member.IDValidator is a validator for the "id" field. It is called by the builders before save.
	member.IDValidator = memberDescID.Validators[0].(func(uint32) error)
	menuMixin := schema.Menu{}.Mixin()
	menuMixinFields0 := menuMixin[0].Fields()
	_ = menuMixinFields0
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescRemark is the schema descriptor for remark field.
	menuDescRemark := menuMixinFields0[4].Descriptor()
	// menu.DefaultRemark holds the default value on creation for the remark field.
	menu.DefaultRemark = menuDescRemark.Default.(string)
	// menuDescSort is the schema descriptor for sort field.
	menuDescSort := menuMixinFields0[5].Descriptor()
	// menu.DefaultSort holds the default value on creation for the sort field.
	menu.DefaultSort = menuDescSort.Default.(int32)
	// menu.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	menu.SortValidator = menuDescSort.Validators[0].(func(int32) error)
	// menuDescState is the schema descriptor for state field.
	menuDescState := menuMixinFields0[6].Descriptor()
	// menu.DefaultState holds the default value on creation for the state field.
	menu.DefaultState = menuDescState.Default.(int32)
	// menu.StateValidator is a validator for the "state" field. It is called by the builders before save.
	menu.StateValidator = menuDescState.Validators[0].(func(int32) error)
	// menuDescName is the schema descriptor for name field.
	menuDescName := menuFields[0].Descriptor()
	// menu.DefaultName holds the default value on creation for the name field.
	menu.DefaultName = menuDescName.Default.(string)
	// menu.NameValidator is a validator for the "name" field. It is called by the builders before save.
	menu.NameValidator = func() func(string) error {
		validators := menuDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// menuDescTitle is the schema descriptor for title field.
	menuDescTitle := menuFields[1].Descriptor()
	// menu.DefaultTitle holds the default value on creation for the title field.
	menu.DefaultTitle = menuDescTitle.Default.(string)
	// menu.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	menu.TitleValidator = menuDescTitle.Validators[0].(func(string) error)
	// menuDescParentID is the schema descriptor for parent_id field.
	menuDescParentID := menuFields[2].Descriptor()
	// menu.DefaultParentID holds the default value on creation for the parent_id field.
	menu.DefaultParentID = menuDescParentID.Default.(uint32)
	// menuDescType is the schema descriptor for type field.
	menuDescType := menuFields[3].Descriptor()
	// menu.DefaultType holds the default value on creation for the type field.
	menu.DefaultType = menuDescType.Default.(int32)
	// menuDescPath is the schema descriptor for path field.
	menuDescPath := menuFields[4].Descriptor()
	// menu.DefaultPath holds the default value on creation for the path field.
	menu.DefaultPath = menuDescPath.Default.(string)
	// menuDescComponent is the schema descriptor for component field.
	menuDescComponent := menuFields[5].Descriptor()
	// menu.DefaultComponent holds the default value on creation for the component field.
	menu.DefaultComponent = menuDescComponent.Default.(string)
	// menuDescIcon is the schema descriptor for icon field.
	menuDescIcon := menuFields[6].Descriptor()
	// menu.DefaultIcon holds the default value on creation for the icon field.
	menu.DefaultIcon = menuDescIcon.Default.(string)
	// menu.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	menu.IconValidator = menuDescIcon.Validators[0].(func(string) error)
	// menuDescIsExt is the schema descriptor for is_ext field.
	menuDescIsExt := menuFields[7].Descriptor()
	// menu.DefaultIsExt holds the default value on creation for the is_ext field.
	menu.DefaultIsExt = menuDescIsExt.Default.(bool)
	// menuDescExtURL is the schema descriptor for ext_url field.
	menuDescExtURL := menuFields[8].Descriptor()
	// menu.DefaultExtURL holds the default value on creation for the ext_url field.
	menu.DefaultExtURL = menuDescExtURL.Default.(string)
	// menu.ExtURLValidator is a validator for the "ext_url" field. It is called by the builders before save.
	menu.ExtURLValidator = menuDescExtURL.Validators[0].(func(string) error)
	// menuDescPermissions is the schema descriptor for permissions field.
	menuDescPermissions := menuFields[9].Descriptor()
	// menu.DefaultPermissions holds the default value on creation for the permissions field.
	menu.DefaultPermissions = menuDescPermissions.Default.([]string)
	// menuDescRedirect is the schema descriptor for redirect field.
	menuDescRedirect := menuFields[10].Descriptor()
	// menu.DefaultRedirect holds the default value on creation for the redirect field.
	menu.DefaultRedirect = menuDescRedirect.Default.(string)
	// menuDescCurrentActiveMenu is the schema descriptor for current_active_menu field.
	menuDescCurrentActiveMenu := menuFields[11].Descriptor()
	// menu.DefaultCurrentActiveMenu holds the default value on creation for the current_active_menu field.
	menu.DefaultCurrentActiveMenu = menuDescCurrentActiveMenu.Default.(string)
	// menuDescKeepAlive is the schema descriptor for keep_alive field.
	menuDescKeepAlive := menuFields[12].Descriptor()
	// menu.DefaultKeepAlive holds the default value on creation for the keep_alive field.
	menu.DefaultKeepAlive = menuDescKeepAlive.Default.(bool)
	// menuDescVisible is the schema descriptor for visible field.
	menuDescVisible := menuFields[13].Descriptor()
	// menu.DefaultVisible holds the default value on creation for the visible field.
	menu.DefaultVisible = menuDescVisible.Default.(bool)
	// menuDescHideTab is the schema descriptor for hide_tab field.
	menuDescHideTab := menuFields[14].Descriptor()
	// menu.DefaultHideTab holds the default value on creation for the hide_tab field.
	menu.DefaultHideTab = menuDescHideTab.Default.(bool)
	// menuDescHideMenu is the schema descriptor for hide_menu field.
	menuDescHideMenu := menuFields[15].Descriptor()
	// menu.DefaultHideMenu holds the default value on creation for the hide_menu field.
	menu.DefaultHideMenu = menuDescHideMenu.Default.(bool)
	// menuDescHideBreadcrumb is the schema descriptor for hide_breadcrumb field.
	menuDescHideBreadcrumb := menuFields[16].Descriptor()
	// menu.DefaultHideBreadcrumb holds the default value on creation for the hide_breadcrumb field.
	menu.DefaultHideBreadcrumb = menuDescHideBreadcrumb.Default.(bool)
	// menuDescID is the schema descriptor for id field.
	menuDescID := menuMixinFields0[0].Descriptor()
	// menu.IDValidator is a validator for the "id" field. It is called by the builders before save.
	menu.IDValidator = menuDescID.Validators[0].(func(uint32) error)
	postMixin := schema.Post{}.Mixin()
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescRemark is the schema descriptor for remark field.
	postDescRemark := postMixinFields0[4].Descriptor()
	// post.DefaultRemark holds the default value on creation for the remark field.
	post.DefaultRemark = postDescRemark.Default.(string)
	// postDescSort is the schema descriptor for sort field.
	postDescSort := postMixinFields0[5].Descriptor()
	// post.DefaultSort holds the default value on creation for the sort field.
	post.DefaultSort = postDescSort.Default.(int32)
	// post.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	post.SortValidator = postDescSort.Validators[0].(func(int32) error)
	// postDescState is the schema descriptor for state field.
	postDescState := postMixinFields0[6].Descriptor()
	// post.DefaultState holds the default value on creation for the state field.
	post.DefaultState = postDescState.Default.(int32)
	// post.StateValidator is a validator for the "state" field. It is called by the builders before save.
	post.StateValidator = postDescState.Validators[0].(func(int32) error)
	// postDescName is the schema descriptor for name field.
	postDescName := postFields[0].Descriptor()
	// post.NameValidator is a validator for the "name" field. It is called by the builders before save.
	post.NameValidator = postDescName.Validators[0].(func(string) error)
	// postDescID is the schema descriptor for id field.
	postDescID := postMixinFields0[0].Descriptor()
	// post.IDValidator is a validator for the "id" field. It is called by the builders before save.
	post.IDValidator = postDescID.Validators[0].(func(uint32) error)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescRemark is the schema descriptor for remark field.
	roleDescRemark := roleMixinFields0[4].Descriptor()
	// role.DefaultRemark holds the default value on creation for the remark field.
	role.DefaultRemark = roleDescRemark.Default.(string)
	// roleDescSort is the schema descriptor for sort field.
	roleDescSort := roleMixinFields0[5].Descriptor()
	// role.DefaultSort holds the default value on creation for the sort field.
	role.DefaultSort = roleDescSort.Default.(int32)
	// role.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	role.SortValidator = roleDescSort.Validators[0].(func(int32) error)
	// roleDescState is the schema descriptor for state field.
	roleDescState := roleMixinFields0[6].Descriptor()
	// role.DefaultState holds the default value on creation for the state field.
	role.DefaultState = roleDescState.Default.(int32)
	// role.StateValidator is a validator for the "state" field. It is called by the builders before save.
	role.StateValidator = roleDescState.Validators[0].(func(int32) error)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleMixinFields0[0].Descriptor()
	// role.IDValidator is a validator for the "id" field. It is called by the builders before save.
	role.IDValidator = roleDescID.Validators[0].(func(uint32) error)
	tenantMixin := schema.Tenant{}.Mixin()
	tenantMixinFields0 := tenantMixin[0].Fields()
	_ = tenantMixinFields0
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescRemark is the schema descriptor for remark field.
	tenantDescRemark := tenantMixinFields0[4].Descriptor()
	// tenant.DefaultRemark holds the default value on creation for the remark field.
	tenant.DefaultRemark = tenantDescRemark.Default.(string)
	// tenantDescSort is the schema descriptor for sort field.
	tenantDescSort := tenantMixinFields0[5].Descriptor()
	// tenant.DefaultSort holds the default value on creation for the sort field.
	tenant.DefaultSort = tenantDescSort.Default.(int32)
	// tenant.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	tenant.SortValidator = tenantDescSort.Validators[0].(func(int32) error)
	// tenantDescState is the schema descriptor for state field.
	tenantDescState := tenantMixinFields0[6].Descriptor()
	// tenant.DefaultState holds the default value on creation for the state field.
	tenant.DefaultState = tenantDescState.Default.(int32)
	// tenant.StateValidator is a validator for the "state" field. It is called by the builders before save.
	tenant.StateValidator = tenantDescState.Validators[0].(func(int32) error)
	// tenantDescName is the schema descriptor for name field.
	tenantDescName := tenantFields[0].Descriptor()
	// tenant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tenant.NameValidator = tenantDescName.Validators[0].(func(string) error)
	// tenantDescID is the schema descriptor for id field.
	tenantDescID := tenantMixinFields0[0].Descriptor()
	// tenant.IDValidator is a validator for the "id" field. It is called by the builders before save.
	tenant.IDValidator = tenantDescID.Validators[0].(func(uint32) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescRemark is the schema descriptor for remark field.
	userDescRemark := userMixinFields0[4].Descriptor()
	// user.DefaultRemark holds the default value on creation for the remark field.
	user.DefaultRemark = userDescRemark.Default.(string)
	// userDescSort is the schema descriptor for sort field.
	userDescSort := userMixinFields0[5].Descriptor()
	// user.DefaultSort holds the default value on creation for the sort field.
	user.DefaultSort = userDescSort.Default.(int32)
	// user.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	user.SortValidator = userDescSort.Validators[0].(func(int32) error)
	// userDescState is the schema descriptor for state field.
	userDescState := userMixinFields0[6].Descriptor()
	// user.DefaultState holds the default value on creation for the state field.
	user.DefaultState = userDescState.Default.(int32)
	// user.StateValidator is a validator for the "state" field. It is called by the builders before save.
	user.StateValidator = userDescState.Validators[0].(func(int32) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[2].Descriptor()
	// user.DefaultNickname holds the default value on creation for the nickname field.
	user.DefaultNickname = userDescNickname.Default.(string)
	// user.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	user.NicknameValidator = func() func(string) error {
		validators := userDescNickname.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(nickname string) error {
			for _, fn := range fns {
				if err := fn(nickname); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[3].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = func() func(string) error {
		validators := userDescPhone.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(phone string) error {
			for _, fn := range fns {
				if err := fn(phone); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[4].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[5].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = func() func(string) error {
		validators := userDescAvatar.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(avatar string) error {
			for _, fn := range fns {
				if err := fn(avatar); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescDescription is the schema descriptor for description field.
	userDescDescription := userFields[6].Descriptor()
	// user.DefaultDescription holds the default value on creation for the description field.
	user.DefaultDescription = userDescDescription.Default.(string)
	// user.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	user.DescriptionValidator = func() func(string) error {
		validators := userDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAuthority is the schema descriptor for authority field.
	userDescAuthority := userFields[7].Descriptor()
	// user.DefaultAuthority holds the default value on creation for the authority field.
	user.DefaultAuthority = userDescAuthority.Default.(int8)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(uint32) error)
}