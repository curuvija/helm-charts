
## Todo

[ ] Finish configuration for github pages
[ ] Finish query-exporter helm chart and publish it
[ ] Check how to add custom domain
[ ] Check best practices for writing Helm charts https://codersociety.com/blog/articles/helm-best-practices 

## Checklist:

1. Run helm lint to check if chart is well-formed
2. Use Polaris for static code analysis
3. Create unit and integration tests
4. Test helm chart deployment using Helm client
5. Use helm docs https://github.com/norwoodj/helm-docs to document values
6. Add NOTES.txt in templates folder
7. Publish chart using Helm chart releaser

Check what can you do with:

* https://github.com/marketplace/actions/helm-chart-testing - A GitHub Action for installing the helm/chart-testing CLI tool.
* https://github.com/marketplace/actions/kind-cluster - A GitHub Action for Kubernetes IN Docker - local clusters for testing Kubernetes using kubernetes-sigs/kind
* https://helm.sh/docs/howto/chart_releaser_action/

## Reference

* https://medium.com/@mattiaperi/create-a-public-helm-chart-repository-with-github-pages-49b180dbb417
* https://tech.paulcz.net/blog/creating-a-helm-chart-monorepo-part-1/
* https://helm.sh/docs/howto/chart_releaser_action/