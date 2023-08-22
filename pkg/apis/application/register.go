package application

const (
	// API Group
	Group string = "argoproj.io"

	// Application constants
	ApplicationKind      string = "Application"
	ApplicationSingular  string = "application"
	ApplicationPlural    string = "applications"
	ApplicationShortName string = "app"
	ApplicationFullName  string = ApplicationPlural + "." + Group

	// AppProject constants
	AppProjectKind      string = "AppProject"
	AppProjectSingular  string = "appproject"
	AppProjectPlural    string = "appprojects"
	AppProjectShortName string = "appproject"
	AppProjectFullName  string = AppProjectPlural + "." + Group

	// ApplicationSet constants
	ApplicationSetKind      string = "ApplicationSet"
	ApplicationSetSingular  string = "applicationset"
	ApplicationSetShortName string = "appset"
	ApplicationSetPlural    string = "applicationsets"
	ApplicationSetFullName  string = ApplicationSetPlural + "." + Group

	// SyncStrategy constants
	ApplicationSetSyncStrategyKind      string = "ApplicationSetSyncStrategy"
	ApplicationSetSyncStrategySingular  string = "applicationsetsyncstrategy"
	ApplicationSetSyncStrategyShortName string = "applicationsetsyncstrategy"
	ApplicationSetSyncStrategyPlural    string = "applicationsetsyncstrategies"
	ApplicationSetSyncStrategyFullName  string = ApplicationSetSyncStrategyPlural + "." + Group

	// ClusterSyncStrategy constants
	ClusterApplicationSetSyncStrategyKind             string = "ClusterApplicationSetSyncStrategy"
	ClusterClusterAApplicationSetSyncStrategySingular string = "clusterapplicationsetsyncstrategy"
	ClusterApplicationSetSyncStrategyShortName        string = "clusterapplicationsetsyncstrategy"
	ClusterApplicationSetSyncStrategyPlural           string = "clusterapplicationsetsyncstrategies"
	ClusterApplicationSetSyncStrategyFullName         string = ClusterApplicationSetSyncStrategyPlural + "." + Group
)
