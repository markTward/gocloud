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
        def config = readYaml file: './cicd.yaml'

        stage('setup') {
            println "Config CICD ==> ${config}"
        }

        stage ('test') {
            container('golang') {
                sh 'go version'
            }
        }
    }
}