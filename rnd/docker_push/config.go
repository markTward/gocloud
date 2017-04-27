package main

type App struct {
	Name string
}
type Github struct {
	Repo string
}

// type Registry struct {
// 	Provider map[string][]string
// }
//
type Platform struct {
	Url     string
	Project string
	Repo    string
}

type Workflow struct {
	Enabled bool
	Github  struct {
		Repo   string
		Branch string
	}

	CIProvider struct {
		Name string
		Plan string
	}

	Platform struct {
		Name    string
		Project string
		Cluster string
	}

	CDProvider struct {
		Name      string
		Release   string
		Namespace string
		ChartDir  string
	}
}
