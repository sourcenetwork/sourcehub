package relationship

import "errors"

var ErrRegisterObject = errors.New("register object")
var ErrCannotSetOwnerRelationship = errors.New("cannot create Relationship for owner relation")
