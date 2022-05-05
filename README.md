On this page you can find info when working with Helm charts in this repo.

# Table of contents
- [Table of contents](#table-of-contents)
  - [Development tools](#development-tools)
  - [Publishing helm chart - Checklist](#publishing-helm-chart---checklist)
  - [Docs generation - Checklist](#docs-generation---checklist)
  - [Quality checks](#quality-checks)
    - [Static code analysis](#static-code-analysis)
  - [Kubeval](#kubeval)
    - [Kube linter](#kube-linter)
    - [Datree](#datree)
    - [Checkov](#checkov)
    - [Polaris](#polaris)
    - [CT](#ct)
- [Security analysis](#security-analysis)
- [Testing](#testing)
  - [Terratest](#terratest)
  - [CT - Chart Testing](#ct---chart-testing)
  - [Helm unittest](#helm-unittest)
  - [Kubetest](#kubetest)
  - [Todo (Improvements)](#todo-improvements)
  - [Reference](#reference)
## Development tools

List of tools/software you need on your computer to work with Helm:

* Git
* Helm
* Chart releaser - https://github.com/helm/chart-releaser/releases
* helm-docs
* VS Code
* Ct
* Chart releaser
* Polaris
* Golang
* Terratest
* git-chglog - https://github.com/git-chglog/git-chglog
## Publishing helm chart - Checklist

List of tasks to do when creating/publishing Helm charts:

1. Create helm chart from ``starter`` chart (helm create mychart --starter C:\Users\micu2112\Desktop\my\temp2\helm-charts\charts\starter\) - it has to be an absoluth path
2. Run kube-linter to check if chart is well-formed (``kube-linter lint charts\query-exporter``) and clear all errors you can
3. Create unit and integration tests
4. Test helm chart deployment using Helm client
5. Use helm docs https://github.com/norwoodj/helm-docs to document values (use pre commit https://github.com/norwoodj/helm-docs#pre-commit-hook)
6. Add NOTES.txt in templates folder
7. Package helm chart using ``cr`` command for example ``cr package charts/query-exporter/`` (package will be stored in ``.cr-release-packages/`` folder)
8.  Now upload release for example ``cr upload -o curuvija --git-repo helm-charts --package-path .cr-release-packages/ --token <token here>``
9.  Checkout to ``gh-pages`` branch
10. Delete ``index.yaml``
11. Create index for your Helm chart repo ``cr index --index-path ./index.yaml --package-path .cr-release-packages/ --owner curuvija --git-repo helm-charts --charts-repo https://curuvija.github.io/helm-charts/``
12. Push changes to ``gh-pages``
13. Create release notes manually following this example https://gist.github.com/juampynr/4c18214a8eb554084e21d6e288a18a2c (use CHANGELOG.md file)

## Docs generation - Checklist

For docs generation I use helm-docs https://github.com/norwoodj/helm-docs. This is the list of steps to produce README.md file
for the Helm chart:

1. Populate Chart.yaml fields https://helm.sh/docs/topics/charts/
2. You need to download 'helm-docs' executable from this page https://github.com/norwoodj/helm-docs/releases.
3. Create README.md.gotmpl and put it inside Helm chart (ignore this file with .helmignore)
4. Use examples found here https://github.com/norwoodj/helm-docs/tree/master/example-charts to make your template
5. Run ``helm-docs -c <helm chart folder path here>`` for example ``helm-docs -c query-exporter/`` to generate docs
6. Configure pre-commit hook to generate docs on commit https://github.com/norwoodj/helm-docs#pre-commit-hook

## Quality checks

### Static code analysis

Analyse the code and fix common issues. Use multiple different tools to get the best results and avoid to skip something.

## Kubeval

Find project here: https://kubeval.instrumenta.dev/

* case1: kubeval somefile.yaml

### Kube linter

* case1 - scan helm chart:  Run kube-linter to check if chart is well-formed (``kube-linter lint charts\query-exporter``) and clear all errors you can

### Datree

* case1 - scan templates: output templates ``helm template query-exporter .\charts\query-exporter\ > .\temp\query-exporter-templates.yaml`` and run ``datree test .\temp\query-exporter-templates.yaml`` and analyse results

### Checkov

* case1 - scan templates: output templates ``helm template query-exporter .\charts\query-exporter\ > .\temp\query-exporter-templates.yaml`` and run ``checkov -d .\temp\`` (evaluate results)

### Polaris

* case1 - scan templates: output templates ``helm template query-exporter .\charts\query-exporter\ > .\temp\query-exporter-templates.yaml`` and run ``polaris audit --audit-path .\temp\oracledb-exporter-podmonitor.yaml``
* case2 - scan helm chart: polaris audit --helm-chart .\charts\query-exporter\ --only-show-failed-tests
* case2 - display dashboard with results: polaris dashboard --audit-path .\temp\oracledb-exporter-podmonitor.yaml --port 9999

Then open your browser and inspect results:

``http://localhost:9999/``

### CT

* case1 - scan all charts: ``ct lint --chart-dirs .\charts\ --all --chart-yaml-schema C:\tools\etc\chart_schema.yaml --lint-conf C:\tools\etc\lintconf.yaml``
* case1 - scan specific charts: ``ct lint --chart-dirs .\charts\ --charts .\charts\query-exporter\ --chart-yaml-schema C:\tools\etc\chart_schema.yaml --lint-conf C:\tools\etc\lintconf.yaml``

Find more about it here https://github.com/helm/chart-testing.
# Security analysis

* kube-bench: 
* trivy: 

# Testing

Write tests to validate helm template output and integration tests when you deploy it onto cluster (you can probably combine this with kube-bench scan and trivy)
## Terratest

* terratest: https://terratest.gruntwork.io/docs/getting-started/quick-start/
## CT - Chart Testing

* ct: ``ct install --charts .\charts\query-exporter\``

Find more info here https://github.com/helm/chart-testing.

## Helm unittest

https://github.com/quintush/helm-unittest

## Kubetest

https://kubetest.readthedocs.io/en/latest/index.html

## Todo (Improvements)

- [x] Finish configuration for github pages
- [x] Finish query-exporter helm chart and publish it
- [ ] Check how to add custom domain to Gihub pages
- [ ] Check best practices for writing Helm charts https://codersociety.com/blog/articles/helm-best-practices and plan for changes
- [ ] Check checklist for publishing releases here Check this one too next time https://helm.sh/docs/community/release_checklist/ and plan for changes
- [ ] https://github.com/marketplace/actions/helm-chart-testing - Implement GitHub Action for installing the helm/chart-testing CLI tool.
- [ ] https://github.com/marketplace/actions/kind-cluster - A GitHub Action for Kubernetes IN Docker - local clusters for testing Kubernetes using kubernetes-sigs/kind
- [ ] Create workflow to use chart releaser https://helm.sh/docs/howto/chart_releaser_action/
- [ ] Implement automatic helm chart scan on commit
- [ ] Create script to rebuild complete development environment (download tools, install docker, kind...)
- [ ] Find way to run OracleDB to run integration tests for oracledb-exporter

## Reference

* https://medium.com/@mattiaperi/create-a-public-helm-chart-repository-with-github-pages-49b180dbb417
* https://tech.paulcz.net/blog/creating-a-helm-chart-monorepo-part-1/
* https://helm.sh/docs/howto/chart_releaser_action/

