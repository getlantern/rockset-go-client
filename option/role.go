package option

import (
	"github.com/getlantern/rockset-go-client/openapi"
)

// RoleOptions is used to hold optional workspace settings
type RoleOptions struct {
	// Description of the role
	Description *string
	Privileges  []openapi.Privilege
}

type RoleOption func(p *RoleOptions)

// WithRoleDescription is used to optionally set a role description.
func WithRoleDescription(desc string) RoleOption {
	return func(o *RoleOptions) {
		o.Description = &desc
	}
}

// WithGlobalPrivilege is used to add a global action to a role.
func WithGlobalPrivilege(action GlobalAction) RoleOption {
	return func(o *RoleOptions) {
		a := action.String()
		o.Privileges = append(o.Privileges, openapi.Privilege{
			Action: &a,
		})
	}
}

// WithIntegrationPrivilege is used to add an integration action to a role.
func WithIntegrationPrivilege(action IntegrationAction, integration string) RoleOption {
	return func(o *RoleOptions) {
		a := action.String()
		o.Privileges = append(o.Privileges, openapi.Privilege{
			Action:       &a,
			ResourceName: &integration,
		})
	}
}

// WithWorkspacePrivilege is used to add a workspace action to a role. If WithCluster isn't specified, the privilege
// applied to all clusters. Only the last WithCluster is what is used.
func WithWorkspacePrivilege(action WorkspaceAction, workspace string, options ...ClusterPrivileges) RoleOption {
	return func(o *RoleOptions) {
		a := action.String()

		p := openapi.Privilege{
			Action:       &a,
			ResourceName: &workspace,
		}

		for _, opt := range options {
			opt(&p)
		}

		o.Privileges = append(o.Privileges, p)
	}
}

type ClusterPrivileges func(*openapi.Privilege)

// WithCluster is used to specify a cluster for a WithWorkspacePrivilege.
func WithCluster(name string) ClusterPrivileges {
	return func(p *openapi.Privilege) {
		p.Cluster = &name
	}
}

// AllClusters is the string representation of all current and future clusters.
const AllClusters = "*ALL*"

// WithAllCluster is used to specify all current and future clusters for a WithWorkspacePrivilege.
func WithAllCluster() ClusterPrivileges {
	return func(p *openapi.Privilege) {
		all := AllClusters
		p.Cluster = &all
	}
}
