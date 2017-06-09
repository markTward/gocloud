node('k8s') {

    stage 'preparation'

    // output environment
    sh 'env | sort'
    sh 'pwd'

    // git checkout
    checkout scm
    sh 'ls -la'

    @Grab('org.yaml:snakeyaml:1.17')

    import org.yaml.snakeyaml.Yaml

    Yaml parser = new Yaml()
    List example = parser.load(("./cicd.yaml" as File).text)

    example.each{println it.subject}

}