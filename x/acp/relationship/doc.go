// package relationship deals with relationship persistance for the ACP module.
//
// Relationships are required to meet certain criteria before they can be passed onto Zanzi for storage.
// Due to the discretionary and public nature of the ACP module, prior to relationship storage,
// it's necessary validating that the relationship actor is allowed to create realationship with
// the specified relation for the specified object.
//
// An exemple of that would be: bob tries to submit relationship (file:foo.txt, read, charlie).
// Before storing the relationship the ACP package validates that bob is allowed to create read relations
// for file:foo.txt.
// This validation is done using the manages rules in a policy.
package relationship
