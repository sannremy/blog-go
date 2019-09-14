# Continuous Deployment using GitHub Actions

At the time I am writing, GitHub Actions are still in beta version, the syntax changed from HCL to YAML and they introduced a bunch of new features which have piqued my curiosity.

I decided to switch my Continuous Deployment (CD) pipeline from Google Cloud Build to GitHub Actions, and it went surprisingly smooth.

## Automated deployment steps

Here are the steps for [the automated CD pipeline](https://github.com/srchea/homepage/blob/master/.github/workflows/) of this website:

 1. Code checkout from Git.
 2. Build front-end code in Node.js.
 3. Build back-end code in Go.
 4. Authenticate on Google Cloud.
 5. Deploy on Google Cloud App Engine.

### Using GitHub Actions

GitHub Actions were pretty easy to setup: [the documentation](https://help.github.com/en/categories/automating-your-workflow-with-github-actions) is clear and [a lot of resources](https://github.com/actions/) helps to quickly integrate into your existing infrastructure and tools.

```yaml
name: Deploy

on:
  push:
    branches:
    - master
  pull_request:
    branches:
    - master

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout
      uses: actions/checkout@master

    - name: Node install/build
      uses: actions/setup-node@master
      with:
        node-version: 10.x
    - run: yarn install
    - run: yarn run build

    - name: Go build
      uses: actions/setup-go@master
      with:
        go-version: 1.12.x
    - run: go build

    - name: GCP auth
      uses: actions/gcloud/auth@master
      env:
        GCLOUD_AUTH: ${{ secrets.GCLOUD_AUTH }}

    - name: GCP deploy
      uses: actions/gcloud/cli@master
      with:
        args: app deploy --quiet app.yaml
      env:
        CLOUDSDK_CORE_PROJECT: ${{ secrets.CLOUDSDK_CORE_PROJECT }}
```

Secret keys are managed in the settings of the specific repository.

### Using Google Cloud Build

This is the same pipeline using Google Cloud Build. No authentication step is needed here, as the permission is already configured in Cloud Build itself.

```yaml
steps:
- name: node:10-alpine
  entrypoint: yarn
  args: ['install']
  id: '01-yarn-install'

- name: node:10-alpine
  entrypoint: yarn
  args: ['build']
  id: '02-yarn-build'
  waitFor:
  - '01-yarn-install'

- name: golang:1.12
  env: ['GOPATH=/go']
  args: ['go', 'build']
  id: '03-go-build'
  waitFor:
  - '02-yarn-build'

- name: 'gcr.io/cloud-builders/gcloud'
  env: ['GOPATH=/go']
  args: ['app', 'deploy']
  id: '04-gcloud-app-deploy'
  waitFor:
  - '03-go-build'

timeout: '600s'
```

## Execution durations

The whole pipeline runs for an average of 6m 48s on GitHub Actions and 5m 54s on Google Cloud Build. This is acceptable as GitHub Actions is free for open source projects whereas Google Cloud Build gives 2 hours free per month.
