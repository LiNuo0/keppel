INSERT INTO accounts (name, auth_tenant_id, upstream_peer_hostname, required_labels, metadata_json, next_blob_sweep_at, next_storage_sweep_at, next_federation_announcement_at, in_maintenance, external_peer_url, external_peer_username, external_peer_password, platform_filter, gc_policies_json) VALUES ('test1', 'test1authtenant', '', '', '', 14400, NULL, NULL, FALSE, '', '', '', '', '[]');

INSERT INTO blob_mounts (blob_id, repo_id, can_be_deleted_at) VALUES (4, 1, NULL);
INSERT INTO blob_mounts (blob_id, repo_id, can_be_deleted_at) VALUES (5, 1, NULL);

INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (1, 'test1', 'sha256:2afc94a21f8a7af5b7eac32e3a3acabfd2db3cb80da1631a995eeee413171bc1', 1048919, '6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b', 3600, 3600, '', 12600, '', NULL);
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (2, 'test1', 'sha256:442f91fa9998460f28e8ff7023e5ddca679f7d2b51dc5498e8aba249678cc7f8', 1048919, 'd4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35', 3600, 3600, '', 12600, '', NULL);
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (3, 'test1', 'sha256:3ae14a50df760250f0e97faf429cc4541c832ed0de61ad5b6ac25d1d695d1a6e', 1048919, '4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce', 3600, 3600, '', 12600, '', NULL);
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (4, 'test1', 'sha256:5c82ce800cfc28b78e32a606bcae512145273f6ab39a2b114879beb992329b53', 1048919, '4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a', 3600, 3600, '', NULL, '', NULL);
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, validation_error_message, can_be_deleted_at, media_type, blocks_vuln_scanning) VALUES (5, 'test1', 'sha256:17715f3edba77811b09cc23421c88f0b904c33871f526a1a0a426ad309cbb8ad', 1048919, 'ef2d127de37b942baad06145e54b0c619a1f22327b2ebbcfbec78f5564afe39d', 3600, 3600, '', NULL, '', NULL);

INSERT INTO quotas (auth_tenant_id, manifests) VALUES ('test1authtenant', 100);

INSERT INTO repos (id, account_name, name, next_blob_mount_sweep_at, next_manifest_sync_at, next_gc_at) VALUES (1, 'test1', 'foo', NULL, NULL, NULL);
