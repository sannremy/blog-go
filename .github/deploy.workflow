workflow "New workflow" {
  on = "push"
  resolves = ["GCP Authenticate"]
}

action "GCP Authenticate" {
  uses = "actions/gcloud/auth@master"
  secrets = ["GCLOUD_AUTH"]
}

action "GCP Deploy" {
  needs = ["GCP Authenticate"]
  uses = "actions/gcloud/cli@master"
  args = "app deploy"
}
