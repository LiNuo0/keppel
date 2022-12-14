INSERT INTO accounts (name, auth_tenant_id, upstream_peer_hostname, required_labels, metadata_json, next_blob_sweep_at, next_storage_sweep_at, next_federation_announcement_at, in_maintenance, external_peer_url, external_peer_username, external_peer_password, platform_filter, gc_policies_json) VALUES ('test1', 'test1authtenant', '', '', '', NULL, NULL, NULL, FALSE, '', '', '', '', '[]');

INSERT INTO blob_mounts (blob_id, repo_id, can_be_deleted_at) VALUES (1, 1, NULL);
INSERT INTO blob_mounts (blob_id, repo_id, can_be_deleted_at) VALUES (2, 1, NULL);
INSERT INTO blob_mounts (blob_id, repo_id, can_be_deleted_at) VALUES (3, 1, NULL);

INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (1, 'test1', 'sha256:442f91fa9998460f28e8ff7023e5ddca679f7d2b51dc5498e8aba249678cc7f8', 1048919, '6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b', 3600, 3600, '', NULL, 'application/vnd.docker.image.rootfs.diff.tar.gzip', NULL);
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (2, 'test1', 'sha256:3ae14a50df760250f0e97faf429cc4541c832ed0de61ad5b6ac25d1d695d1a6e', 1048919, 'd4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35', 3600, 3600, '', NULL, 'application/vnd.docker.image.rootfs.diff.tar.gzip', NULL);
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (3, 'test1', 'sha256:92b29e540b6fcadd4e07525af1546c7eff1bb9a8ef0ef249e0b234cdb13dbea3', 1412, '4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce', 3600, 3600, '', NULL, 'application/vnd.docker.container.image.v1+json', NULL);

INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:e255ca60e7cfef94adfcd95d78f1eb44404c4f5887cbf506dd5799489a42606c', 1);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:e255ca60e7cfef94adfcd95d78f1eb44404c4f5887cbf506dd5799489a42606c', 2);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:e255ca60e7cfef94adfcd95d78f1eb44404c4f5887cbf506dd5799489a42606c', 3);

INSERT INTO manifest_contents (repo_id, digest, content) VALUES (1, 'sha256:e255ca60e7cfef94adfcd95d78f1eb44404c4f5887cbf506dd5799489a42606c', '{"config":{"digest":"sha256:92b29e540b6fcadd4e07525af1546c7eff1bb9a8ef0ef249e0b234cdb13dbea3","mediaType":"application/vnd.docker.container.image.v1+json","size":1412},"layers":[{"digest":"sha256:442f91fa9998460f28e8ff7023e5ddca679f7d2b51dc5498e8aba249678cc7f8","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919},{"digest":"sha256:3ae14a50df760250f0e97faf429cc4541c832ed0de61ad5b6ac25d1d695d1a6e","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919}],"mediaType":"application/vnd.docker.distribution.manifest.v2+json","schemaVersion":2}');

INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, validation_error_message, last_pulled_at, labels_json, gc_status_json, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:e255ca60e7cfef94adfcd95d78f1eb44404c4f5887cbf506dd5799489a42606c', 'application/vnd.docker.distribution.manifest.v2+json', 2099842, 3600, 3600, '', NULL, '', '', 1, 1);

INSERT INTO quotas (auth_tenant_id, manifests) VALUES ('test1authtenant', 100);

INSERT INTO repos (id, account_name, name, next_blob_mount_sweep_at, next_manifest_sync_at, next_gc_at) VALUES (1, 'test1', 'foo', 7200, NULL, NULL);

INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at, checked_at, index_started_at, index_finished_at, index_state, check_duration_secs) VALUES (1, 'sha256:e255ca60e7cfef94adfcd95d78f1eb44404c4f5887cbf506dd5799489a42606c', 'Pending', '', 3600, NULL, NULL, NULL, '', NULL);
