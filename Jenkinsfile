podTemplate(label: 'jenkins-pipeline', containers: [
    containerTemplate(name: 'docker', image: 'docker:1.11.1', command: 'cat', ttyEnabled: true),
    containerTemplate(name: 'golang', image: 'golang:1.8.1', command: 'cat', ttyEnabled: true),
    containerTemplate(name: 'cicd', image: 'marktward/gocloud-cicd:minikube', command: 'cat', ttyEnabled: true),
],
volumes:[
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
]){

    node('jenkins-pipeline') {
        checkout scm

        def gitCommit = sh(returnStdout: true, script: 'git rev-parse HEAD').trim().take(7)
        def config = readYaml file: './cicd.yaml'

        stage('setup') {
            sh 'pwd'
            sh 'ls -la'
            sh 'env | sort'
            dir('gocloud-cicd') {
                git url: 'https://github.com/markTward/gocloud-cicd.git', branch: 'jenkins'
            }
            sh 'ls -la'
        }

        stage ('test') {
            container('golang') {
                sh 'go env'
                sh 'ls -la'
                sh 'go get -d -t -v -race ./...'
                sh 'go test -v ./...'
            }
        }

        stage ('build') {
            container('docker') {
                sh 'docker version'
                println "build image: ${config.app.name}:${gitCommit}"
                sh "docker build -t ${config.app.name}:${gitCommit} -f Dockerfile ."
                sh "docker tag ${config.app.name}:${gitCommit} marktward/${config.app.name}:${gitCommit}"
                sh "docker images"

                //withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'dockerHub',
                //            usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
                //  sh "docker login -u ${env.USERNAME} -p ${env.PASSWORD}"
                //}
                //sh "docker push marktward/${config.app.name}:${gitCommit}"
            }
        }

        stage ('deploy') {
            container('cicd') {
                sh 'gocloud-cicd deploy --help'
                sh 'docker version'
                sh 'docker images'
                sh 'which gcloud && gcloud version'
                sh 'which kubectl && kubectl version'
            }
        }
    }
}