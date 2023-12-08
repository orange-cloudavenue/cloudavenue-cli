# Contributing

We welcome contributions of all kinds including code, issues or documentation.

##  Contributing Code

To contribute code, please follow this steps:

1. Communicate with us on the issue you want to work on
2. Make your changes
3. Test your changes
4. Update the documentation if needed and examples too
5. Ensure to run `make generate` without issues
6. Open a pull request
7. If needed, make a changelog of your changes

Ensure to use a good commit hygiene and follow the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.

##  Contributing documentation

Documentation is generated from the code using, you don't need to update the documentation manually in the `docs/` folder.

##  Development environment

Get git submodules:

```console
make submodules
```

Build packages generation:

```console
make build
```

Install CLI locally:

```console
make install
```

Run tests:

```console
make test
```

Test and Build
```console
make generate
```


##  Changelog format

We use the go-changelog to generate and update the changelog from files created in the .changelog/ directory. It is important that when you raise your Pull Request, there is a changelog entry which describes the changes your contribution makes. Not all changes require an entry in the changelog, guidance follows on what changes do.

The changelog format requires an entry in the following format, where HEADER corresponds to the changelog category, and the entry is the changelog entry itself. The entry should be included in a file in the .changelog directory with the naming convention {PR-NUMBER}.txt. For example, to create a changelog entry for pull request 1234, there should be a file named .changelog/1234.txt.

``````markdown
```release-note:{HEADER}
{ENTRY}
```
``````

## Pull request types to CHANGELOG

The CHANGELOG is intended to show operator-impacting changes to the codebase for a particular version. If every change or commit to the code resulted in an entry, the CHANGELOG would become less useful for operators. The lists below are general guidelines and examples for when a decision needs to be made to decide whether a change should have an entry.

### Changes that should have a CHANGELOG entry

#### New feature

A new feature entry should only contain the name of the feat, and use the `release-note:feature` header.

``````markdown
```release-note:feature
New vdc
```
``````

#### Bug fixes

A new bug entry should use the `release-note:bug` header and have a prefix indicating the command and subcommand it corresponds to, a colon, then followed by a brief summary.

``````markdown
```release-note:bug
s3/List: Fix wrong print
```
``````

#### Resource and provider enhancements

A new enhancement entry should use the `release-note:enhancement` header and have a prefix indicating the command and subcommand it corresponds to, a colon, then followed by a brief summary.

``````markdown
```release-note:enhancement
vdc/create: Add new argument
```
``````

#### Deprecations

A deprecation entry should use the `release-note:note` header and have a prefix indicating the command and subcommand it corresponds to, a colon, then followed by a brief summary.

``````markdown
```release-note:note
t0/create: The old_subcommand create is being deprecated in favor of the new_subcommand add to support new feature
```
``````

#### Breaking changes and removals

A breaking-change entry should use the `release-note:breaking-change` header and have a prefix indicating the command and subcommand it corresponds to, a colon, then followed by a brief summary.

``````markdown
```release-note:breaking-change
vdc/delete: This is a breaking change
```
``````

### Changes that should _not_ have a CHANGELOG entry

- Documentation updates
- Testing updates
- Code refactoring
