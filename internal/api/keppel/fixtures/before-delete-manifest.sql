INSERT INTO accounts (name, auth_tenant_id, upstream_peer_hostname, required_labels, metadata_json, next_blob_sweep_at, next_storage_sweep_at, next_federation_announcement_at, in_maintenance, external_peer_url, external_peer_username, external_peer_password, platform_filter, gc_policies_json) VALUES ('test1', 'tenant1', '', '', '', NULL, NULL, NULL, FALSE, '', '', '', '', '[]');
INSERT INTO accounts (name, auth_tenant_id, upstream_peer_hostname, required_labels, metadata_json, next_blob_sweep_at, next_storage_sweep_at, next_federation_announcement_at, in_maintenance, external_peer_url, external_peer_username, external_peer_password, platform_filter, gc_policies_json) VALUES ('test2', 'tenant2', '', '', '', NULL, NULL, NULL, FALSE, '', '', '', '', '[]');

INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:3ee5f0d83bf791f0fb4d750a5719ce19d6d352ef7e5a4264e4b760f0f9c15014', 'application/vnd.docker.distribution.manifest.v2+json', 2000, 12000, 12000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:48341d92e2c078cb4203d231be6402df6794f7114ff465e51174b293caba2438', 'application/vnd.docker.distribution.manifest.v2+json', 8000, 18000, 18000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:69801b353a8f248b5399788172cf8a7758625782651ae1e8b733fd3f5cd875a8', 'application/vnd.docker.distribution.manifest.v2+json', 5000, 15000, 15000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:7b0c710c832f5e2d6ba6b0d459531380d8127d86b6ddf9f5e9e7df2f27f16479', 'application/vnd.docker.distribution.manifest.v2+json', 1000, 11000, 11000, '', 11100, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:937e3ee177ea363e5076a0196bf7bfcbbfc6316a519b4b042cff1f1529584334', 'application/vnd.docker.distribution.manifest.v2+json', 9000, 19000, 19000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:9b8f65607d891ebc9ee18add4f866748456ebce2d8f0bd9c9a8e508871617f27', 'application/vnd.docker.distribution.manifest.v2+json', 10000, 20000, 20000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:cc8cd41cef907c4d216069122c4b89936211361f9050a717a1e37ad1862e952f', 'application/vnd.docker.distribution.manifest.v2+json', 6000, 16000, 16000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:ea3ced18d8c9f5ed3017afbc235db40d7af1a0d3ad50c4d49f7c1549322266c3', 'application/vnd.docker.distribution.manifest.v2+json', 4000, 14000, 14000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:fc1df255dfe9a3d6d2d53746ade768d6cc6578c08b2a4bbc9d6c19153b673791', 'application/vnd.docker.distribution.manifest.v2+json', 3000, 13000, 13000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:ff5a51a65ed0412e35a105b0c9d745e3f8bb49fa64c09f3958230e8bfbbe3272', 'application/vnd.docker.distribution.manifest.v2+json', 7000, 17000, 17000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:22dfc40c154d98ebea12e7f903423d6c49f543265dd67003210b02c013cb637a', 'application/vnd.docker.distribution.manifest.v2+json', 8000, 28000, 28000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:41122349d311a07751ca89355e920157458227652629aa742f3643fbcad246bc', 'application/vnd.docker.distribution.manifest.v2+json', 1000, 21000, 21000, '', 21100, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:543695320e5ed157afafdd9e3f95383c1a27c7d468537fa02389cdaf27b77858', 'application/vnd.docker.distribution.manifest.v2+json', 2000, 22000, 22000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:5b8b4d29020ea5b1bc427c40a0cab2bf944be057ec482110f1d12b68008cd286', 'application/vnd.docker.distribution.manifest.v2+json', 4000, 24000, 24000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:5cc7f71eb464de4bb7ad9c0356a7fe564e4e9ea0df09364c54f495d0dc2c12e6', 'application/vnd.docker.distribution.manifest.v2+json', 10000, 30000, 30000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:726b40b4923b8fc1ff67c2cf4a840ac4b9751f8da18738216d385ad6189c7861', 'application/vnd.docker.distribution.manifest.v2+json', 7000, 27000, 27000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:8865b272185386b62a9cf9c7dd721964e69b7beb4fc161e0faa25339b22f4242', 'application/vnd.docker.distribution.manifest.v2+json', 3000, 23000, 23000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:993c36a18400df7ab61fe4713437cb9ab32947d6b4b9bc319cf5809043bd7adf', 'application/vnd.docker.distribution.manifest.v2+json', 9000, 29000, 29000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:c8139665614f96cfe0af03533e152056b438cc8656d8402fea99e2f275164050', 'application/vnd.docker.distribution.manifest.v2+json', 5000, 25000, 25000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (2, 'sha256:f5f7f84272974fb96e7bc3dec50d8ce5651ccdfbdf2c0bb7145c970ef68fde22', 'application/vnd.docker.distribution.manifest.v2+json', 6000, 26000, 26000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:2432dbb70926df292595fa1c2dee4933cf81130572a2e29a27256d3bdb8e07e5', 'application/vnd.docker.distribution.manifest.v2+json', 10000, 40000, 40000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:37418419cb1f8b74b17c321c44b9e33800ea88ba30a1fb669b1d39e557e18827', 'application/vnd.docker.distribution.manifest.v2+json', 6000, 36000, 36000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:3fea653723dc7414de78df193bba18a242e09ed493e6ec0da8b02bf11267cfbc', 'application/vnd.docker.distribution.manifest.v2+json', 8000, 38000, 38000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:72cd6e8422c407fb6d098690f1130b7ded7ec2f7f5e1d30bd9d521f015363793', 'application/vnd.docker.distribution.manifest.v2+json', 2000, 32000, 32000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:c9bcf7a93bbcc2dec8eb37814135d3f166f971785ff872506e903abed2c49fa5', 'application/vnd.docker.distribution.manifest.v2+json', 5000, 35000, 35000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:cc5f6cf32f361e9fb50506c70a866c25bf5956ff095cd5c968ebb14f22412a63', 'application/vnd.docker.distribution.manifest.v2+json', 1000, 31000, 31000, '', 31100, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:ce041765675ad4d93378e20bd3a7d0d97ddcf3385fb6341581b21d4bc9e3e69e', 'application/vnd.docker.distribution.manifest.v2+json', 3000, 33000, 33000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:d891da1515b09db345eaf7b04e5b8d67597018130fe493cb239d380e8327d4d0', 'application/vnd.docker.distribution.manifest.v2+json', 4000, 34000, 34000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:ecdaf79b192e5eb9d894c4108b530b64a453762b215bf58e3f102b9cdb39c25f', 'application/vnd.docker.distribution.manifest.v2+json', 7000, 37000, 37000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (3, 'sha256:f5576efe561214ce478995fd3cad0181ac257b8fe19f3e84e731c15b45a51776', 'application/vnd.docker.distribution.manifest.v2+json', 9000, 39000, 39000, '', NULL, '{"foo":"is there"}', '{"protected_by_recent_upload":true}', 20001, 20002);

INSERT INTO repos (id, account_name, name, next_blob_mount_sweep_at, next_manifest_sync_at, next_gc_at) VALUES (1, 'test1', 'repo1-1', NULL, NULL, NULL);
INSERT INTO repos (id, account_name, name, next_blob_mount_sweep_at, next_manifest_sync_at, next_gc_at) VALUES (2, 'test1', 'repo1-2', NULL, NULL, NULL);
INSERT INTO repos (id, account_name, name, next_blob_mount_sweep_at, next_manifest_sync_at, next_gc_at) VALUES (3, 'test2', 'repo2-1', NULL, NULL, NULL);

INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (1, 'first', 'sha256:7b0c710c832f5e2d6ba6b0d459531380d8127d86b6ddf9f5e9e7df2f27f16479', 20001, 20101);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (1, 'second', 'sha256:3ee5f0d83bf791f0fb4d750a5719ce19d6d352ef7e5a4264e4b760f0f9c15014', 20003, NULL);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (1, 'stillfirst', 'sha256:7b0c710c832f5e2d6ba6b0d459531380d8127d86b6ddf9f5e9e7df2f27f16479', 20002, NULL);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (2, 'first', 'sha256:41122349d311a07751ca89355e920157458227652629aa742f3643fbcad246bc', 20001, 20101);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (2, 'second', 'sha256:543695320e5ed157afafdd9e3f95383c1a27c7d468537fa02389cdaf27b77858', 20003, NULL);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (2, 'stillfirst', 'sha256:41122349d311a07751ca89355e920157458227652629aa742f3643fbcad246bc', 20002, NULL);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (3, 'first', 'sha256:cc5f6cf32f361e9fb50506c70a866c25bf5956ff095cd5c968ebb14f22412a63', 20001, 20101);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (3, 'second', 'sha256:72cd6e8422c407fb6d098690f1130b7ded7ec2f7f5e1d30bd9d521f015363793', 20003, NULL);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (3, 'stillfirst', 'sha256:cc5f6cf32f361e9fb50506c70a866c25bf5956ff095cd5c968ebb14f22412a63', 20002, NULL);

INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:3ee5f0d83bf791f0fb4d750a5719ce19d6d352ef7e5a4264e4b760f0f9c15014', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:48341d92e2c078cb4203d231be6402df6794f7114ff465e51174b293caba2438', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:69801b353a8f248b5399788172cf8a7758625782651ae1e8b733fd3f5cd875a8', 'Pending', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:7b0c710c832f5e2d6ba6b0d459531380d8127d86b6ddf9f5e9e7df2f27f16479', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:937e3ee177ea363e5076a0196bf7bfcbbfc6316a519b4b042cff1f1529584334', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:9b8f65607d891ebc9ee18add4f866748456ebce2d8f0bd9c9a8e508871617f27', 'Pending', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:cc8cd41cef907c4d216069122c4b89936211361f9050a717a1e37ad1862e952f', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:ea3ced18d8c9f5ed3017afbc235db40d7af1a0d3ad50c4d49f7c1549322266c3', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:fc1df255dfe9a3d6d2d53746ade768d6cc6578c08b2a4bbc9d6c19153b673791', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:ff5a51a65ed0412e35a105b0c9d745e3f8bb49fa64c09f3958230e8bfbbe3272', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:22dfc40c154d98ebea12e7f903423d6c49f543265dd67003210b02c013cb637a', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:41122349d311a07751ca89355e920157458227652629aa742f3643fbcad246bc', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:543695320e5ed157afafdd9e3f95383c1a27c7d468537fa02389cdaf27b77858', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:5b8b4d29020ea5b1bc427c40a0cab2bf944be057ec482110f1d12b68008cd286', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:5cc7f71eb464de4bb7ad9c0356a7fe564e4e9ea0df09364c54f495d0dc2c12e6', 'Pending', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:726b40b4923b8fc1ff67c2cf4a840ac4b9751f8da18738216d385ad6189c7861', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:8865b272185386b62a9cf9c7dd721964e69b7beb4fc161e0faa25339b22f4242', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:993c36a18400df7ab61fe4713437cb9ab32947d6b4b9bc319cf5809043bd7adf', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:c8139665614f96cfe0af03533e152056b438cc8656d8402fea99e2f275164050', 'Pending', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (2, 'sha256:f5f7f84272974fb96e7bc3dec50d8ce5651ccdfbdf2c0bb7145c970ef68fde22', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:2432dbb70926df292595fa1c2dee4933cf81130572a2e29a27256d3bdb8e07e5', 'Pending', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:37418419cb1f8b74b17c321c44b9e33800ea88ba30a1fb669b1d39e557e18827', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:3fea653723dc7414de78df193bba18a242e09ed493e6ec0da8b02bf11267cfbc', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:72cd6e8422c407fb6d098690f1130b7ded7ec2f7f5e1d30bd9d521f015363793', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:c9bcf7a93bbcc2dec8eb37814135d3f166f971785ff872506e903abed2c49fa5', 'Pending', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:cc5f6cf32f361e9fb50506c70a866c25bf5956ff095cd5c968ebb14f22412a63', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:ce041765675ad4d93378e20bd3a7d0d97ddcf3385fb6341581b21d4bc9e3e69e', 'High', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:d891da1515b09db345eaf7b04e5b8d67597018130fe493cb239d380e8327d4d0', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:ecdaf79b192e5eb9d894c4108b530b64a453762b215bf58e3f102b9cdb39c25f', 'Clean', '', 0, NULL, NULL, NULL, '', NULL);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (3, 'sha256:f5576efe561214ce478995fd3cad0181ac257b8fe19f3e84e731c15b45a51776', 'High', '', 0, NULL, NULL, NULL, '', NULL);
