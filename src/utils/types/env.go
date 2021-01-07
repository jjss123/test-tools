package types

//Environment variable for injecting

const (
	EnvCheckEnvError          = 123                        // exist code: env check failed
	EnvServiceName            = "JVES_SERVICE_NAME"        // service name
	EnvServiceVersion         = "JVES_SERVICE_VERSION"     // service version/release
	EnvSpaceID                = "JVES_SPACE_ID"            // space id(jvessel used)
	EnvInstanceID             = "JVES_INSTANCE_ID"         // instance id(expose to pass user)
	EnvHostName               = "JVES_HOSTNAME"            // the name of container or vm
	EnvResourceId             = "JVES_RESOURCE_ID"         // the id of container or vm
	EnvKind                   = "JVES_KIND"                // container or vm
	EnvProbeURL               = "JVES_PROBE_URL"           // probe heartbeat url, websocket addr
	EnvProbePeriod            = "JVES_PROBE_PERIOD"        // probe execute heartbeat period
	EnvProbeClientDownloadURL = "JVES_PROBE_DOWNLOAD_URL"  // download probe_client url
	EnvHealthCheckScript      = "JVES_HEALTH_CHECK_SCRIPT" // user health check script
	EnvLogDir                 = "JVES_LOG_DIR"             // user log file dir
	EnvWorkDir                = "JVES_WORK_DIR"            // jvessel log file dir
	EnvLogDirRegex            = "JVES_LOG_DIR_REGEX"       // user log file dir regex
	EnvLogServiceRegex        = "JVES_LOG_SERVICE_REGEX"   // log service file dir regex
	EnvDownloadURL            = "JVES_DOWNLOAD_URL"        // probe client download url
	EnvMetricDir              = "JVES_METRIC_DIR"          // user metric file dir
	EnvBillingDir             = "JVES_BILLING_DIR"         // user billing file dir
	EnvLogbookDir             = "JVES_LOGBOOK_DIR"         // user logbook file dir
	EnvKafkaBrokers           = "JVES_KAFKA_BROKERS"       // kafka brokers
	EnvTopicPrefix            = "JVES_TOPIC_PREFIX"        // kafka topic prefix
	EnvDomain                 = "JVES_DOMAIN"              // domain suffix, xxx.JVESSEL_DOMAIN
	EnvExposedDomain          = "JVES_EXPOSED_DOMAIN"      // exposed domain suffix
	EnvCPU                    = "JVES_CPU"                 // cpu num
	EnvMemory                 = "JVES_MEM"                 // memory num
	EnvDiskSize               = "JVES_LOCALDISK"           // local disk num
	EnvRegion                 = "JVES_REGION"              // region
	EnvIpVersion              = "JVES_IP_VERSION"          // ip version
	EnvDNS                    = "JVES_DNS"
	EnvConfigMap              = "JVES_CONFIGMAP_URL"
	EnvAZ                     = "JVES_AZ"
	EnvPin                    = "JVES_PIN"
)

/* Note:  FIXME: 'ENV in shell code and config file' should keep consistence with 'ENV in go code'

tool/release/satellite.sh
tool/release/filebeat.yml
conf/worker.yaml
*/
