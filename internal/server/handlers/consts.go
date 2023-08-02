package handlers

// k8s Resources
const (
	Deployment  = "deployments"
	Pod         = "pods"
	Service     = "services"
	Deamonset   = "deamonsets"
	Replicaset  = "replicasets"
	Statefulset = "statefulsets"
	Job         = "jobs"
	ConfigMaps  = "configmaps"
	Secret      = "secrets"
	Ingress     = "ingresses"
	Namespace   = "namespaces"
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
