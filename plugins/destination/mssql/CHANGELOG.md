# Changelog

## [4.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v4.1.0...plugins-destination-mssql-v4.1.1) (2023-05-18)


### Bug Fixes

* Upgrade to plugin-sdk v3.5.2 (Fixes delete-stale for incremental tables) ([#10856](https://github.com/cloudquery/cloudquery/issues/10856)) ([b581394](https://github.com/cloudquery/cloudquery/commit/b5813944da89aa8350be3084783a20c5f1adb49f))

## [4.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v4.0.0...plugins-destination-mssql-v4.1.0) (2023-05-18)


### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* **mssql:** Migrate to `github.com/cloudquery/plugin-sdk/v3` ([#10809](https://github.com/cloudquery/cloudquery/issues/10809)) ([7cacb57](https://github.com/cloudquery/cloudquery/commit/7cacb57dff1f87083238d2a2843b31b130b561a3)), closes [#10724](https://github.com/cloudquery/cloudquery/issues/10724)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v3.0.5...plugins-destination-mssql-v4.0.0) (2023-05-03)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([74271d5](https://github.com/cloudquery/cloudquery/commit/74271d5d9e3a93206cbaeee5540c22c6a32d29a5))

## [3.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v3.0.4...plugins-destination-mssql-v3.0.5) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [3.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v3.0.3...plugins-destination-mssql-v3.0.4) (2023-04-04)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azidentity to v1.2.2 ([#9600](https://github.com/cloudquery/cloudquery/issues/9600)) ([0c5162c](https://github.com/cloudquery/cloudquery/commit/0c5162ce519483909eda8c9bcd7118d145b286c7))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [3.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v3.0.2...plugins-destination-mssql-v3.0.3) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [3.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v3.0.1...plugins-destination-mssql-v3.0.2) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## [3.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v3.0.0...plugins-destination-mssql-v3.0.1) (2023-03-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.4...plugins-destination-mssql-v3.0.0) (2023-03-01)


### ⚠ BREAKING CHANGES

* **deps:** Update MSSQL plugin-sdk to v1.38.2. You'll need to drop the database and start fresh due to a change in the schema for all timestamp columns from `datetimeoffset` to `datetime2`.

### Features

* **deps:** Update MSSQL plugin-sdk to v1.38.2 ([b1870b9](https://github.com/cloudquery/cloudquery/commit/b1870b9dc3698c1306a8d82cf520d1f780bc671d))


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to c95f2b4 ([#8560](https://github.com/cloudquery/cloudquery/issues/8560)) ([9c3bd5b](https://github.com/cloudquery/cloudquery/commit/9c3bd5b68f9741a360fde6c54bf3f5f3efe06d9e))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.3.1 ([#8567](https://github.com/cloudquery/cloudquery/issues/8567)) ([d17e2aa](https://github.com/cloudquery/cloudquery/commit/d17e2aa63e8258ecc3c6815431222d1c5f0a06bf))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [2.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.3...plugins-destination-mssql-v2.1.4) (2023-02-21)


### Bug Fixes

* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## [2.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.2...plugins-destination-mssql-v2.1.3) (2023-02-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [2.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.1...plugins-destination-mssql-v2.1.2) (2023-02-07)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to f062dba ([#7531](https://github.com/cloudquery/cloudquery/issues/7531)) ([59d5575](https://github.com/cloudquery/cloudquery/commit/59d55758b0951553b8d246d1e78b4e3917ff1976))
* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azidentity to v1.2.1 ([#7540](https://github.com/cloudquery/cloudquery/issues/7540)) ([3b5c838](https://github.com/cloudquery/cloudquery/commit/3b5c83832064d729ad1097728f7d12aedbbb9400))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))

## [2.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.0...plugins-destination-mssql-v2.1.1) (2023-01-31)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.0.0...plugins-destination-mssql-v2.1.0) (2023-01-24)


### Features

* **mssql:** Use `LEFT JOIN` for overwrite ([#7086](https://github.com/cloudquery/cloudquery/issues/7086)) ([97c7981](https://github.com/cloudquery/cloudquery/commit/97c79816850b62989670916fd93f4f52528ce2b3))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v1.0.0...plugins-destination-mssql-v2.0.0) (2023-01-23)


### ⚠ BREAKING CHANGES

* **mssql:** Change column types from `varchar` to `nvarchar` to properly store Unicode characters

### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))
* **mssql:** Change column types from `varchar` to `nvarchar` to properly store Unicode characters ([748c7df](https://github.com/cloudquery/cloudquery/commit/748c7df498dfd7ff342c3e380af1ee75f5d1034b))
* **mssql:** Fix JSON transformer to store unescaped data ([748c7df](https://github.com/cloudquery/cloudquery/commit/748c7df498dfd7ff342c3e380af1ee75f5d1034b))

## 1.0.0 (2023-01-19)


### Features

* **mssql:** Microsoft SQL Server destination ([#6417](https://github.com/cloudquery/cloudquery/issues/6417)) ([c71bc88](https://github.com/cloudquery/cloudquery/commit/c71bc887280c9efb05afb98321a611bc33e387b0))
