# DevOps For The Desperate

ECK:
https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-deploy-eck.html

My Progress:
Part I
- [X] Chapter 1
- [X] Chapter 2
- [X] Chapter 3
- [X] Chapter 4
- [X] Chapter 5
    - IP addresses:
        10.0.2.15/24 fe80::d2:a9ff:fe00:7a16/64 
        192.168.56.4/24 fe80::a00:27ff:fe1b:89bd/64
Part II
- [X] Chapter 6
    - recover the minikube/kubernetes environment from sudden shutdown:
        ```
        minikube start --drver=virtualbox
        minikube kubectl -- get services
        # In its own terminal:
        minikube tunnel
        ```
    - Minikube installation instructions:
        https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fdebian+package
        *As root:*
        ```
        # Go to https://github.com/kubernetes/minikube/releases/
        #  find the version (such as the most recent) you want
        #  download the .deb file, and install like so:
        apt install ./minikube_<release/version.number>_amd64.deb
        ```
        # ```
        # curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
        # apt install ./minikube_latest_admd64.deb
        # ```
    - Install Docker:
        https://docs.docker.com/engine/install/debian/
        *As root:*
        ```
        # Add Docker's official GPG key:
        apt-get update
        apt-get install ca-certificates curl
        install -m 0755 -d /etc/apt/keyrings
        curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
        chmod a+r /etc/apt/keyrings/docker.asc

        # Add the repository to Apt sources:
        echo \
          "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
          $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
          sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
        apt-get update

        # Actually install the latest version:
        apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
        ```
    - docker build warnings:
        ```
         3 warnings found (use docker --debug to expand):
         - FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 7)
            - Literally 'as' -> 'AS' (it does have better readability)
         - LegacyKeyValueFormat: "ENV key=value" should be used instead of legacy "ENV key value" format (line 9)
         - LegacyKeyValueFormat: "ENV key=value" should be used instead of legacy "ENV key value" format (line 10)
        ```
- [X] Chapter 7
- [] Chapter 8
    - Install Skaffold:
        - Google's instructions for installing their own product is wrong, there is no publically available 'release' branch.
            GOOGLE SUCKS
        - https://github.com/GoogleContainerTools/skaffold/releases/tag/v2.13.2
        ```
        # As root:
        curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v2.13.2/skaffold-linux-amd64 && \
        chmod +x skaffold && \
        mv skaffold /usr/local/bin
        ```
    - GOOGLE majorly sucks
        The first Skaffold error was that the URL provided on: https://skaffold.dev/docs/install/ was returning an error (up until today)
        skaffold was causing a reference error in docker by converting the "/" to an "_" in the image tag
        ```
        Generating tags...
         - dftd/telnet-server -> ./dftd_telnet-server:0c778a6
        ```
        After changing the api version in the skaffold.yaml file, it gave me the correct URL for getting the latest release.
          This gave an even worse error. So, I restored the api version, and the same error persisted.

    - Install *container-structure-test*:
        - https://github.com/GoogleContainerTools/container-structure-test
        ```
        curl -LO https://github.com/GoogleContainerTools/container-structure-test/releases/latest/download/container-structure-test-linux-amd64 && \
        chmod +x container-structure-test-linux-amd64 && \
        sudo mv container-structure-test-linux-amd64 /usr/local/bin/container-structure-test
        ```
Part III
- [] Chapter 9
- [] Chapter 10

![dftd](book-cover.png "Book front cover")

This repo is for the book [DevOps for the Desperate](https://nostarch.com/devops-desperate).

## Directories

Each one of these directories contain different material for different parts in the book.
Here is a quick overview of each directory:

* _ansible_: Contains the playbook and tasks to follow along in the first section of the book.

* _monitoring_: Contains the k8s manifest files to install Prometheus, Alertmanager, and Grafana.

* _runbooks_: Contains a simple runbook for `telnet-server` and examples of a runbook broken up by alert. These are used in the alerts for Chapter 9.

* _telnet-server_: Contains the sample application that is used throughout two-thirds of the book.

* _vagrant_: Contains the Vagrantfile for the VM used in the first section of the book.

* _apple-silicon_: Contains the Vagrantfile for the VM used in the first section of the book and minikube instructions.
