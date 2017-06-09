node('k8s') {

    sh 'env | sort'
    sh 'pwd'

    checkout scm

    def inputFile = readFile('./cicd.json')
    def config = new groovy.json.JsonSlurperClassic().parseText(inputFile)
    println "pipeline config ==> ${config}"

}