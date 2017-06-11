podTemplate(label: 'jenkins-pipeline', containers: [
    containerTemplate(name: 'jnlp', image: 'jenkinsci/jnlp-slave:2.62', args: '${computer.jnlpmac} ${computer.name}', workingDir: '/home/jenkins', resourceRequestCpu: '200m', resourceLimitCpu: '200m', resourceRequestMemory: '256Mi', resourceLimitMemory: '256Mi'),
    containerTemplate(name: 'docker', image: 'docker:1.11.1', command: 'cat', ttyEnabled: true),
    containerTemplate(name: 'golang', image: 'golang:1.8.1', command: 'cat', ttyEnabled: true),
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
                sh "docker push marktward/${config.app.name}:${gitCommit}"
            }
        }

    }
}