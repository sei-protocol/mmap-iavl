<!--
Guiding Principles:
Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.
Usage:
Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:
* (<tag>) \#<issue-number> message
The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.
Types of changes (Stanzas):
"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking Protobuf, gRPC and REST routes used by end-users.
"CLI Breaking" for breaking CLI commands.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.
Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

* 11 Jan 2023: Initial version started in cronos
* 13 Jan 2023: Change changeset encoding from protobuf to plain one
* 17 Jan 2023:
    * Add delete field to change set to support empty value
    * Add section about compression on snapshot format
* 27 Jan 2023:
    * Update metadata file format
    * Encode key length with 4 bytes instead of 2.
* 24 Feb 2023:
    * Reduce node size without hash from 32bytes to 16bytes, leverage properties of post-order traversal.
    * Merge key-values into single kvs file, build optional MPHF hash table to index it.
* 26 Jul 2023:
  * Forked and modified bu sei