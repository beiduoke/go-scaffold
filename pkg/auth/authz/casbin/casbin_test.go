package casbin

import (
	"context"
	"fmt"
	"testing"

	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	"github.com/stretchr/testify/assert"
)

var (
	allProjects = authz.Projects{
		"(unassigned)",
		"project1",
		"project2",
		"project3",
		"project4",
		"project5",
		"project6",
	}
)

func TestFilterAuthorizedPairs(t *testing.T) {
	ctx := context.Background()
	s, err := NewAuthorized(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	policies := map[string]interface{}{
		"policies": []PolicyRule{
			{PType: "p", V0: "bobo", V1: "/api/*", V2: "(GET)|(POST)", V3: "*"},
			{PType: "p", V0: "bobo01", V1: "/api/users", V2: "GET", V3: "*"},
			{PType: "p", V0: "admin_role", V1: "/api/*", V2: "(GET)|(POST)", V3: "*"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "*"},
		},
		"projects": authz.MakeProjects(),
	}

	err = s.SetPolicies(ctx, policies, nil)
	assert.Nil(t, err)

	tests := []struct {
		authorityId string
		path        string
		action      string
		equal       authz.Pairs
	}{
		{
			authorityId: "admin",
			path:        "/api/login",
			action:      "POST",
			equal:       authz.Pairs{authz.MakePair("/api/login", "POST")},
		},
		{
			authorityId: "admin",
			path:        "/api/logout",
			action:      "POST",
			equal:       authz.Pairs{authz.MakePair("/api/logout", "POST")},
		},
		{
			authorityId: "bobo",
			path:        "/api/login",
			action:      "POST",
			equal:       authz.Pairs{authz.MakePair("/api/login", "POST")},
		},
		{
			authorityId: "bobo01",
			path:        "/api/login",
			action:      "POST",
			equal:       authz.Pairs{},
		},
		{
			authorityId: "bobo01",
			path:        "/api/users",
			action:      "GET",
			equal:       authz.Pairs{authz.MakePair("/api/users", "GET")},
		},
		{
			authorityId: "bobo01",
			path:        "/api/users",
			action:      "POST",
			equal:       authz.Pairs{},
		},
	}

	for _, test := range tests {
		t.Run(test.authorityId, func(t *testing.T) {
			subjects := authz.MakeSubjects(authz.Subject(test.authorityId))
			pairs := authz.MakePairs(authz.MakePair(test.path, test.action))
			r, err := s.FilterAuthorizedPairs(ctx, subjects, pairs)
			assert.Nil(t, err)
			assert.EqualValues(t, test.equal, r)
			//fmt.Println(r, err)
		})
	}
}

func TestFilterAuthorizedProjects(t *testing.T) {
	ctx := context.Background()
	s, err := NewAuthorized(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	policies := map[string]interface{}{
		"policies": []PolicyRule{
			{PType: "p", V0: "bobo", V1: "/api/*", V2: "(GET)|(POST)", V3: "project1"},
			{PType: "p", V0: "bobo", V1: "/api/*", V2: "(GET)|(POST)", V3: "project2"},
			{PType: "p", V0: "bobo01", V1: "/api/users", V2: "GET", V3: "*"},
			{PType: "p", V0: "admin_role", V1: "/api/*", V2: "(GET)|(POST)", V3: "*"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "*"},
		},
		"projects": allProjects,
	}

	err = s.SetPolicies(ctx, policies, nil)
	assert.Nil(t, err)

	subjects := authz.Subjects{"bobo"}

	r, err := s.FilterAuthorizedProjects(ctx, subjects)
	assert.Nil(t, err)
	fmt.Println(r)

	tests := []struct {
		subjects authz.Subjects
		equal    authz.Projects
	}{
		{
			subjects: authz.MakeSubjects("bobo"),
			equal:    authz.Projects{"project1", "project2"},
		},
		{
			subjects: authz.MakeSubjects("bobo01"),
			equal:    allProjects,
		},
		{
			subjects: authz.MakeSubjects("admin"),
			equal:    allProjects,
		},
		{
			subjects: authz.MakeSubjects("admin_role"),
			equal:    allProjects,
		},
	}

	for _, test := range tests {
		t.Run(string(test.subjects[0]), func(t *testing.T) {
			r, err := s.FilterAuthorizedProjects(ctx, test.subjects)
			assert.Nil(t, err)
			assert.EqualValues(t, test.equal, r)
			//fmt.Println(r, err)
		})
	}
}

func TestProjectsAuthorized(t *testing.T) {
	ctx := context.Background()
	s, err := NewAuthorized(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	policies := map[string]interface{}{
		"policies": []PolicyRule{
			{PType: "p", V0: "bobo", V1: "/api/*", V2: "(GET)|(POST)", V3: "project1"},
			{PType: "p", V0: "bobo", V1: "/api/*", V2: "(GET)|(POST)", V3: "project2"},
			{PType: "p", V0: "bobo01", V1: "/api/users", V2: "GET", V3: "*"},
			{PType: "p", V0: "admin_role", V1: "/api/*", V2: "(GET)|(POST)", V3: "*"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "project1"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "project2"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "project3"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "project4"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "project5"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "project6"},
			{PType: "g", V0: "admin", V1: "admin_role", V2: "(unassigned)"},
		},
		"projects": allProjects,
	}

	err = s.SetPolicies(ctx, policies, nil)
	assert.Nil(t, err)

	subjects := authz.Subjects{"bobo"}
	action := authz.Action("GET")
	resource := authz.Resource("/api/users")
	projects := authz.Projects{"project1"}
	r, err := s.ProjectsAuthorized(ctx, subjects, action, resource, projects)
	assert.Nil(t, err)
	fmt.Println(r)

	tests := []struct {
		subjects authz.Subjects
		action   authz.Action
		resource authz.Resource
		projects authz.Projects
		equal    authz.Projects
	}{
		{
			subjects: authz.MakeSubjects("bobo"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{"project1"},
		},
		{
			subjects: authz.MakeSubjects("bobo"),
			action:   authz.Action("POST"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{"project1"},
		},
		{
			subjects: authz.MakeSubjects("bobo"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/projects"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{"project1"},
		},
		{
			subjects: authz.MakeSubjects("bobo"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project2"),
			equal:    authz.Projects{"project2"},
		},
		{
			subjects: authz.MakeSubjects("bobo"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project3"),
			equal:    authz.Projects{},
		},
		{
			subjects: authz.MakeSubjects("bobo"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api1/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{},
		},
		{
			subjects: authz.MakeSubjects("bobo"),
			action:   authz.Action("DELETE"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{},
		},
		{
			subjects: authz.MakeSubjects("bobo999"),
			action:   authz.Action("DELETE"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{},
		},
		{
			subjects: authz.MakeSubjects("bobo01"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{"project1"},
		},
		{
			subjects: authz.MakeSubjects("bobo01"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects(allProjects...),
			equal:    allProjects,
		},
		{
			subjects: authz.MakeSubjects("admin"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{"project1"},
		},
		{
			subjects: authz.MakeSubjects("admin"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects(allProjects...),
			equal:    allProjects,
		},
		{
			subjects: authz.MakeSubjects("admin_role"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects("project1"),
			equal:    authz.Projects{"project1"},
		},
		{
			subjects: authz.MakeSubjects("admin_role"),
			action:   authz.Action("GET"),
			resource: authz.Resource("/api/users"),
			projects: authz.MakeProjects(allProjects...),
			equal:    allProjects,
		},
	}

	for _, test := range tests {
		t.Run(string(test.subjects[0]), func(t *testing.T) {
			r, err := s.ProjectsAuthorized(ctx, test.subjects, test.action, test.resource, test.projects)
			assert.Nil(t, err)
			assert.EqualValues(t, test.equal, r)
			//fmt.Println(r, err)
		})
	}
}
