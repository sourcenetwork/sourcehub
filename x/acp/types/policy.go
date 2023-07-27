package types

import (
	"github.com/sourcenetwork/sourcehub/utils"
)

// Sort performs an in place sorting of all entities in a Policy.
// Every resource, along with its relations and permissions, are sorted by name.
func (pol *Policy) Sort() {
	resourceExtractor := func(resource *Resource) string { return resource.Name }
	relationExtractor := func(relation *Relation) string { return relation.Name }
	permissionExtractor := func(permission *Permission) string { return permission.Name }

	utils.AsSortable(pol.Resources, resourceExtractor).Sort()

	for _, resource := range pol.Resources {
		utils.AsSortable(resource.Relations, relationExtractor).Sort()
		utils.AsSortable(resource.Permissions, permissionExtractor).Sort()
	}
}
