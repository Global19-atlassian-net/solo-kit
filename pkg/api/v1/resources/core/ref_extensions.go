package core

func (r ResourceRef) Strings() (string, string, string) {
	return r.Cluster, r.Namespace, r.Name
}

func (r ResourceRef) Key() string {
	key := r.Name
	if r.Namespace != "" {
		key = r.Namespace + "." + key
	}
	if r.Cluster != "" {
		key = r.Cluster + "." + key
	}
	return key
}
