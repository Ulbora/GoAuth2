# Contributing guidelines

## Why
GoAuth2 is written in Golang and if you would like to get experience with the language, this is a good place to start.
If you have an interest in security, this is also a good place to start by helping to build a Golang Oauth2 server solution that is licensed under the GPLV3 license.

## Pull Request Checklist
Before sending your pull requests, make sure you followed this list.

- Read [contributing guidelines](CONTRIBUTING.md).
- Read [Code of Conduct](CODE_OF_CONDUCT.md).


## Project organization

* Branch `master` is always stable and release-ready.
* Branch `develop` is for development and merged into `master` when stable.
* Feature branches should be created for adding new features and merged into `develop` when ready.
* Bug fix branches should be created for fixing bugs and merged into `develop` when ready.

## Submitting a pull request

1. Find an issue to work on, or create a new one. *Avoid duplicates, please check existing issues!*
2. Fork the repo, or make sure you are synced with the latest changes on `develop`.
3. Create a new branch with a sweet name: `git checkout -b issue_<##>_<description>`.
4. Do some programming.
5. Write unit tests.
6. Keep your code nice and clean by adhering to the coding standards & guidelines below.
7. Don't break unit tests or functionality.
8. Update the documentation header comments if needed.
9. **Rebase on `develop` branch and resolve any conflicts _before submitting a pull request!_**
10. Submit a pull request to the `develop` branch.
