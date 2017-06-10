node('k8s') {

    checkout scm

    def config = readYaml file: './cicd.yaml'

    stage("preparation") {
        sh 'env | sort'
        sh 'pwd'
        sh 'ls -la'
        println "Config CICD ==> ${config}"
        println "CICD Tools ==> ${config.provider.cicd}"
    }

}