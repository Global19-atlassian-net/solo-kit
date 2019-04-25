package core

func (r ResourceRef) Strings() (string, string) {
	return r.Namespace, r.Name
}

func (r ResourceRef) Key() string {
	key := r.Name
	if r.Namespace != "" {
		key = r.Namespace + "." + key
	}
	return key
}
