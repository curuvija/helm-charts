pipeline {
    agent { label 'k8s-helm-pod' }

    options{
        buildDiscarder(logRotator(numToKeepStr: '5', daysToKeepStr: '5'))
    }

    environment {
        GITHUB_TOKEN = credentials('github_curuvija_jcasc')
        HELM_CHART_PATH = 'oracledb-exporter/'
        GITHUB_REPO_NAME = 'oracledb-exporter'
    }


    stages {
        // TODO: kube-linter -> https://github.com/stackrox/kube-linter (kube-score, kubeconform, kubeeval, datree, kics -> https://kics.io/index.html
        // TODO: check also https://analysis-tools.dev/tag/kubernetes
        // TODO: kube-hunter -> https://aquasecurity.github.io/kube-hunter/
        stage('Lint') {
            steps {
                container('helm') {
                    sh 'helm lint ${HELM_CHART_PATH}'
                }
            }
        }
        // TODO: polaris https://github.com/FairwindsOps/polaris
        stage('Scan for security issues') {
            steps {
                sh 'echo something'
            }
        }
        // TODO: Terratest (you need to install prometheus crds)
        stage('Test') {
            steps {
                container('helm') {
                    sh 'cd tests && go test ./...'
                }
            }
        }
        stage('Template') {
            steps {
                container('helm') {
                    sh 'helm template ${HELM_CHART_PATH}'
                }
            }
        }
        stage('Dry run') {
            steps {
                sh 'echo something'
            }
        }
        stage('Generate docs') {
            steps {
                container('helm') {
                    sh 'helm-docs ${HELM_CHART_PATH} && cp ${HELM_CHART_PATH}README.md .'
                }
            }
        }
        // TODO: git-chglog - https://github.com/git-chglog/git-chglog
        stage('Changelog') {
            steps {
                sh 'echo something'
            }
        }
        // TODO: generate provenance file -> https://helm.sh/docs/topics/provenance/
        stage('Package Helm chart') {
            steps {
                container('helm') {
                    sh 'cr package ${HELM_CHART_PATH}'
                }
            }
        }
        stage('Upload Helm chart to releases') {
            when {
                allOf {
                    expression {
                        env.BRANCH_NAME == 'master'
                    }
                }
            }
            steps {
                container('helm') {
                    sh 'cr upload -o curuvija --git-repo ${GITHUB_REPO_NAME} --package-path .cr-release-packages/ --token ${GITHUB_TOKEN_PSW}'
                }
            }
        }
        stage('Publish') {
            when {
                allOf {
                    expression {
                        env.BRANCH_NAME == 'master'
                    }
                }
            }
            steps {
                sh 'echo something'
            }
        }
        stage('Tag') {
            when {
                allOf {
                    expression {
                        env.BRANCH_NAME == 'master'
                    }
                }
            }
            steps {
                sh 'echo something'
            }
        }
    }
}
