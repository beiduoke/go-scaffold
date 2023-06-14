package data

import (
	"context"

	"github.com/pkg/errors"
)

type Policy struct {
	// 策略ID
	ID string
	// 策略资源
	Resource string
	// 策略方法/动作
	Action string
}

// AddCasbinPolicy 角色策略设置
func (d *Data) RoleSetPolicy(ctx context.Context, domainId string, roleId string, rule ...Policy) error {
	var (
		rulePolicy  = make([][]string, 0, len(rule))
		mapResource = make(map[string]string)
	)
	for _, v := range rule {
		if v.ID == "" {
			return errors.Errorf("policies %s id not null", v.Resource)
		}
		_, ok := mapResource[v.Resource]
		if v.Action == "" {
			v.Action = "*"
		}
		if v.Resource != "" && !ok {
			rulePolicy = append(rulePolicy, []string{v.Resource, v.Action, domainId})
			mapResource[v.Resource] = v.ID
		}
	}
	_, err := d.enforcer.DeletePermissionsForUser(roleId)
	if err != nil {
		return err
	}
	if len(rulePolicy) > 0 {
		_, err = d.enforcer.AddPermissionsForUser(roleId, rulePolicy...)
	}
	return err
}

// RoleUpdatePolicyResource 角色策略资源更新
func (d *Data) RoleUpdatePolicyResource(ctx context.Context, oldResource string, newResource string) error {
	// 策略ID查询策略列表
	oldPolicies := d.enforcer.GetFilteredPolicy(1, oldResource)
	newPolicies := make([][]string, 0, len(oldPolicies))
	for _, v := range oldPolicies {
		newPolicies = append(newPolicies, []string{v[0], newResource, v[2], v[3]})
	}
	// 批量修改策略资源
	_, err := d.enforcer.UpdatePolicies(oldPolicies, newPolicies)
	return err
}

// RoleDeletePolicyResource 角色策略资源删除
func (d *Data) RoleDeletePolicyResource(ctx context.Context, resource string) error {
	// 批量修改策略资源
	_, err := d.enforcer.DeletePermission(resource)
	return err
}
