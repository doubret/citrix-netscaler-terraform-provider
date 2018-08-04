package netscaler

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/bindings"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/resources"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Username to login to the NetScaler",
			DefaultFunc: schema.EnvDefaultFunc("NS_LOGIN", "nsroot"),
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password to login to the NetScaler",
			DefaultFunc: schema.EnvDefaultFunc("NS_PASSWORD", "nsroot"),
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The URL to the API",
			DefaultFunc: schema.EnvDefaultFunc("NS_URL", nil),
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"netscaler_appflowaction":                                        resources.NetscalerAppflowaction(),
		"netscaler_appflowcollector":                                     resources.NetscalerAppflowcollector(),
		"netscaler_appflowpolicy":                                        resources.NetscalerAppflowpolicy(),
		"netscaler_appflowpolicylabel":                                   resources.NetscalerAppflowpolicylabel(),
		"netscaler_appfwpolicy":                                          resources.NetscalerAppfwpolicy(),
		"netscaler_appqoeaction":                                         resources.NetscalerAppqoeaction(),
		"netscaler_appqoepolicy":                                         resources.NetscalerAppqoepolicy(),
		"netscaler_auditnslogaction":                                     resources.NetscalerAuditnslogaction(),
		"netscaler_auditnslogpolicy":                                     resources.NetscalerAuditnslogpolicy(),
		"netscaler_auditsyslogaction":                                    resources.NetscalerAuditsyslogaction(),
		"netscaler_auditsyslogpolicy":                                    resources.NetscalerAuditsyslogpolicy(),
		"netscaler_authorizationpolicy":                                  resources.NetscalerAuthorizationpolicy(),
		"netscaler_authorizationpolicylabel":                             resources.NetscalerAuthorizationpolicylabel(),
		"netscaler_caaction":                                             resources.NetscalerCaaction(),
		"netscaler_cachecontentgroup":                                    resources.NetscalerCachecontentgroup(),
		"netscaler_cachepolicy":                                          resources.NetscalerCachepolicy(),
		"netscaler_cachepolicylabel":                                     resources.NetscalerCachepolicylabel(),
		"netscaler_capolicy":                                             resources.NetscalerCapolicy(),
		"netscaler_cmpaction":                                            resources.NetscalerCmpaction(),
		"netscaler_cmppolicy":                                            resources.NetscalerCmppolicy(),
		"netscaler_cmppolicylabel":                                       resources.NetscalerCmppolicylabel(),
		"netscaler_csaction":                                             resources.NetscalerCsaction(),
		"netscaler_cspolicy":                                             resources.NetscalerCspolicy(),
		"netscaler_cspolicylabel":                                        resources.NetscalerCspolicylabel(),
		"netscaler_dbdbprofile":                                          resources.NetscalerDbdbprofile(),
		"netscaler_dnsaction64":                                          resources.NetscalerDnsaction64(),
		"netscaler_dnspolicy64":                                          resources.NetscalerDnspolicy64(),
		"netscaler_dnsprofile":                                           resources.NetscalerDnsprofile(),
		"netscaler_dospolicy":                                            resources.NetscalerDospolicy(),
		"netscaler_feoaction":                                            resources.NetscalerFeoaction(),
		"netscaler_feopolicy":                                            resources.NetscalerFeopolicy(),
		"netscaler_filteraction":                                         resources.NetscalerFilteraction(),
		"netscaler_filterpolicy":                                         resources.NetscalerFilterpolicy(),
		"netscaler_lbgroup":                                              resources.NetscalerLbgroup(),
		"netscaler_lbmetrictable":                                        resources.NetscalerLbmetrictable(),
		"netscaler_lbmonitor":                                            resources.NetscalerLbmonitor(),
		"netscaler_lbprofile":                                            resources.NetscalerLbprofile(),
		"netscaler_lbvserver":                                            resources.NetscalerLbvserver(),
		"netscaler_lbwlm":                                                resources.NetscalerLbwlm(),
		"netscaler_netprofile":                                           resources.NetscalerNetprofile(),
		"netscaler_nshttpprofile":                                        resources.NetscalerNshttpprofile(),
		"netscaler_nstcpprofile":                                         resources.NetscalerNstcpprofile(),
		"netscaler_policydataset":                                        resources.NetscalerPolicydataset(),
		"netscaler_policyexpression":                                     resources.NetscalerPolicyexpression(),
		"netscaler_policypatset":                                         resources.NetscalerPolicypatset(),
		"netscaler_policystringmap":                                      resources.NetscalerPolicystringmap(),
		"netscaler_pqpolicy":                                             resources.NetscalerPqpolicy(),
		"netscaler_responderaction":                                      resources.NetscalerResponderaction(),
		"netscaler_responderpolicy":                                      resources.NetscalerResponderpolicy(),
		"netscaler_responderpolicylabel":                                 resources.NetscalerResponderpolicylabel(),
		"netscaler_rewriteaction":                                        resources.NetscalerRewriteaction(),
		"netscaler_rewritepolicy":                                        resources.NetscalerRewritepolicy(),
		"netscaler_rewritepolicylabel":                                   resources.NetscalerRewritepolicylabel(),
		"netscaler_scpolicy":                                             resources.NetscalerScpolicy(),
		"netscaler_server":                                               resources.NetscalerServer(),
		"netscaler_service":                                              resources.NetscalerService(),
		"netscaler_servicegroup":                                         resources.NetscalerServicegroup(),
		"netscaler_spilloveraction":                                      resources.NetscalerSpilloveraction(),
		"netscaler_spilloverpolicy":                                      resources.NetscalerSpilloverpolicy(),
		"netscaler_tmsessionaction":                                      resources.NetscalerTmsessionaction(),
		"netscaler_tmsessionpolicy":                                      resources.NetscalerTmsessionpolicy(),
		"netscaler_tmtrafficaction":                                      resources.NetscalerTmtrafficaction(),
		"netscaler_tmtrafficpolicy":                                      resources.NetscalerTmtrafficpolicy(),
		"netscaler_transformaction":                                      resources.NetscalerTransformaction(),
		"netscaler_transformpolicy":                                      resources.NetscalerTransformpolicy(),
		"netscaler_transformpolicylabel":                                 resources.NetscalerTransformpolicylabel(),
		"netscaler_transformprofile":                                     resources.NetscalerTransformprofile(),
		"netscaler_videooptimizationaction":                              resources.NetscalerVideooptimizationaction(),
		"netscaler_videooptimizationpolicy":                              resources.NetscalerVideooptimizationpolicy(),
		"netscaler_videooptimizationpolicylabel":                         resources.NetscalerVideooptimizationpolicylabel(),
		"netscaler_appflowglobal_appflowpolicy_binding":                  bindings.NetscalerAppflowglobalAppflowpolicyBinding(),
		"netscaler_appflowpolicylabel_appflowpolicy_binding":             bindings.NetscalerAppflowpolicylabelAppflowpolicyBinding(),
		"netscaler_authorizationpolicylabel_authorizationpolicy_binding": bindings.NetscalerAuthorizationpolicylabelAuthorizationpolicyBinding(),
		"netscaler_lbmetrictable_metric_binding":                         bindings.NetscalerLbmetrictableMetricBinding(),
		"netscaler_lbmonitor_metric_binding":                             bindings.NetscalerLbmonitorMetricBinding(),
		"netscaler_lbmonitor_sslcertkey_binding":                         bindings.NetscalerLbmonitorSslcertkeyBinding(),
		"netscaler_lbvserver_appflowpolicy_binding":                      bindings.NetscalerLbvserverAppflowpolicyBinding(),
		"netscaler_lbvserver_appfwpolicy_binding":                        bindings.NetscalerLbvserverAppfwpolicyBinding(),
		"netscaler_lbvserver_appqoepolicy_binding":                       bindings.NetscalerLbvserverAppqoepolicyBinding(),
		"netscaler_lbvserver_auditnslogpolicy_binding":                   bindings.NetscalerLbvserverAuditnslogpolicyBinding(),
		"netscaler_lbvserver_auditsyslogpolicy_binding":                  bindings.NetscalerLbvserverAuditsyslogpolicyBinding(),
		"netscaler_lbvserver_authorizationpolicy_binding":                bindings.NetscalerLbvserverAuthorizationpolicyBinding(),
		"netscaler_lbvserver_cachepolicy_binding":                        bindings.NetscalerLbvserverCachepolicyBinding(),
		"netscaler_lbvserver_capolicy_binding":                           bindings.NetscalerLbvserverCapolicyBinding(),
		"netscaler_lbvserver_cmppolicy_binding":                          bindings.NetscalerLbvserverCmppolicyBinding(),
		"netscaler_lbvserver_dnspolicy64_binding":                        bindings.NetscalerLbvserverDnspolicy64Binding(),
		"netscaler_lbvserver_feopolicy_binding":                          bindings.NetscalerLbvserverFeopolicyBinding(),
		"netscaler_lbvserver_filterpolicy_binding":                       bindings.NetscalerLbvserverFilterpolicyBinding(),
		"netscaler_lbvserver_pqpolicy_binding":                           bindings.NetscalerLbvserverPqpolicyBinding(),
		"netscaler_lbvserver_responderpolicy_binding":                    bindings.NetscalerLbvserverResponderpolicyBinding(),
		"netscaler_lbvserver_rewritepolicy_binding":                      bindings.NetscalerLbvserverRewritepolicyBinding(),
		"netscaler_lbvserver_scpolicy_binding":                           bindings.NetscalerLbvserverScpolicyBinding(),
		"netscaler_lbvserver_service_binding":                            bindings.NetscalerLbvserverServiceBinding(),
		"netscaler_lbvserver_servicegroup_binding":                       bindings.NetscalerLbvserverServicegroupBinding(),
		"netscaler_lbvserver_spilloverpolicy_binding":                    bindings.NetscalerLbvserverSpilloverpolicyBinding(),
		"netscaler_lbvserver_tmtrafficpolicy_binding":                    bindings.NetscalerLbvserverTmtrafficpolicyBinding(),
		"netscaler_lbvserver_transformpolicy_binding":                    bindings.NetscalerLbvserverTransformpolicyBinding(),
		"netscaler_lbvserver_videooptimizationpolicy_binding":            bindings.NetscalerLbvserverVideooptimizationpolicyBinding(),
		"netscaler_policydataset_value_binding":                          bindings.NetscalerPolicydatasetValueBinding(),
		"netscaler_policypatset_pattern_binding":                         bindings.NetscalerPolicypatsetPatternBinding(),
		"netscaler_policystringmap_pattern_binding":                      bindings.NetscalerPolicystringmapPatternBinding(),
		"netscaler_service_dospolicy_binding":                            bindings.NetscalerServiceDospolicyBinding(),
		"netscaler_service_lbmonitor_binding":                            bindings.NetscalerServiceLbmonitorBinding(),
		"netscaler_service_scpolicy_binding":                             bindings.NetscalerServiceScpolicyBinding(),
		"netscaler_servicegroup_lbmonitor_binding":                       bindings.NetscalerServicegroupLbmonitorBinding(),
		"netscaler_servicegroup_servicegroupmember_binding":              bindings.NetscalerServicegroupServicegroupmemberBinding(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := nitro.NewNitroClient(d.Get("endpoint").(string), d.Get("username").(string), d.Get("password").(string))

	return c, nil
}