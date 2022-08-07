package topology

type Topology struct {
	Services   []*ServiceTier `json:"services" yaml:"services"`
	ServiceMap map[string]*ServiceTier
}

func (t *Topology) GetServiceTier(serviceName string) *ServiceTier {
	return t.ServiceMap[serviceName]
}

func (t *Topology) LoadServiceMap() {
	t.ServiceMap = make(map[string]*ServiceTier)
	for _, s := range t.Services {
		t.ServiceMap[s.ServiceName] = s
		s.LoadRouteMap()
	}
}
