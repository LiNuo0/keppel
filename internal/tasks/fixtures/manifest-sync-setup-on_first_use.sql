INSERT INTO accounts (name, auth_tenant_id, upstream_peer_hostname) VALUES ('test1', 'test1authtenant', 'registry.example.org');

INSERT INTO blob_mounts (blob_id, repo_id) VALUES (1, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (2, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (3, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (4, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (5, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (6, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (7, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (8, 1);
INSERT INTO blob_mounts (blob_id, repo_id) VALUES (9, 1);

INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (1, 'test1', 'sha256:804845712601c0fff29e63faaa1804fd15f18bd6206a5a6d3f0c1c78c628eb2d', 1412, '6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b', 3600, 3600, 'application/vnd.docker.container.image.v1+json');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (2, 'test1', 'sha256:eb56d5d5d6a0b061ca49785b5a29e899e68208cdb87853f150e97ecb90d17d92', 1048919, '', 0, 0, 'application/vnd.docker.image.rootfs.diff.tar.gzip');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (3, 'test1', 'sha256:e737c274a038006cd0423ea3526c5c154e025ca2d47c544f54f5a88ee8ac2a94', 1048919, '', 0, 0, 'application/vnd.docker.image.rootfs.diff.tar.gzip');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (4, 'test1', 'sha256:e9a71cfd6a79779b61045293d8046b8ca6b13594883eb498a801266dee31593f', 1412, 'd4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35', 3600, 3600, 'application/vnd.docker.container.image.v1+json');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (5, 'test1', 'sha256:0c9a857a8177888de25f59e4793d785d8c58760cd371dda579f132e897f09401', 1048919, '', 0, 0, 'application/vnd.docker.image.rootfs.diff.tar.gzip');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (6, 'test1', 'sha256:893775dc1030685e5d834c0c4108a7224cfb05440a5e2aba8321142c33cf82e1', 1048919, '', 0, 0, 'application/vnd.docker.image.rootfs.diff.tar.gzip');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (7, 'test1', 'sha256:5f3a9f3054429c026b38e11c2475cc5fdf569ec964e64d13abc2ba769aa5f4c4', 1412, '4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce', 3600, 3600, 'application/vnd.docker.container.image.v1+json');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (8, 'test1', 'sha256:908c681fdc861d81d3f2cf3c760b52c66b126f7e54354d93b7df9a8a2b94e3f2', 1048919, '', 0, 0, 'application/vnd.docker.image.rootfs.diff.tar.gzip');
INSERT INTO blobs (id, account_name, digest, size_bytes, storage_id, pushed_at, validated_at, media_type) VALUES (9, 'test1', 'sha256:ec7b058b0e860e9880dc4827452c379a06b452e11fcbae3e0392174d6493fd62', 1048919, '', 0, 0, 'application/vnd.docker.image.rootfs.diff.tar.gzip');

INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:0b5b811e448f3003b03549e2f2c58c89ca8ea944b4594c20c4ca7ea0885024cb', 7);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:0b5b811e448f3003b03549e2f2c58c89ca8ea944b4594c20c4ca7ea0885024cb', 8);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:0b5b811e448f3003b03549e2f2c58c89ca8ea944b4594c20c4ca7ea0885024cb', 9);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', 1);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', 2);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', 3);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3', 4);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3', 5);
INSERT INTO manifest_blob_refs (repo_id, digest, blob_id) VALUES (1, 'sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3', 6);

INSERT INTO manifest_contents (repo_id, digest, content) VALUES (1, 'sha256:0b5b811e448f3003b03549e2f2c58c89ca8ea944b4594c20c4ca7ea0885024cb', '{"config":{"digest":"sha256:5f3a9f3054429c026b38e11c2475cc5fdf569ec964e64d13abc2ba769aa5f4c4","mediaType":"application/vnd.docker.container.image.v1+json","size":1412},"layers":[{"digest":"sha256:908c681fdc861d81d3f2cf3c760b52c66b126f7e54354d93b7df9a8a2b94e3f2","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919},{"digest":"sha256:ec7b058b0e860e9880dc4827452c379a06b452e11fcbae3e0392174d6493fd62","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919}],"mediaType":"application/vnd.docker.distribution.manifest.v2+json","schemaVersion":2}');
INSERT INTO manifest_contents (repo_id, digest, content) VALUES (1, 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', '{"config":{"digest":"sha256:804845712601c0fff29e63faaa1804fd15f18bd6206a5a6d3f0c1c78c628eb2d","mediaType":"application/vnd.docker.container.image.v1+json","size":1412},"layers":[{"digest":"sha256:eb56d5d5d6a0b061ca49785b5a29e899e68208cdb87853f150e97ecb90d17d92","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919},{"digest":"sha256:e737c274a038006cd0423ea3526c5c154e025ca2d47c544f54f5a88ee8ac2a94","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919}],"mediaType":"application/vnd.docker.distribution.manifest.v2+json","schemaVersion":2}');
INSERT INTO manifest_contents (repo_id, digest, content) VALUES (1, 'sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3', '{"config":{"digest":"sha256:e9a71cfd6a79779b61045293d8046b8ca6b13594883eb498a801266dee31593f","mediaType":"application/vnd.docker.container.image.v1+json","size":1412},"layers":[{"digest":"sha256:0c9a857a8177888de25f59e4793d785d8c58760cd371dda579f132e897f09401","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919},{"digest":"sha256:893775dc1030685e5d834c0c4108a7224cfb05440a5e2aba8321142c33cf82e1","mediaType":"application/vnd.docker.image.rootfs.diff.tar.gzip","size":1048919}],"mediaType":"application/vnd.docker.distribution.manifest.v2+json","schemaVersion":2}');
INSERT INTO manifest_contents (repo_id, digest, content) VALUES (1, 'sha256:edc51c987fd8b320a7496ca6bd97c9e5534368f8f8ee9d7ede8a489ee93fec18', '{"manifests":[{"digest":"sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223","mediaType":"application/vnd.docker.distribution.manifest.v2+json","platform":{"architecture":"amd64","os":"linux"},"size":592},{"digest":"sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3","mediaType":"application/vnd.docker.distribution.manifest.v2+json","platform":{"architecture":"arm","os":"linux"},"size":592}],"mediaType":"application/vnd.docker.distribution.manifest.list.v2+json","schemaVersion":2}');

INSERT INTO manifest_manifest_refs (repo_id, parent_digest, child_digest) VALUES (1, 'sha256:edc51c987fd8b320a7496ca6bd97c9e5534368f8f8ee9d7ede8a489ee93fec18', 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223');
INSERT INTO manifest_manifest_refs (repo_id, parent_digest, child_digest) VALUES (1, 'sha256:edc51c987fd8b320a7496ca6bd97c9e5534368f8f8ee9d7ede8a489ee93fec18', 'sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3');

INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, last_pulled_at, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:0b5b811e448f3003b03549e2f2c58c89ca8ea944b4594c20c4ca7ea0885024cb', 'application/vnd.docker.distribution.manifest.v2+json', 2099842, 3600, 3600, 42, 1, 1);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, last_pulled_at, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', 'application/vnd.docker.distribution.manifest.v2+json', 2099842, 3600, 3600, 32, 1, 1);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, last_pulled_at, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3', 'application/vnd.docker.distribution.manifest.v2+json', 2099842, 3600, 3600, 52, 1, 1);
INSERT INTO manifests (repo_id, digest, media_type, size_bytes, pushed_at, validated_at, min_layer_created_at, max_layer_created_at) VALUES (1, 'sha256:edc51c987fd8b320a7496ca6bd97c9e5534368f8f8ee9d7ede8a489ee93fec18', 'application/vnd.docker.distribution.manifest.list.v2+json', 4200211, 3600, 3600, 1, 1);

INSERT INTO peers (hostname, our_password) VALUES ('registry.example.org', 'a4cb6fae5b8bb91b0b993486937103dab05eca93');

INSERT INTO quotas (auth_tenant_id, manifests) VALUES ('test1authtenant', 100);

INSERT INTO repos (id, account_name, name) VALUES (1, 'test1', 'foo');

INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (1, 'latest', 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', 3600, 32);
INSERT INTO tags (repo_id, name, digest, pushed_at, last_pulled_at) VALUES (1, 'other', 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', 3600, 52);

INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at) VALUES (1, 'sha256:0b5b811e448f3003b03549e2f2c58c89ca8ea944b4594c20c4ca7ea0885024cb', 'Pending', '', 3600);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at) VALUES (1, 'sha256:207a16511ab28a6c3ff0ad6e483ba79fb59a9ebf3721c94e4b91b825bfecf223', 'Pending', '', 3600);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at) VALUES (1, 'sha256:7cefe08689844ee549c7168fd99cf844a7c6117e09e1b728cdd6a18e4645d8b3', 'Pending', '', 3600);
INSERT INTO vuln_info (repo_id, digest, status, message, next_check_at) VALUES (1, 'sha256:edc51c987fd8b320a7496ca6bd97c9e5534368f8f8ee9d7ede8a489ee93fec18', 'Pending', '', 3600);
