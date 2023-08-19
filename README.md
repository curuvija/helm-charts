Helm repository where I work on my helm charts.

## Publishing helm chart - Checklist

List of tasks to do when creating/publishing Helm charts:

* Create helm chart from ``starter`` chart (helm create mychart --starter C:\Users\micu2112\Desktop\my\temp2\helm-charts\charts\starter\) - it has to be an absoluth path
* Run kube-linter to check if chart is well-formed (``kube-linter lint charts\query-exporter``) and clear all errors you can
* add artifacthub annotations https://artifacthub.io/docs/topics/annotations/helm/ (check also https://artifacthub.io/docs/topics/repositories/helm-charts/)
* Create unit and integration tests
* Test helm chart deployment using Helm client
* Use helm docs https://github.com/norwoodj/helm-docs to document values (use pre commit https://github.com/norwoodj/helm-docs#pre-commit-hook)
* Add NOTES.txt in templates folder
* Package helm chart using ``cr`` command for example ``cr package charts/query-exporter/`` (package will be stored in ``.cr-release-packages/`` folder) (TODO: sign the chart)
* Now upload release for example ``cr upload -o curuvija --git-repo helm-charts --package-path .cr-release-packages/ --token <token here>``
* Checkout to ``gh-pages`` branch
* Update ``index.yaml``
* Create index for your Helm chart repo ``cr index --index-path ./index.yaml --package-path .cr-release-packages/ --owner curuvija --git-repo helm-charts --charts-repo https://curuvija.github.io/helm-charts/``
* Push changes to ``gh-pages``
* Create release notes manually following this example https://gist.github.com/juampynr/4c18214a8eb554084e21d6e288a18a2c (use CHANGELOG.md file)






