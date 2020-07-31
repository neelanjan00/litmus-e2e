package types

// TestDetails is for collecting all the test-related details
type TestDetails struct {
	ExperimentName         string
	EngineName             string
	OperatorName           string
	ChaosNamespace         string
	RbacPath               string
	ExperimentPath         string
	EnginePath             string
	AnsibleRbacPath        string
	AnsibleExperimentPath  string
	AnsibleEnginePath      string
	AnsibleExperimentImage string
	AppNS                  string
	AppLabel               string
	JobCleanUpPolicy       string
	AnnotationCheck        string
	ApplicationNodeName    string
	GoExperimentImage      string
	InstallLitmus          string
	OperatorImage          string
	ImagePullPolicy        string
	RunnerImage            string
	ChaosDuration          int
	AdminRbacPath          string
}
