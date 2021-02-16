# Top Commands

All the following commands require [TEAM API TOKEN](https://www.terraform.io/docs/cloud/users-teams-organizations/api-tokens.html#team-api-tokens). 
You can run `tecli configure create` to configure TECLI options. Alternatively, you can export [environment varibles](https://github.com/awslabs/tecli/wiki/Environment-Variables).

To list all workspaces part of an organization:
```
tecli workspace list -o=${TFC_ORGANIZATION} -p=${PROFILE}
```

To find a workspace by name (instead of listing all workspaces and look for its ID):
```
tecli workspace find-by-name --organization=${TFC_ORGANIZATION} --name=${TFC_WORKSPACE_NAME}
```

To create a workspace and allow destroy plans:
```
tecli workspace create --organization=${TFC_ORGANIZATION} --name=${TFC_WORKSPACE_NAME} --allow-destroy-plan=true
```

To create a plan (if you want to upload your code to Terraform Cloud):

```
tecli configuration-version create --workspace-id=${W_ID}
tecli configuration-version upload --url=${CV_UPLOAD_URL} --path=./
tecli run create --workspace-id=${W_ID} --comment="${COMMENT}" 
```

To check the staus of a run:
```
tecli run read --id=${RUN_ID}
```

You combine some `BASH` scripting and check if your plan has finished:

``` 
while true; do STATUS=$(tecli run read --id=${RUN_ID} | jq -r ".Status"); if [ "${STATUS}" != "pending" ]; then break; else echo "RUN STATUS:${STATUS}, IF 'pending' TRY DISCARD PREVIOUS PLANS. SLEEP 5 seconds" && sleep 5; fi; done
```

To display the logs of a plan:
```
tecli plan logs --id=${PLAN_ID}
```

To leave a comment on a plan:

```
tecli run create --workspace-id=${W_ID} --comment="${COMMENT}" --is-destroy=true
```

To discard a run:
```
tecli run discard --id=${RUN_ID}
```

To discard all runs:
```
tecli run discard-all --workspace-id=${W_ID}
```

To apply a plan:
```
tecli run apply --id=${RUN_ID} --comment="${COMMENT}"
```

To display the apply logs:
```
tecli apply logs --id=${APPLY_ID}
```