package utils

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"k8s.io/klog/v2"
)

type LdapConfig struct {
	UserBase            string
	GroupBase           string
	AppGroupBase        string
	CustomerGroupBase   string
	OpsGroupBase        string
	AdminGroupBase      string
	Host                string
	Port                int
	UseSSL              bool
	StartTLS            bool
	SkipTLSVerification bool
	BindDN              string
	BindPassword        string
	UserFilter          string
	UserKey             string
	GroupFilter         string
	Attributes          []string
}

func LoadConfig() LdapConfig {

	env := os.Getenv("GO_DOT_ENV")
	if env != "" {
		filePath := filepath.Join("dev", ".env")

		err := godotenv.Load(filePath)
		if err != nil {
			klog.Warningf("failed to load environment file: %v", err)
		}
	}

	ldapPort, errLdapPort := strconv.Atoi(getEnv("LDAP_PORT", "389"))
	Checkf(errLdapPort, "Invalid LDAP_PORT, must be an integer")

	useSSL, errLdapSSL := strconv.ParseBool(getEnv("LDAP_USE_SSL", "false"))
	Checkf(errLdapSSL, "Invalid LDAP_USE_SSL, must be a boolean")

	skipTLSVerification, errSkipTLS := strconv.ParseBool(getEnv("LDAP_SKIP_TLS_VERIFICATION", "true"))
	Checkf(errSkipTLS, "Invalid LDAP_SKIP_TLS_VERIFICATION, must be a boolean")

	startTLS, errStartTLS := strconv.ParseBool(getEnv("LDAP_START_TLS", "false"))
	Checkf(errStartTLS, "Invalid LDAP_START_TLS, must be a boolean")

	if len(os.Getenv("LDAP_PORT")) > 0 {
		envLdapPort, err := strconv.Atoi(os.Getenv("LDAP_PORT"))
		Check(err)
		ldapPort = envLdapPort
		if ldapPort == 389 && os.Getenv("LDAP_SKIP_TLS") == "false" {
			skipTLSVerification = false
		}
		if ldapPort == 636 && os.Getenv("LDAP_SKIP_TLS") == "false" {
			skipTLSVerification = false
			useSSL = true
		}
	}

	ldapUserFilter := getEnv("LDAP_USERFILTER", "(cn=%s)")

	ldapConfig := LdapConfig{
		UserBase:            os.Getenv("LDAP_USERBASE"),
		UserKey:             os.Getenv("LDAP_USERKEY"),
		GroupBase:           os.Getenv("LDAP_GROUPBASE"),
		AppGroupBase:        os.Getenv("LDAP_APP_GROUPBASE"),
		CustomerGroupBase:   os.Getenv("LDAP_CUSTOMER_OPS_GROUPBASE"),
		OpsGroupBase:        os.Getenv("LDAP_OPS_GROUPBASE"),
		AdminGroupBase:      os.Getenv("LDAP_ADMINS_GROUPBASE"),
		Host:                os.Getenv("LDAP_SERVER"),
		Port:                ldapPort,
		UseSSL:              useSSL,
		StartTLS:            startTLS,
		SkipTLSVerification: skipTLSVerification,
		BindDN:              os.Getenv("LDAP_BINDDN"),
		BindPassword:        os.Getenv("LDAP_PASSWD"),
		UserFilter:          ldapUserFilter,
		GroupFilter:         "(member=%s)",
		Attributes:          []string{"givenName", "sn", "mail", "uid", "cn", "userPrincipalName"},
	}

	return ldapConfig

}
