# Table: k8s_core_pvs

This table shows data for Kubernetes (K8s) Core Persistent Volumes (PVs).

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|uid (PK)|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|spec_capacity|JSON|
|spec_persistent_volume_source|JSON|
|spec_access_modes|StringArray|
|spec_claim_ref|JSON|
|spec_persistent_volume_reclaim_policy|String|
|spec_storage_class_name|String|
|spec_mount_options|StringArray|
|spec_volume_mode|String|
|spec_node_affinity|JSON|
|status_phase|String|
|status_message|String|
|status_reason|String|