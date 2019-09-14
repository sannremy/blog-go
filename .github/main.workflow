workflow "Deploy" {
  on = "push"
  resolves = ["Filter branch master", "Yarn Install", "Yarn Build", "Go Build", "GCP Authenticate", "GCP Deploy"]
}

action "Filter branch master" {
  uses = "actions/bin/filter@master"
  args = "branch master"
}

action "Yarn Install" {
  uses = "Borales/actions-yarn@master"
  args = "install"
  needs = ["Filter branch master"]
}

action "Yarn Build" {
  uses = "Borales/actions-yarn@master"
  args = "build"
  needs = ["Yarn Install"]
}

action "Go Build" {
  uses = "cedrickring/golang-action/go1.12@1.3.0"
  needs = ["Filter branch master"]
}

action "GCP Authenticate" {
  needs = ["Filter branch master"]
  uses = "actions/gcloud/auth@master"
  secrets = ["GCLOUD_AUTH"]
}

action "GCP Deploy" {
  needs = ["GCP Authenticate", "Go Build"]
  uses = "actions/gcloud/cli@master"
  args = "app deploy --quiet app.yaml"
  secrets = ["CLOUDSDK_CORE_PROJECT"]
}
