#!/bin/bash

commands='
agentPool
agentToken
applies
configurationVersion
configure
costEstimate
notificationConfiguration
oAuthClient
oAuthToken
oauth
organizationMembership
organizationToken
organization
planExport
plan
policy
policyCheck
policySetParameter
policySet
registryModule
root
runTrigger
run
sshKey
stateVersionOutput
stateVersion
teamAccesse
teamMember
teamToken
teams
userToken
user
variable
version
workspace
'

for command in ${commands}
do
    if [[ ${command} == 'oAuthClient' ]]; then
        echo "Working on command ${command} now"
        cp command_template.tmpl "cobra/cmd/${command}.go"
        sed "s,COMMAND_LC_,${command},g"  "cobra/cmd/${command}.go" > "cobra/cmd/${command}.go.1"
        command_uc="$(tr '[:lower:]' '[:upper:]' <<< ${command:0:1})${command:1}"
        sed "s,COMMAND_UC_,${command_uc},g"  "cobra/cmd/${command}.go.1" > "cobra/cmd/${command}.go"
        rm -f "cobra/cmd/${command}.go.1"

        # controller
        cp command_controller_template.tmpl "cobra/controller/${command}.go"
        sed "s,COMMAND_LC_,${command},g"  "cobra/controller/${command}.go" > "cobra/controller/${command}.go.1"
        sed "s,COMMAND_UC_,${command_uc},g"  "cobra/controller/${command}.go.1" > "cobra/controller/${command}.go"
        rm -f "cobra/controller/${command}.go.1"

        # aid
        cp command_aid_template.tmpl "cobra/aid/${command}.go"
        sed "s,COMMAND_LC_,${command},g"  "cobra/aid/${command}.go" > "cobra/aid/${command}.go.1"
        sed "s,COMMAND_UC_,${command_uc},g"  "cobra/aid/${command}.go.1" > "cobra/aid/${command}.go"
        rm -f "cobra/aid/${command}.go.1"

        # manual
        cp command_manual_template.tmpl "box/resources/manual/${command}.yaml"
        sed "s,COMMAND_LC_,${command},g"  "box/resources/manual/${command}.yaml" > "box/resources/manual/${command}.yaml.1"
        sed "s,COMMAND_UC_,${command_uc},g"  "box/resources/manual/${command}.yaml.1" > "box/resources/manual/${command}.yaml"
        rm -f "box/resources/manual/${command}.yaml.1"

        # test
        cp command_test_template.tmpl "tests/command_${command}_test.go"
        sed "s,COMMAND_LC_,${command},g"  "tests/command_${command}_test.go" > "tests/command_${command}_test.go.1"
        sed "s,COMMAND_UC_,${command_uc},g"  "tests/command_${command}_test.go.1" > "tests/command_${command}_test.go"
        rm -f "tests/command_${command}_test.go.1"

        exit 0
    fi
done