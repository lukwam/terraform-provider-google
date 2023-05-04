// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceSecretManagerSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretManagerSecretCreate,
		Read:   resourceSecretManagerSecretRead,
		Update: resourceSecretManagerSecretUpdate,
		Delete: resourceSecretManagerSecretDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecretManagerSecretImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"replication": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `The replication policy of the secret data attached to the Secret. It cannot be changed
after the Secret has been created.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"automatic": {
							Type:         schema.TypeBool,
							Optional:     true,
							ForceNew:     true,
							Description:  `The Secret will automatically be replicated without any restrictions.`,
							ExactlyOneOf: []string{"replication.0.automatic", "replication.0.user_managed"},
						},
						"user_managed": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `The Secret will automatically be replicated without any restrictions.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"replicas": {
										Type:        schema.TypeList,
										Required:    true,
										ForceNew:    true,
										Description: `The list of Replicas for this Secret. Cannot be empty.`,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"location": {
													Type:        schema.TypeString,
													Required:    true,
													ForceNew:    true,
													Description: `The canonical IDs of the location to replicate data. For example: "us-east1".`,
												},
												"customer_managed_encryption": {
													Type:        schema.TypeList,
													Optional:    true,
													ForceNew:    true,
													Description: `Customer Managed Encryption for the secret.`,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kms_key_name": {
																Type:        schema.TypeString,
																Required:    true,
																ForceNew:    true,
																Description: `Describes the Cloud KMS encryption key that will be used to protect destination secret.`,
															},
														},
													},
												},
											},
										},
									},
								},
							},
							ExactlyOneOf: []string{"replication.0.automatic", "replication.0.user_managed"},
						},
					},
				},
			},
			"secret_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `This must be unique within the project.`,
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels assigned to this Secret.

Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}

Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}

No more than 64 labels can be assigned to a given resource.

An object containing a list of "key": value pairs. Example:
{ "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"rotation": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The rotation time and period for a Secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be set to configure rotation.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"next_rotation_time": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Timestamp in UTC at which the Secret is scheduled to rotate.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
							RequiredWith: []string{"rotation.0.rotation_period"},
						},
						"rotation_period": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The Duration between rotation notifications. Must be in seconds and at least 3600s (1h) and at most 3153600000s (100 years).
If rotationPeriod is set, 'next_rotation_time' must be set. 'next_rotation_time' will be advanced by this period when the service automatically sends rotation notifications.`,
						},
					},
				},
				RequiredWith: []string{"topics"},
			},
			"topics": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*/topics/*.
For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.`,
						},
					},
				},
			},
			"ttl": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The TTL for the Secret.
A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Secret was created.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the Secret. Format:
'projects/{{project}}/secrets/{{secret_id}}'`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecretManagerSecretCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandSecretManagerSecretLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	replicationProp, err := expandSecretManagerSecretReplication(d.Get("replication"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("replication"); !isEmptyValue(reflect.ValueOf(replicationProp)) && (ok || !reflect.DeepEqual(v, replicationProp)) {
		obj["replication"] = replicationProp
	}
	topicsProp, err := expandSecretManagerSecretTopics(d.Get("topics"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("topics"); !isEmptyValue(reflect.ValueOf(topicsProp)) && (ok || !reflect.DeepEqual(v, topicsProp)) {
		obj["topics"] = topicsProp
	}
	expireTimeProp, err := expandSecretManagerSecretExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !isEmptyValue(reflect.ValueOf(expireTimeProp)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandSecretManagerSecretTtl(d.Get("ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl"); !isEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	rotationProp, err := expandSecretManagerSecretRotation(d.Get("rotation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rotation"); !isEmptyValue(reflect.ValueOf(rotationProp)) && (ok || !reflect.DeepEqual(v, rotationProp)) {
		obj["rotation"] = rotationProp
	}

	url, err := ReplaceVars(d, config, "{{SecretManagerBasePath}}projects/{{project}}/secrets?secretId={{secret_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Secret: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Secret: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Secret: %s", err)
	}
	if err := d.Set("name", flattenSecretManagerSecretName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/secrets/{{secret_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Secret %q: %#v", d.Id(), res)

	return resourceSecretManagerSecretRead(d, meta)
}

func resourceSecretManagerSecretRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{SecretManagerBasePath}}projects/{{project}}/secrets/{{secret_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Secret: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecretManagerSecret %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}

	if err := d.Set("name", flattenSecretManagerSecretName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}
	if err := d.Set("create_time", flattenSecretManagerSecretCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}
	if err := d.Set("labels", flattenSecretManagerSecretLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}
	if err := d.Set("replication", flattenSecretManagerSecretReplication(res["replication"], d, config)); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}
	if err := d.Set("topics", flattenSecretManagerSecretTopics(res["topics"], d, config)); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}
	if err := d.Set("expire_time", flattenSecretManagerSecretExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}
	if err := d.Set("rotation", flattenSecretManagerSecretRotation(res["rotation"], d, config)); err != nil {
		return fmt.Errorf("Error reading Secret: %s", err)
	}

	return nil
}

func resourceSecretManagerSecretUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Secret: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandSecretManagerSecretLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	topicsProp, err := expandSecretManagerSecretTopics(d.Get("topics"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("topics"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, topicsProp)) {
		obj["topics"] = topicsProp
	}
	expireTimeProp, err := expandSecretManagerSecretExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	rotationProp, err := expandSecretManagerSecretRotation(d.Get("rotation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rotation"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, rotationProp)) {
		obj["rotation"] = rotationProp
	}

	url, err := ReplaceVars(d, config, "{{SecretManagerBasePath}}projects/{{project}}/secrets/{{secret_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Secret %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("topics") {
		updateMask = append(updateMask, "topics")
	}

	if d.HasChange("expire_time") {
		updateMask = append(updateMask, "expireTime")
	}

	if d.HasChange("rotation") {
		updateMask = append(updateMask, "rotation")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Secret %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Secret %q: %#v", d.Id(), res)
	}

	return resourceSecretManagerSecretRead(d, meta)
}

func resourceSecretManagerSecretDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Secret: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{SecretManagerBasePath}}projects/{{project}}/secrets/{{secret_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Secret %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Secret")
	}

	log.Printf("[DEBUG] Finished deleting Secret %q: %#v", d.Id(), res)
	return nil
}

func resourceSecretManagerSecretImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/secrets/(?P<secret_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<secret_id>[^/]+)",
		"(?P<secret_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/secrets/{{secret_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecretManagerSecretName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretReplication(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["automatic"] =
		flattenSecretManagerSecretReplicationAutomatic(original["automatic"], d, config)
	transformed["user_managed"] =
		flattenSecretManagerSecretReplicationUserManaged(original["userManaged"], d, config)
	return []interface{}{transformed}
}
func flattenSecretManagerSecretReplicationAutomatic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v != nil
}

func flattenSecretManagerSecretReplicationUserManaged(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["replicas"] =
		flattenSecretManagerSecretReplicationUserManagedReplicas(original["replicas"], d, config)
	return []interface{}{transformed}
}
func flattenSecretManagerSecretReplicationUserManagedReplicas(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"location":                    flattenSecretManagerSecretReplicationUserManagedReplicasLocation(original["location"], d, config),
			"customer_managed_encryption": flattenSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption(original["customerManagedEncryption"], d, config),
		})
	}
	return transformed
}
func flattenSecretManagerSecretReplicationUserManagedReplicasLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryptionKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryptionKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretTopics(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name": flattenSecretManagerSecretTopicsName(original["name"], d, config),
		})
	}
	return transformed
}
func flattenSecretManagerSecretTopicsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretRotation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["next_rotation_time"] =
		flattenSecretManagerSecretRotationNextRotationTime(original["nextRotationTime"], d, config)
	transformed["rotation_period"] =
		flattenSecretManagerSecretRotationRotationPeriod(original["rotationPeriod"], d, config)
	return []interface{}{transformed}
}
func flattenSecretManagerSecretRotationNextRotationTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretRotationRotationPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecretManagerSecretLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandSecretManagerSecretReplication(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutomatic, err := expandSecretManagerSecretReplicationAutomatic(original["automatic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutomatic); val.IsValid() && !isEmptyValue(val) {
		transformed["automatic"] = transformedAutomatic
	}

	transformedUserManaged, err := expandSecretManagerSecretReplicationUserManaged(original["user_managed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUserManaged); val.IsValid() && !isEmptyValue(val) {
		transformed["userManaged"] = transformedUserManaged
	}

	return transformed, nil
}

func expandSecretManagerSecretReplicationAutomatic(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || !v.(bool) {
		return nil, nil
	}

	return struct{}{}, nil
}

func expandSecretManagerSecretReplicationUserManaged(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedReplicas, err := expandSecretManagerSecretReplicationUserManagedReplicas(original["replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReplicas); val.IsValid() && !isEmptyValue(val) {
		transformed["replicas"] = transformedReplicas
	}

	return transformed, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicas(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedLocation, err := expandSecretManagerSecretReplicationUserManagedReplicasLocation(original["location"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
			transformed["location"] = transformedLocation
		}

		transformedCustomerManagedEncryption, err := expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption(original["customer_managed_encryption"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCustomerManagedEncryption); val.IsValid() && !isEmptyValue(val) {
			transformed["customerManagedEncryption"] = transformedCustomerManagedEncryption
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicasLocation(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryptionKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryptionKmsKeyName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretTopics(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecretManagerSecretTopicsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecretManagerSecretTopicsName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretExpireTime(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretTtl(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretRotation(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNextRotationTime, err := expandSecretManagerSecretRotationNextRotationTime(original["next_rotation_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNextRotationTime); val.IsValid() && !isEmptyValue(val) {
		transformed["nextRotationTime"] = transformedNextRotationTime
	}

	transformedRotationPeriod, err := expandSecretManagerSecretRotationRotationPeriod(original["rotation_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRotationPeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["rotationPeriod"] = transformedRotationPeriod
	}

	return transformed, nil
}

func expandSecretManagerSecretRotationNextRotationTime(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretRotationRotationPeriod(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
