<!--
Thank you for helping to improve Cloud Avenue CLI !
-->

### Description of your changes

<!--
Briefly describe what this pull request does. Be sure to direct your reviewers'
attention to anything that needs special consideration.

We love pull requests that resolve an open issue. If yours does, you
can uncomment the below line to indicate which issue your PR fixes, for example
"Fixes #500":

-->

If you submit change in the provider code, please make sure to:

- [ ] If Needed add a changelog file
- [ ] Write or modify coverage tests
- [ ] Run make generate to ensure the doc was updated properly

### How has this code been coverage and tested
```
go test -coverprofile=coverage.out ./cmd/... && go tool cover -func=coverage.out
.... <PUSH YOUR RESULT HERE>
```

<!--
Before reviewers can be confident in the correctness of this pull request, it
needs to tested and shown to be correct. Briefly describe the testing that has
already been done or which is planned for this change.
-->