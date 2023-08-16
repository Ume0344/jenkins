#!/bin/bash


function install_java {
    # Install java
    sudo apt update -y
    sudo apt install openjdk-17-jre
}


function install_jenkins {
    # Install Jenkins

    if [ -n `which java` ]
    then
        echo "Java is installed, Installing Jenkins"
        curl -fsSL https://pkg.jenkins.io/debian/jenkins.io-2023.key | sudo tee \
        /usr/share/keyrings/jenkins-keyring.asc > /dev/null
        echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
        https://pkg.jenkins.io/debian binary/ | sudo tee \
        /etc/apt/sources.list.d/jenkins.list > /dev/null

        sudo apt-get update -y
        sudo apt-get install jenkins

    else
        echo "Java is not installed, Installing Java"
        install_java
        install_jenkins
    fi
}

install_java
install_jenkins
