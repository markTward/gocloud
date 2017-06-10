node('k8s') {

    stage("preparation") {
        sh 'env | sort'
        sh 'pwd'

        checkout scm

        def config = readYaml file: './cicd.yaml'
        println "Config CICD ==> ${config}"
    }

    stage("scope test") {
        println "Config CICD ==> ${config}"
    }

}