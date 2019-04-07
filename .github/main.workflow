workflow "Deploy" {
  on = "push"
  resolves = ["Filter branch master", "Build", "GCP Authenticate", "GCP Deploy"]
}

action "Filter branch master" {
  uses = "actions/bin/filter@master"
  args = "branch master"
}

action "Build" {
  uses = "cedrickring/golang-action@1.2.0"
  needs = ["Filter branch master"]
}

action "GCP Authenticate" {
  needs = ["Filter branch master"]
  uses = "actions/gcloud/auth@master"
  secrets = ["GCLOUD_AUTH"]
}

action "GCP Deploy" {
  needs = ["GCP Authenticate", "Build"]
  uses = "actions/gcloud/cli@master"
  args = "app deploy --quiet app.yaml"
  secrets = ["CLOUDSDK_CORE_PROJECT"]
}
