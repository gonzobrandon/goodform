#!/usr/bin/env bash

NO_COLOR="\033[0m"
OK_COLOR="\033[32;01m"
ERROR_COLOR="\033[31;01m"
WARN_COLOR="\033[33;01m"

target_branch=master
if (( "$#" > 0 ))
then
	target_branch=$1
fi
echo "Target branch: $target_branch"

msg() {
  echo
  echo -en "$1 * $2 $NO_COLOR"
  echo
}

info() {
  echo -en "$NO_COLOR * $1 $NO_COLOR"
}

warn() {
  msg "$WARN_COLOR" "$1"
}

error() {
  msg "$ERROR_COLOR" "$1"
}

ok() {
  msg "$OK_COLOR" "$1"
}

ready() {
  echo -e "$OK_COLOR OK!"
}

check_for_multi_buildpack() {
  info "Checking BUILDPACK_URL to allow multiple buildpack configs"
  multi_buildpack_url="https://github.com/ddollar/heroku-buildpack-multi.git"
  buildpack_url=$(heroku config:get BUILDPACK_URL)

  if [[ $buildpack_url != $multi_buildpack_url ]]; then
    heroku config:add BUILDPACK_URL=$multi_buildpack_url > /dev/null
    warn "BUILDPACK_URL updated! to $multi_buildpack_url"
  else
    ready
  fi
}

check_for_netrc_buildpack() {
  info "Checking netrc buildpack to allow private fetching"

  netrc_buildpack_url="https://github.com/timshadel/heroku-buildpack-github-netrc"
  buildpacks_file=".buildpacks"

  grep "$netrc_buildpack_url" $buildpacks_file &> /dev/null

  case $? in
    0)
      ready
      ;;

    2)
      error "Missing $buildpacks_file file"
      exit 1
      ;;

    *)
      echo -e "$netrc_buildpack_url\n$(cat $buildpacks_file)" > "$buildpacks_file"
      warn "Added netrc buildpack to $buildpacks_file"
  esac
}

check_for_user_env_compile() {
  info "Checking user-env-compile in heroku labs"
  heroku labs | grep user-env-compile | grep + > /dev/null

  if [ $? -ne 0 ]; then
    heroku labs:enable user-env-compile > /dev/null
    warn "user-env-compile added to heroku labs"
  else
    ready
  fi
}


check_for_github_token() {
  info "Checking GITHUB_AUTH_TOKEN to allow access to private repos"
  github_auth_token=$(heroku config:get GITHUB_AUTH_TOKEN)

  if [[ $github_auth_token == "" ]]; then
    error "No GitHub token in ENV"
    echo -n "GitHub token: "
    read token
    heroku config:add GITHUB_AUTH_TOKEN=$token > /dev/null
    warn "Token added to the ENV"
  else
    ready
  fi
}

push_to_github() {
  ok "Pushing to GitHub"
  git push origin ${target_branch}
}

deploy_application() {
  ok "Deploying application"
  echo "Running command: git push heroku ${target_branch}:master"
  git push heroku ${target_branch}:master
}

check_for_multi_buildpack
check_for_netrc_buildpack
check_for_user_env_compile
check_for_github_token

push_to_github
deploy_application
