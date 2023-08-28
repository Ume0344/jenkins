**[Tutorial](https://www.youtube.com/watch?v=FX322RVNGj4&t=104s)**

**Agile Model**

- Project is divided into several sprints (2 weeks). In each sprint, we have product backlogs, these backlogs are reviewed and assigned to the development team in spint planning. In  sprint, devlopement team plan, code, test and deliver the product features (daily scrum/standup meeting). Then after two weeks, these features are sent to client for review to get their feedback. 
- Client requirements are understood better through constant feedback and product are delivered faster.
- Disadvantages - Products are tested on only developer test environment not on production environment. And, developers and operations work separately. When product fails in production, operation team sends the product back to devlopment team to find the issue which causes delay in projects. Here comes, DevOps.

**DevOps Phases**

- Plan
- Code (Git)
- Build
- Test
- Integrate (Jenkins)
- Deploy
- Operate (Ansible)
- Monitor (Prometheus)

**DevOps Process**

*Continuous Integration*

- Plan -> Code -> Build -> Test

- If some bugs occurs, it is commuincated to devlopment team.

*Continuous Deployment*

- Deployment -> Operate (to ensure that the applications are available, responsive, and performing well for end users) -> Monitor ( to gain insights into how the applications are behaving in production and to identify any deviations from expected behavior)


<center><h2><b>Jenkins</b></h2></center>

**Install Jenkins on Ubuntu 18.04.6**

Please run the following command to install jenkins;

`chmod +x jenkins_installation.sh && ./jenkins_installation.sh`

**Jenkins Pipeline**

- Commit Code
- Build
- Test
- Release
- Ready to deploy to production

**Jenkins Architecture**
- Jenkin does not support multiple languages build. For that, you need to have multiple distributed jenkins servers.

**Jenkins Configuration**
- We need to configure an email address with jenkins (Manage Jenkins -> System ->  E-mail Notification) where it sends out the notifications for its process. To configure gmail account, create a app password under 2-step verification by following this [link](https://stackoverflow.com/questions/26594097/javamail-exception-javax-mail-authenticationfailedexception-534-5-7-9-applicatio/72592946#72592946) 

- We need to configure some tools. (Manage Jenkins -> Tools -> Git -> Path to Git executable = /usr/bin/git (comes from `which git` on linux) )

**Connect Github Repo with Jenkins Job**
- Create a new job -> Go to Source Code Management -> Add http url of repository you want to attach -> Add Git Credentials with git username and password.

- Go to build steps -> Execute shell -> Paste this command to run go program `/usr/local/go/bin/go run hello.go`

- Build the job. 

**[Buidling Jenkins Pipeline through Groovy Scripts](https://www.jenkins.io/doc/pipeline/tour/hello-world/)**

We can create a pipeline through groovy scripts on UI. 
- On Dashboard, New item -> Pipeline -> ok -> Pipeline (Pipeline script) -> paste the script there -> save. Build the pipeline. 

A basic groovy script is written to run docker container with golang image
```
/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.21.0-alpine3.18' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
```

Add jenkins user to docker group 

`sudo usermod -a -G docker jenkins`

**Creating Jenkinsfile and using it from GitHub respository to build pipeline**

- We just have to add Jenkinsfile to our repository and select *Pipeline script from SCM*. 

- Whenever we use tool, we need to configure the tool in Manage Jenkins.
    - Manage Jenkins -> Tool Configuration -> Select Go tool -> Add Go installation -> Save.

**Add Email Notifications for every Build**

- Configure email address and apppassword(gmail) in Manage -> Configure -> Email Notification. 
- Add a post build section in Jenkinsfile.

**Types of Pipelines**

*Declarative* - Does not depend on groovey scripting

*Scripted* - Based on groovey script

**DevOps Approach to create Pipeline**
- Always create a Jenkinsfile in your project.
- Create a job on Jenkins UI to pull git repository (containing Jenkinsfile) and run the Jenkinsfile.

**DevOps Interview Question**
- Difference between agile and devops (Agile focuses on small, incremental and rapid release of products with customer feedback informs of sprints and releasing/integrating the product every two weeks and devops is continuous integration and delivery through automated pipelines that we now have multiple releases of product a day because of devops)
- Benefits of DevOps - Continuous delivery of software, less complex problems to man age, early detection and correction of bugs. Faster delivery of features, improved communication among different teams.
- How to implement devops in a project? 
    - Stage 1 - Analyze the existing process and implementation of project, identify improvment areas and construct the roadmap for implementation.
    - Stage 2 - A proof of concept is done and shred with customer for the approval. After getting approval, actual implemntation can be carried out.
    - Stage 3 - Implement the devops phases step by step; version control -> integration -> testing -> deploying -> delivery -> monitoring.

- Difference between continuous delivery and continuous deployment?

- what is configuration management in DevOps?

- How continuous monitorying helps in maintaining entire architecture of the system?
    - Detecting, identifying and reporting any faults in entire infrtructure of the system.
    - Monitors the status of servers and applications.
    - Ensures that services are running properly.

- What is infrastructure as code?

- What is Jenkins.
- Key aspects of jenkins pipeline?
    - Pipeline
    - Node
    - Step
    - Stage

- Types of pieplines in jenkins (Declarative vs scripted)

- How we can copy jenkins from one server to another.

- Security mechanisms jenkins uses to authenticate users.

- restarting the jenkins 
    - (jenkins url)/restart  # forces to restart causing all builds to stop
    - (jenkins url)/safeRestart # let builds finish and then restarts
