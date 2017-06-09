node('k8s') {

    stage 'preparation'

    // output environment
    sh 'env | sort'
    sh 'pwd'

    // git checkout
    checkout scm
    sh 'ls -la'

    org.yaml.snakeyaml.Yaml parser = new org.yaml.snakeyaml.Yaml()
    def config = parser.load(("./cicd.yaml" as File).text)

    println "after yaml load"
    println config

}