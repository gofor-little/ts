# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.2.7 - 2021-08-21
### Added
* Added Go 1.17 support.
* Added a changelog.
* Added a code of conduct.

## v0.2.6 - 2021-02-20
### Added
* Added Go 1.16 support.

## v0.2.5 - 2021-02-05
### Changed
* **BREAKING**: ```Slice.Elements``` is no longer an exported property and now can be access via the ```Slice.GetElements``` method.

### Fixed
* Fixed possible race conditions on ```Slice``` methods.

## v0.2.4 - 2021-02-05
### Fixed
* Fixed possible race conditions on ```LinkedList``` methods.

## v0.2.3 - 2021-02-05
### Fixed
* Fixed possible race conditions on ```Slice``` methods.

## v0.2.2 - 2021-02-05
### Fixed
* Fixed possible race conditions on ```LinkedList``` methods.

## v0.2.1 - 2020-12-31
### Changed
* Updated readme.

## v0.2.0 - 2020-12-31
### Added
* Added Go 1.15 support.
* Added ```Slice.Exists``` method.

## v0.1.0 - 2020-06-11
### Added
* Added ```LinkedList``` struct.
* Added ```Slice``` struct.
