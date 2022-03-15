On this page you can find info when working with Helm charts in this repo.

# Table of contents
- [Table of contents](#table-of-contents)
  - [Development tools](#development-tools)
  - [Publishing helm chart - Checklist](#publishing-helm-chart---checklist)
  - [Docs generation - Checklist](#docs-generation---checklist)
  - [Todo (Improvements)](#todo-improvements)
  - [Reference](#reference)
## Development tools

List of tools/software you need on your computer to work with Helm:

* Git
* Helm
* Chart releaser - https://github.com/helm/chart-releaser/releases
* helm-docs
* VS Code
* Chart releaser
* Polaris
* Golang
* Terratest
## Publishing helm chart - Checklist

List of tasks to do when publishing Helm charts:

1. Run helm lint to check if chart is well-formed
2. Use Polaris for static code analysis
3. Create unit and integration tests
4. Test helm chart deployment using Helm client
5. Use helm docs https://github.com/norwoodj/helm-docs to document values (use pre commit https://github.com/norwoodj/helm-docs#pre-commit-hook)
6. Add NOTES.txt in templates folder
7. Checkout to ``gh-pages`` branch
8. Package helm chart using ``cr`` command for example ``cr package charts/query-exporter/`` (package will be stored in ``.cr-release-packages/`` folder)
9.  Now upload release for example ``cr upload -o curuvija --git-repo helm-charts --package-path .cr-release-packages/ --token --token <token here>``
10. Create/Update index for your Helm chart repo ``cr index --index-path ./index.yaml --package-path .cr-release-packages/ --owner curuvija --git-repo helm-charts --charts-repo https://curuvija.github.io/helm-charts/``
11. Push changes to ``gh-pages``

## Docs generation - Checklist

For docs generation I use helm-docs https://github.com/norwoodj/helm-docs. This is the list of steps to produce README.md file
for the Helm chart:

1. Populate Chart.yaml fields https://helm.sh/docs/topics/charts/
2. You need to download 'helm-docs' executable from this page https://github.com/norwoodj/helm-docs/releases.
3. Create README.md.gotmpl and put it inside Helm chart (ignore this file with .helmignore)
4. Use examples found here https://github.com/norwoodj/helm-docs/tree/master/example-charts to make your template
5. Run ``helm-docs -c <helm chart folder path here>`` for example ``helm-docs -c query-exporter/`` to generate docs
6. Configure pre-commit hook to generate docs on commit https://github.com/norwoodj/helm-docs#pre-commit-hook

## Todo (Improvements)

- [x] Finish configuration for github pages
- [x] Finish query-exporter helm chart and publish it
- [ ] Check how to add custom domain to Gihub pages
- [ ] Check best practices for writing Helm charts https://codersociety.com/blog/articles/helm-best-practices and plan for changes
- [ ] Check checklist for publishing releases here Check this one too next time https://helm.sh/docs/community/release_checklist/ and plan for changes
- [ ] https://github.com/marketplace/actions/helm-chart-testing - Implement GitHub Action for installing the helm/chart-testing CLI tool.
- [ ] https://github.com/marketplace/actions/kind-cluster - A GitHub Action for Kubernetes IN Docker - local clusters for testing Kubernetes using kubernetes-sigs/kind
- [ ] Create workflow to use chart releaser https://helm.sh/docs/howto/chart_releaser_action/

## Reference

* https://medium.com/@mattiaperi/create-a-public-helm-chart-repository-with-github-pages-49b180dbb417
* https://tech.paulcz.net/blog/creating-a-helm-chart-monorepo-part-1/
* https://helm.sh/docs/howto/chart_releaser_action/

