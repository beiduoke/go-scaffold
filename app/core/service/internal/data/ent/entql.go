// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/dept"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/member"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/menu"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/post"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/predicate"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/role"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/tenant"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 7)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   dept.Table,
			Columns: dept.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: dept.FieldID,
			},
		},
		Type: "Dept",
		Fields: map[string]*sqlgraph.FieldSpec{
			dept.FieldCreatedAt: {Type: field.TypeTime, Column: dept.FieldCreatedAt},
			dept.FieldUpdatedAt: {Type: field.TypeTime, Column: dept.FieldUpdatedAt},
			dept.FieldDeletedAt: {Type: field.TypeTime, Column: dept.FieldDeletedAt},
			dept.FieldRemark:    {Type: field.TypeString, Column: dept.FieldRemark},
			dept.FieldSort:      {Type: field.TypeInt32, Column: dept.FieldSort},
			dept.FieldState:     {Type: field.TypeInt32, Column: dept.FieldState},
			dept.FieldName:      {Type: field.TypeString, Column: dept.FieldName},
			dept.FieldParentID:  {Type: field.TypeUint32, Column: dept.FieldParentID},
			dept.FieldAncestors: {Type: field.TypeJSON, Column: dept.FieldAncestors},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: member.FieldID,
			},
		},
		Type: "Member",
		Fields: map[string]*sqlgraph.FieldSpec{
			member.FieldCreatedAt:   {Type: field.TypeTime, Column: member.FieldCreatedAt},
			member.FieldUpdatedAt:   {Type: field.TypeTime, Column: member.FieldUpdatedAt},
			member.FieldDeletedAt:   {Type: field.TypeTime, Column: member.FieldDeletedAt},
			member.FieldRemark:      {Type: field.TypeString, Column: member.FieldRemark},
			member.FieldSort:        {Type: field.TypeInt32, Column: member.FieldSort},
			member.FieldState:       {Type: field.TypeInt32, Column: member.FieldState},
			member.FieldUsername:    {Type: field.TypeString, Column: member.FieldUsername},
			member.FieldPassword:    {Type: field.TypeString, Column: member.FieldPassword},
			member.FieldNickname:    {Type: field.TypeString, Column: member.FieldNickname},
			member.FieldPhone:       {Type: field.TypeString, Column: member.FieldPhone},
			member.FieldEmail:       {Type: field.TypeString, Column: member.FieldEmail},
			member.FieldAvatar:      {Type: field.TypeString, Column: member.FieldAvatar},
			member.FieldDescription: {Type: field.TypeString, Column: member.FieldDescription},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: menu.FieldID,
			},
		},
		Type: "Menu",
		Fields: map[string]*sqlgraph.FieldSpec{
			menu.FieldCreatedAt:         {Type: field.TypeTime, Column: menu.FieldCreatedAt},
			menu.FieldUpdatedAt:         {Type: field.TypeTime, Column: menu.FieldUpdatedAt},
			menu.FieldDeletedAt:         {Type: field.TypeTime, Column: menu.FieldDeletedAt},
			menu.FieldRemark:            {Type: field.TypeString, Column: menu.FieldRemark},
			menu.FieldSort:              {Type: field.TypeInt32, Column: menu.FieldSort},
			menu.FieldState:             {Type: field.TypeInt32, Column: menu.FieldState},
			menu.FieldName:              {Type: field.TypeString, Column: menu.FieldName},
			menu.FieldTitle:             {Type: field.TypeString, Column: menu.FieldTitle},
			menu.FieldParentID:          {Type: field.TypeUint32, Column: menu.FieldParentID},
			menu.FieldType:              {Type: field.TypeInt32, Column: menu.FieldType},
			menu.FieldPath:              {Type: field.TypeString, Column: menu.FieldPath},
			menu.FieldComponent:         {Type: field.TypeString, Column: menu.FieldComponent},
			menu.FieldIcon:              {Type: field.TypeString, Column: menu.FieldIcon},
			menu.FieldIsExt:             {Type: field.TypeBool, Column: menu.FieldIsExt},
			menu.FieldExtURL:            {Type: field.TypeString, Column: menu.FieldExtURL},
			menu.FieldPermissions:       {Type: field.TypeJSON, Column: menu.FieldPermissions},
			menu.FieldRedirect:          {Type: field.TypeString, Column: menu.FieldRedirect},
			menu.FieldCurrentActiveMenu: {Type: field.TypeString, Column: menu.FieldCurrentActiveMenu},
			menu.FieldKeepAlive:         {Type: field.TypeBool, Column: menu.FieldKeepAlive},
			menu.FieldVisible:           {Type: field.TypeBool, Column: menu.FieldVisible},
			menu.FieldHideTab:           {Type: field.TypeBool, Column: menu.FieldHideTab},
			menu.FieldHideMenu:          {Type: field.TypeBool, Column: menu.FieldHideMenu},
			menu.FieldHideBreadcrumb:    {Type: field.TypeBool, Column: menu.FieldHideBreadcrumb},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: post.FieldID,
			},
		},
		Type: "Post",
		Fields: map[string]*sqlgraph.FieldSpec{
			post.FieldCreatedAt: {Type: field.TypeTime, Column: post.FieldCreatedAt},
			post.FieldUpdatedAt: {Type: field.TypeTime, Column: post.FieldUpdatedAt},
			post.FieldDeletedAt: {Type: field.TypeTime, Column: post.FieldDeletedAt},
			post.FieldRemark:    {Type: field.TypeString, Column: post.FieldRemark},
			post.FieldSort:      {Type: field.TypeInt32, Column: post.FieldSort},
			post.FieldState:     {Type: field.TypeInt32, Column: post.FieldState},
			post.FieldName:      {Type: field.TypeString, Column: post.FieldName},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: role.FieldID,
			},
		},
		Type: "Role",
		Fields: map[string]*sqlgraph.FieldSpec{
			role.FieldCreatedAt:         {Type: field.TypeTime, Column: role.FieldCreatedAt},
			role.FieldUpdatedAt:         {Type: field.TypeTime, Column: role.FieldUpdatedAt},
			role.FieldDeletedAt:         {Type: field.TypeTime, Column: role.FieldDeletedAt},
			role.FieldRemark:            {Type: field.TypeString, Column: role.FieldRemark},
			role.FieldSort:              {Type: field.TypeInt32, Column: role.FieldSort},
			role.FieldState:             {Type: field.TypeInt32, Column: role.FieldState},
			role.FieldName:              {Type: field.TypeString, Column: role.FieldName},
			role.FieldDefaultRouter:     {Type: field.TypeString, Column: role.FieldDefaultRouter},
			role.FieldDataScope:         {Type: field.TypeInt32, Column: role.FieldDataScope},
			role.FieldMenuCheckStrictly: {Type: field.TypeInt32, Column: role.FieldMenuCheckStrictly},
			role.FieldDeptCheckStrictly: {Type: field.TypeInt32, Column: role.FieldDeptCheckStrictly},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tenant.Table,
			Columns: tenant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: tenant.FieldID,
			},
		},
		Type: "Tenant",
		Fields: map[string]*sqlgraph.FieldSpec{
			tenant.FieldCreatedAt: {Type: field.TypeTime, Column: tenant.FieldCreatedAt},
			tenant.FieldUpdatedAt: {Type: field.TypeTime, Column: tenant.FieldUpdatedAt},
			tenant.FieldDeletedAt: {Type: field.TypeTime, Column: tenant.FieldDeletedAt},
			tenant.FieldRemark:    {Type: field.TypeString, Column: tenant.FieldRemark},
			tenant.FieldSort:      {Type: field.TypeInt32, Column: tenant.FieldSort},
			tenant.FieldState:     {Type: field.TypeInt32, Column: tenant.FieldState},
			tenant.FieldName:      {Type: field.TypeString, Column: tenant.FieldName},
		},
	}
	graph.Nodes[6] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreatedAt:   {Type: field.TypeTime, Column: user.FieldCreatedAt},
			user.FieldUpdatedAt:   {Type: field.TypeTime, Column: user.FieldUpdatedAt},
			user.FieldDeletedAt:   {Type: field.TypeTime, Column: user.FieldDeletedAt},
			user.FieldRemark:      {Type: field.TypeString, Column: user.FieldRemark},
			user.FieldSort:        {Type: field.TypeInt32, Column: user.FieldSort},
			user.FieldState:       {Type: field.TypeInt32, Column: user.FieldState},
			user.FieldName:        {Type: field.TypeString, Column: user.FieldName},
			user.FieldPassword:    {Type: field.TypeString, Column: user.FieldPassword},
			user.FieldNickName:    {Type: field.TypeString, Column: user.FieldNickName},
			user.FieldRealName:    {Type: field.TypeString, Column: user.FieldRealName},
			user.FieldPhone:       {Type: field.TypeString, Column: user.FieldPhone},
			user.FieldEmail:       {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldBirthday:    {Type: field.TypeTime, Column: user.FieldBirthday},
			user.FieldGender:      {Type: field.TypeInt32, Column: user.FieldGender},
			user.FieldAvatar:      {Type: field.TypeString, Column: user.FieldAvatar},
			user.FieldDescription: {Type: field.TypeString, Column: user.FieldDescription},
			user.FieldAuthority:   {Type: field.TypeInt32, Column: user.FieldAuthority},
		},
	}
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dept.ParentTable,
			Columns: []string{dept.ParentColumn},
			Bidi:    false,
		},
		"Dept",
		"Dept",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dept.ChildrenTable,
			Columns: []string{dept.ChildrenColumn},
			Bidi:    false,
		},
		"Dept",
		"Dept",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
		},
		"Menu",
		"Menu",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
		},
		"Menu",
		"Menu",
	)
	graph.MustAddE(
		"users",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
		},
		"Role",
		"User",
	)
	graph.MustAddE(
		"roles",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
		},
		"User",
		"Role",
	)
	graph.MustAddE(
		"posts",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
		},
		"User",
		"Post",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (dq *DeptQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DeptQuery builder.
func (dq *DeptQuery) Filter() *DeptFilter {
	return &DeptFilter{config: dq.config, predicateAdder: dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DeptMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DeptMutation builder.
func (m *DeptMutation) Filter() *DeptFilter {
	return &DeptFilter{config: m.config, predicateAdder: m}
}

// DeptFilter provides a generic filtering capability at runtime for DeptQuery.
type DeptFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DeptFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *DeptFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(dept.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *DeptFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(dept.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *DeptFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(dept.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql time.Time predicate on the deleted_at field.
func (f *DeptFilter) WhereDeletedAt(p entql.TimeP) {
	f.Where(p.Field(dept.FieldDeletedAt))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *DeptFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(dept.FieldRemark))
}

// WhereSort applies the entql int32 predicate on the sort field.
func (f *DeptFilter) WhereSort(p entql.Int32P) {
	f.Where(p.Field(dept.FieldSort))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *DeptFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(dept.FieldState))
}

// WhereName applies the entql string predicate on the name field.
func (f *DeptFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(dept.FieldName))
}

// WhereParentID applies the entql uint32 predicate on the parent_id field.
func (f *DeptFilter) WhereParentID(p entql.Uint32P) {
	f.Where(p.Field(dept.FieldParentID))
}

// WhereAncestors applies the entql json.RawMessage predicate on the ancestors field.
func (f *DeptFilter) WhereAncestors(p entql.BytesP) {
	f.Where(p.Field(dept.FieldAncestors))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *DeptFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *DeptFilter) WhereHasParentWith(preds ...predicate.Dept) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *DeptFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *DeptFilter) WhereHasChildrenWith(preds ...predicate.Dept) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (mq *MemberQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MemberQuery builder.
func (mq *MemberQuery) Filter() *MemberFilter {
	return &MemberFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MemberMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MemberMutation builder.
func (m *MemberMutation) Filter() *MemberFilter {
	return &MemberFilter{config: m.config, predicateAdder: m}
}

// MemberFilter provides a generic filtering capability at runtime for MemberQuery.
type MemberFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MemberFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *MemberFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(member.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *MemberFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(member.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *MemberFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(member.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql time.Time predicate on the deleted_at field.
func (f *MemberFilter) WhereDeletedAt(p entql.TimeP) {
	f.Where(p.Field(member.FieldDeletedAt))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *MemberFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(member.FieldRemark))
}

// WhereSort applies the entql int32 predicate on the sort field.
func (f *MemberFilter) WhereSort(p entql.Int32P) {
	f.Where(p.Field(member.FieldSort))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *MemberFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(member.FieldState))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *MemberFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(member.FieldUsername))
}

// WherePassword applies the entql string predicate on the password field.
func (f *MemberFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(member.FieldPassword))
}

// WhereNickname applies the entql string predicate on the nickname field.
func (f *MemberFilter) WhereNickname(p entql.StringP) {
	f.Where(p.Field(member.FieldNickname))
}

// WherePhone applies the entql string predicate on the phone field.
func (f *MemberFilter) WherePhone(p entql.StringP) {
	f.Where(p.Field(member.FieldPhone))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *MemberFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(member.FieldEmail))
}

// WhereAvatar applies the entql string predicate on the avatar field.
func (f *MemberFilter) WhereAvatar(p entql.StringP) {
	f.Where(p.Field(member.FieldAvatar))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *MemberFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(member.FieldDescription))
}

// addPredicate implements the predicateAdder interface.
func (mq *MenuQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MenuQuery builder.
func (mq *MenuQuery) Filter() *MenuFilter {
	return &MenuFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MenuMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MenuMutation builder.
func (m *MenuMutation) Filter() *MenuFilter {
	return &MenuFilter{config: m.config, predicateAdder: m}
}

// MenuFilter provides a generic filtering capability at runtime for MenuQuery.
type MenuFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MenuFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *MenuFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(menu.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *MenuFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(menu.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *MenuFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(menu.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql time.Time predicate on the deleted_at field.
func (f *MenuFilter) WhereDeletedAt(p entql.TimeP) {
	f.Where(p.Field(menu.FieldDeletedAt))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *MenuFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(menu.FieldRemark))
}

// WhereSort applies the entql int32 predicate on the sort field.
func (f *MenuFilter) WhereSort(p entql.Int32P) {
	f.Where(p.Field(menu.FieldSort))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *MenuFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(menu.FieldState))
}

// WhereName applies the entql string predicate on the name field.
func (f *MenuFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(menu.FieldName))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *MenuFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(menu.FieldTitle))
}

// WhereParentID applies the entql uint32 predicate on the parent_id field.
func (f *MenuFilter) WhereParentID(p entql.Uint32P) {
	f.Where(p.Field(menu.FieldParentID))
}

// WhereType applies the entql int32 predicate on the type field.
func (f *MenuFilter) WhereType(p entql.Int32P) {
	f.Where(p.Field(menu.FieldType))
}

// WherePath applies the entql string predicate on the path field.
func (f *MenuFilter) WherePath(p entql.StringP) {
	f.Where(p.Field(menu.FieldPath))
}

// WhereComponent applies the entql string predicate on the component field.
func (f *MenuFilter) WhereComponent(p entql.StringP) {
	f.Where(p.Field(menu.FieldComponent))
}

// WhereIcon applies the entql string predicate on the icon field.
func (f *MenuFilter) WhereIcon(p entql.StringP) {
	f.Where(p.Field(menu.FieldIcon))
}

// WhereIsExt applies the entql bool predicate on the is_ext field.
func (f *MenuFilter) WhereIsExt(p entql.BoolP) {
	f.Where(p.Field(menu.FieldIsExt))
}

// WhereExtURL applies the entql string predicate on the ext_url field.
func (f *MenuFilter) WhereExtURL(p entql.StringP) {
	f.Where(p.Field(menu.FieldExtURL))
}

// WherePermissions applies the entql json.RawMessage predicate on the permissions field.
func (f *MenuFilter) WherePermissions(p entql.BytesP) {
	f.Where(p.Field(menu.FieldPermissions))
}

// WhereRedirect applies the entql string predicate on the redirect field.
func (f *MenuFilter) WhereRedirect(p entql.StringP) {
	f.Where(p.Field(menu.FieldRedirect))
}

// WhereCurrentActiveMenu applies the entql string predicate on the current_active_menu field.
func (f *MenuFilter) WhereCurrentActiveMenu(p entql.StringP) {
	f.Where(p.Field(menu.FieldCurrentActiveMenu))
}

// WhereKeepAlive applies the entql bool predicate on the keep_alive field.
func (f *MenuFilter) WhereKeepAlive(p entql.BoolP) {
	f.Where(p.Field(menu.FieldKeepAlive))
}

// WhereVisible applies the entql bool predicate on the visible field.
func (f *MenuFilter) WhereVisible(p entql.BoolP) {
	f.Where(p.Field(menu.FieldVisible))
}

// WhereHideTab applies the entql bool predicate on the hide_tab field.
func (f *MenuFilter) WhereHideTab(p entql.BoolP) {
	f.Where(p.Field(menu.FieldHideTab))
}

// WhereHideMenu applies the entql bool predicate on the hide_menu field.
func (f *MenuFilter) WhereHideMenu(p entql.BoolP) {
	f.Where(p.Field(menu.FieldHideMenu))
}

// WhereHideBreadcrumb applies the entql bool predicate on the hide_breadcrumb field.
func (f *MenuFilter) WhereHideBreadcrumb(p entql.BoolP) {
	f.Where(p.Field(menu.FieldHideBreadcrumb))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *MenuFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *MenuFilter) WhereHasParentWith(preds ...predicate.Menu) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *MenuFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *MenuFilter) WhereHasChildrenWith(preds ...predicate.Menu) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (pq *PostQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PostQuery builder.
func (pq *PostQuery) Filter() *PostFilter {
	return &PostFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PostMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PostMutation builder.
func (m *PostMutation) Filter() *PostFilter {
	return &PostFilter{config: m.config, predicateAdder: m}
}

// PostFilter provides a generic filtering capability at runtime for PostQuery.
type PostFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PostFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *PostFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(post.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *PostFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(post.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *PostFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(post.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql time.Time predicate on the deleted_at field.
func (f *PostFilter) WhereDeletedAt(p entql.TimeP) {
	f.Where(p.Field(post.FieldDeletedAt))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *PostFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(post.FieldRemark))
}

// WhereSort applies the entql int32 predicate on the sort field.
func (f *PostFilter) WhereSort(p entql.Int32P) {
	f.Where(p.Field(post.FieldSort))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *PostFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(post.FieldState))
}

// WhereName applies the entql string predicate on the name field.
func (f *PostFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(post.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (rq *RoleQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RoleQuery builder.
func (rq *RoleQuery) Filter() *RoleFilter {
	return &RoleFilter{config: rq.config, predicateAdder: rq}
}

// addPredicate implements the predicateAdder interface.
func (m *RoleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RoleMutation builder.
func (m *RoleMutation) Filter() *RoleFilter {
	return &RoleFilter{config: m.config, predicateAdder: m}
}

// RoleFilter provides a generic filtering capability at runtime for RoleQuery.
type RoleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RoleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *RoleFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(role.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *RoleFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(role.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *RoleFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(role.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql time.Time predicate on the deleted_at field.
func (f *RoleFilter) WhereDeletedAt(p entql.TimeP) {
	f.Where(p.Field(role.FieldDeletedAt))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *RoleFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(role.FieldRemark))
}

// WhereSort applies the entql int32 predicate on the sort field.
func (f *RoleFilter) WhereSort(p entql.Int32P) {
	f.Where(p.Field(role.FieldSort))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *RoleFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(role.FieldState))
}

// WhereName applies the entql string predicate on the name field.
func (f *RoleFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(role.FieldName))
}

// WhereDefaultRouter applies the entql string predicate on the default_router field.
func (f *RoleFilter) WhereDefaultRouter(p entql.StringP) {
	f.Where(p.Field(role.FieldDefaultRouter))
}

// WhereDataScope applies the entql int32 predicate on the data_scope field.
func (f *RoleFilter) WhereDataScope(p entql.Int32P) {
	f.Where(p.Field(role.FieldDataScope))
}

// WhereMenuCheckStrictly applies the entql int32 predicate on the menu_check_strictly field.
func (f *RoleFilter) WhereMenuCheckStrictly(p entql.Int32P) {
	f.Where(p.Field(role.FieldMenuCheckStrictly))
}

// WhereDeptCheckStrictly applies the entql int32 predicate on the dept_check_strictly field.
func (f *RoleFilter) WhereDeptCheckStrictly(p entql.Int32P) {
	f.Where(p.Field(role.FieldDeptCheckStrictly))
}

// WhereHasUsers applies a predicate to check if query has an edge users.
func (f *RoleFilter) WhereHasUsers() {
	f.Where(entql.HasEdge("users"))
}

// WhereHasUsersWith applies a predicate to check if query has an edge users with a given conditions (other predicates).
func (f *RoleFilter) WhereHasUsersWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("users", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tq *TenantQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TenantQuery builder.
func (tq *TenantQuery) Filter() *TenantFilter {
	return &TenantFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TenantMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TenantMutation builder.
func (m *TenantMutation) Filter() *TenantFilter {
	return &TenantFilter{config: m.config, predicateAdder: m}
}

// TenantFilter provides a generic filtering capability at runtime for TenantQuery.
type TenantFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TenantFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *TenantFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(tenant.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *TenantFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(tenant.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *TenantFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(tenant.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql time.Time predicate on the deleted_at field.
func (f *TenantFilter) WhereDeletedAt(p entql.TimeP) {
	f.Where(p.Field(tenant.FieldDeletedAt))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *TenantFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(tenant.FieldRemark))
}

// WhereSort applies the entql int32 predicate on the sort field.
func (f *TenantFilter) WhereSort(p entql.Int32P) {
	f.Where(p.Field(tenant.FieldSort))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *TenantFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(tenant.FieldState))
}

// WhereName applies the entql string predicate on the name field.
func (f *TenantFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(tenant.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[6].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *UserFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *UserFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *UserFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql time.Time predicate on the deleted_at field.
func (f *UserFilter) WhereDeletedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldDeletedAt))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *UserFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(user.FieldRemark))
}

// WhereSort applies the entql int32 predicate on the sort field.
func (f *UserFilter) WhereSort(p entql.Int32P) {
	f.Where(p.Field(user.FieldSort))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *UserFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(user.FieldState))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}

// WhereNickName applies the entql string predicate on the nick_name field.
func (f *UserFilter) WhereNickName(p entql.StringP) {
	f.Where(p.Field(user.FieldNickName))
}

// WhereRealName applies the entql string predicate on the real_name field.
func (f *UserFilter) WhereRealName(p entql.StringP) {
	f.Where(p.Field(user.FieldRealName))
}

// WherePhone applies the entql string predicate on the phone field.
func (f *UserFilter) WherePhone(p entql.StringP) {
	f.Where(p.Field(user.FieldPhone))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WhereBirthday applies the entql time.Time predicate on the birthday field.
func (f *UserFilter) WhereBirthday(p entql.TimeP) {
	f.Where(p.Field(user.FieldBirthday))
}

// WhereGender applies the entql int32 predicate on the gender field.
func (f *UserFilter) WhereGender(p entql.Int32P) {
	f.Where(p.Field(user.FieldGender))
}

// WhereAvatar applies the entql string predicate on the avatar field.
func (f *UserFilter) WhereAvatar(p entql.StringP) {
	f.Where(p.Field(user.FieldAvatar))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *UserFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(user.FieldDescription))
}

// WhereAuthority applies the entql int32 predicate on the authority field.
func (f *UserFilter) WhereAuthority(p entql.Int32P) {
	f.Where(p.Field(user.FieldAuthority))
}

// WhereHasRoles applies a predicate to check if query has an edge roles.
func (f *UserFilter) WhereHasRoles() {
	f.Where(entql.HasEdge("roles"))
}

// WhereHasRolesWith applies a predicate to check if query has an edge roles with a given conditions (other predicates).
func (f *UserFilter) WhereHasRolesWith(preds ...predicate.Role) {
	f.Where(entql.HasEdgeWith("roles", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasPosts applies a predicate to check if query has an edge posts.
func (f *UserFilter) WhereHasPosts() {
	f.Where(entql.HasEdge("posts"))
}

// WhereHasPostsWith applies a predicate to check if query has an edge posts with a given conditions (other predicates).
func (f *UserFilter) WhereHasPostsWith(preds ...predicate.Post) {
	f.Where(entql.HasEdgeWith("posts", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
