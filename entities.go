package entities

const (
	ObjectTypeTreeMimeType string = "clawio/tree"
)
const (
	ObjectTypeTree ObjectType = iota
	ObjectTypeBLOB
)

type (
	// ObjectType indicates if the object is either a Tree or a BLOB.
	ObjectType int

	// Metadata represents the metadata information retrieved
	// from an object, either tree or blob.
	ObjectInfo interface {
		GetPathSpec() string
		GetSize() uint64
		GetType() ObjectType
		GetMimeType() string
		GetChecksum() string
	}
	// User represents an user of the system.
	// They are created by the authentication service.
	User interface {
		GetUsername() string
		GetEmail() string
		GetDisplayName() string
	}
)
