@NonCPS
def cancelPreviousBuilds() {
  String jobname = env.JOB_NAME
  int currentBuildNum = env.BUILD_NUMBER.toInteger()

  def job = Jenkins.instance.getItemByFullName(jobname)
//   print('jobname: ' + jobname)
  for (build in job.builds) {
    //   print('build.getNumber(): ' + build.getNumber())

    if (build.isBuilding() && currentBuildNum > build.getNumber().toInteger()) {
      build.doStop();
      echo "Build ${build.toString()} cancelled"
    }
  }
}

pipeline {

    environment {
        registry = "197272534240.dkr.ecr.us-east-1.amazonaws.com"
        registryCredential = "ecr:us-east-1:aws_vert"
        dockerImageName = ""
    }

    agent {
        docker {
            image "akaytatsu/cibuilder:latest"
        }
    }

    stages {

        stage('Init') {
            steps {
                cancelPreviousBuilds()
            }
        }

        stage('Code Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Images') {
            steps {
                script {
                    sh 'cp -f src/.env.sample src/.env'
                    sh 'docker-compose -f docker-compose.yml down'
                    sh 'docker-compose -f docker-compose.yml build'
                    sh 'docker-compose -f docker-compose.yml up -d --no-build'
                }
            }
        }

        stage('stop containers') {
            steps {
                script {
                    sh 'docker-compose -f docker-compose.yml down'
                }
            }
        }

        stage('build Container Register Production') {
            when {
                expression {
                    return env.GIT_BRANCH == 'master'
                }
            }

            steps {
                script {
                    docker.withRegistry("https://$registry", registryCredential) {
                        dockerImageName = "skyswitch-prd"
                        dockerImage = docker.build(dockerImageName, "./src")
                        dockerImage.push("$BUILD_NUMBER")
                        dockerImage.push("latest")
                    }
                }

                script{
                    sh "docker rmi $registry/$dockerImageName:$BUILD_NUMBER"
                    sh "docker rmi $registry/$dockerImageName:latest"
                }
            }
        }

        stage('Deploy to Production Environment') {
            when {
                expression {
                    return env.GIT_BRANCH == 'master'
                }
            }

            steps {
                script {
                    withCredentials([string(credentialsId: "ARGOCD_SERVER", variable: 'ARGOCD_SERVER')]) {
                        withCredentials([string(credentialsId: "argocd-production", variable: 'ARGOCD_AUTH_TOKEN')]) {
                            sh "argocd --grpc-web app actions run skyswitch-prd restart --kind Deployment --all"
                        }
                    }
                }
            }
        }
    }

    post {
        always {
            echo "Stop Docker image"
            script{
                sh 'docker-compose -f docker-compose.yml down'
            }
        }

        success {
            echo "Notify success"
            script {
                sh 'docker-compose -f docker-compose.yml down'
            }
        }

        failure {
            echo "Notify failure"
            script {
                sh 'docker-compose -f docker-compose.yml down'
            }
        }

        aborted {
            echo "Notify failure"
            script {
                sh 'docker-compose -f docker-compose.yml down'
            }
        }

        unsuccessful {
            echo "Notify failure"
            script {
                sh 'docker-compose -f docker-compose.yml down'
            }
        }
    }

}
