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

Please run the follwoing command to install jenkins;

`chmod +x jenkins_installation.sh && ./jenkins_installation.sh`
