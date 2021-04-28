# Step 4 - Manually Deploy Into IBM Cloud Foundry

The first part of this step is to manually deploy into IBM Cloud Foundry to ensure it works and help you understand what is happening when it is automated.

Open up a new terminal session and login to the [IBM Cloud CLI](<LINK-HERE>). If you have not already installed it, please do so.

To login use the command

`$ ibmcloud login`

If you have a federated ID, you will need to use

`$ ibmcloud login --sso`

List regions and find your closest region

`$ ibmcloud regions`

Target your chosen region

`$ ibmcloud target -r <your-region>` (use the Name and NOT the Display name - Example: `eu-gb`)

Target the Default resource group

`$ ibmcloud target -g Default`

Finally, target IBM Cloud Foundry

`$ ibmcloud target --cf`

Now the setup is complete, lets write some YAML!

Open the `manifest.yml` file. This file will contain all the configuration IBM Cloud needs to deploy your application. Copy the code snippet below and paste it into your file. **Make sure to change the parts unique to your deployment**

```yaml
applications:
  - name: <unique-name-of-your-application>
    routes:
      - route: <unique-name-of-your-application>.<your-chosen-domain>
    memory: 128M
    buildpack: go_buildpack
```

Broken down, the name of your application will be what it is called in IBM Cloud Foundry. This needs to be a unique name.

The route is made up of two parts, your chosen application name and your chosen domain. To find a list of domains, use the command `ibmcloud cf domains`.

The memory is how much memory you wish your instance to have in IBM Cloud. This application doesn't need much so I have set it to 128M.

The buildpack states the desired buildpack you wish your application to use. This workshop uses a Go application so the buildpack mirrors this. Likewise, if you are using a different language, you will need to specify the correct buildpack in this config file.

An example completed `manifest.yml`:

```yaml
applications:
  - name: github-actions-demo
    routes:
      - route: github-actions-demo.eu-gb.mybluemix.net
    memory: 128M
    buildpack: go_buildpack
```

All that is left to do is push this into IBM Cloud Foundry!

`ibmcloud cf push -f ./manifest.yml`

One the command has run to completion, you should see some output similar to:

```bash
Waiting for app to start...

name:              github-actions-demo
requested state:   started
routes:            github-actions-demo.eu-gb.mybluemix.net
last uploaded:     Tue 27 Apr 17:07:36 UTC 2021
stack:             cflinuxfs3
buildpacks:        go

type:            web
instances:       1/1
memory usage:    128M
start command:   ./bin/github-actions-demo
     state     since                  cpu    memory          disk         details
#0   running   2021-04-27T17:07:50Z   0.5%   16.7M of 128M   7.2M of 1G   
```

Copy the route of your application and past it into a browser search bar. You will see your application running. Example: `github-actions-demo.eu-gb.mybluemix.net`

Once it is all working, commit these changes to your forked repository in GitHub.

Final step is to automate this! 