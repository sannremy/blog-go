workflow "Deploy" {
  on = "push"
  resolves = ["Branch master", "GCP Authenticate", "GCP Deploy"]
}

action "Branch master" {
  uses = "actions/bin/filter@master"
  args = "branch master"
}

action "GCP Authenticate" {
  needs = ["Branch master"]
  uses = "actions/gcloud/auth@master"
  secrets = ["GCLOUD_AUTH"]
}

action "GCP Deploy" {
  needs = ["GCP Authenticate"]
  uses = "actions/gcloud/cli@master"
  args = "app deploy --quiet"
  secrets = ["CLOUDSDK_CORE_PROJECT"]
}
