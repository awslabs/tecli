# Terminalizer

```
tecli configure create --profile=foo-org

tecli workspace create --profile=foo-org --organization=foo-org --name=bar

tecli workspace create --vcs-repo-oauth-token-id=ot-m4nuCHotgiG4JqgZ --vcs-repo-identifier=valter-silva-au/terraform-dummy --organization=foo-org --name=bar

tecli run create --workspace-id=ws-zpwJtDTH1m4QT5b2
tecli run read --id=run-PHt4Ue7pcmyyMyuR | jq .Plan.ID
tecli plan logs --id=plan-kPNiuSxJbZb4zFVb

tecli run apply --id=run-PHt4Ue7pcmyyMyuR
tecli run read --id=run-PHt4Ue7pcmyyMyuR | jq .Apply.ID
tecli apply logs --id=apply-mgyeDqkRtbuf94Jc

tecli workspace delete --profile=foo-org --organization=foo-org --name=bar
```
