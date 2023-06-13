package data

import "context"

type Policies struct {
	ID       uint
	Resource string
	Action   string
}

// AddCasbinPolicies 添加
func (d *Data) RoleSetPolicies(ctx context.Context, domainId string, roleId string, rule ...Policies) error {
	var (
		rulePolicies = make([][]string, 0, len(rule))
		mapResource  = make(map[string]uint)
	)
	_, err := d.enforcer.DeleteRolesForUserInDomain(roleId, domainId)
	if err != nil {
		return err
	}
	for _, v := range rule {
		_, ok := mapResource[v.Resource]
		if v.Action == "" {
			v.Action = "*"
		}
		if v.Resource != "" && !ok {
			rulePolicies = append(rulePolicies, []string{v.Resource, v.Action, domainId})
			mapResource[v.Resource] = v.ID
		}
	}
	_, err = d.enforcer.AddPermissionsForUser(roleId, rulePolicies...)
	return err
}
