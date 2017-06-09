node('k8s') {

    sh 'env | sort'
    sh 'pwd'

    checkout scm

    sh 'ls -la'

    org.yaml.snakeyaml.Yaml parser = new org.yaml.snakeyaml.Yaml()
    def config = parser.load(("./cicd.yaml" as File).text)

    println "pipeline config ==> ${config}"

}