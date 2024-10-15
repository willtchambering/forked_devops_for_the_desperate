# DevOps For The Desperate

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
- [] Chapter 6
    - Minikube installation instructions:
        https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fdebian+package
        *As root:*
        ```
        curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
        apt install ./minikube_latest_admd64.deb
        ```
- [] Chapter 7
- [] Chapter 8
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
