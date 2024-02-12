package types

import (
	"github.com/sourcenetwork/sourcehub/utils"
)

const ManagementPermissionPrefix string = "_can_manage_"

// Sort performs an in place sorting of all entities in a Policy.
// Every resource, along with its relations and permissions, are sorted by name.
func (pol *Policy) Sort() {
	resourceExtractor := func(resource *Resource) string { return resource.Name }
	relationExtractor := func(relation *Relation) string { return relation.Name }
	permissionExtractor := func(permission *Permission) string { return permission.Name }

	utils.FromExtractor(pol.Resources, resourceExtractor).SortInPlace()

	for _, resource := range pol.Resources {
		utils.FromExtractor(resource.Relations, relationExtractor).SortInPlace()
		utils.FromExtractor(resource.Permissions, permissionExtractor).SortInPlace()
	}
}

// GetManagementPermissionName returns the name of the Management Permission
// built for the given Relation
func (pol *Policy) GetManagementPermissionName(relation string) string {
	return ManagementPermissionPrefix + relation
}

// GetResourceByName returns a Resource with the given Name.
// If no resource is found with resourceName, return nil
func (pol *Policy) GetResourceByName(resourceName string) *Resource {
	for _, resource := range pol.Resources {
		if resource.Name == resourceName {
			return resource
		}
	}
	return nil
}

// GetPermissionByName returns a Permission with `name`.
// If no Permission matches name, returns nil
func (res *Resource) GetPermissionByName(name string) *Permission {
	for _, permission := range res.Permissions {
		if permission.Name == name {
			return permission
		}
	}
	return nil
}

// GetRelationByName returns a Relation with `name`.
// If no Relation matches name, returns nil
func (res *Resource) GetRelationByName(name string) *Relation {
	for _, relation := range res.Relations {
		if relation.Name == name {
			return relation
		}
	}
	return nil
}
