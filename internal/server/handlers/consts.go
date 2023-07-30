package handlers

// k8s Resources
const (
	Deployment = "deployments"
	Pod        = "pods"
)

// Request type index - (Validating, Mutating)
const (
	RequestTypeIndex = 2
)

// Request Handler to Source
const (
	Validator = "validating"
	Mutator   = "mutating"
)
