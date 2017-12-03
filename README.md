# docker_download

To build:

    go get ./... && go build

How to use:

    Usage:
      docker_download [command]

    Available Commands:
      help        Help about any command
      layers      Get layers info about specified image
      pull        Downloads layers from specified image

    Flags:
      -h, --help           help for docker_download
          --tag string     tag of the image to get info/imageimage on from docker registry

    Use "docker_download [command] --help" for more information about a command.


Layers:

    ./docker_download layers mysql/mysql-server --tag=5.6.23

    SchemaVersion: 1
    FsLayers
    0 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    1 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    2 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    3 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    4 {sha256:c45bb0e60064f776fdbee153e8c57ee9781273427dfc045a9e4f26fe230a33bb}
    5 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    6 {sha256:adbc50742107a7745478e1840f45670164826d8c6c1e9f50e2e6e645d8386a11}
    7 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    8 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    9 {sha256:e15c386793c2cc0229c9bc91595b248f2702f633d72213caedb8b119019508c4}
    10 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    11 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
    History
    0 e8ce1f579b38b18db838c28640ef392f9296cdd5776407af62c6431f48c074ab /bin/sh -c #(nop) CMD [mysqld]
    1 e8ce1f579b38b18db838c28640ef392f9296cdd5776407af62c6431f48c074ab /bin/sh -c #(nop) CMD [mysqld]
    2 3b72b1ace3bfec667ef0f5f1724c5e485556cf99d6a8737692c993bbf57b92bc /bin/sh -c #(nop) EXPOSE 3306/tcp
    3 7b3714e277db62cf4d470a8c6347bc372e3f4019693444a206d3a795b96d5862 /bin/sh -c #(nop) ENTRYPOINT [/entrypoint.sh]
    4 e77d6faa1eb7ac12f5e51f2806dc31aecd16e8f352ed901d800ad888489e5176 /bin/sh -c #(nop) COPY file:dc3f554223e0005f7d0c32810729a6272fa63c6eaa0fdbafa726e5e6f920aaf6 in /entrypoint.sh
    5 d384db67743134cef842e4435bc992f67a0f371914ba828d7ad34c9ceaa57c0d /bin/sh -c #(nop) VOLUME [/var/lib/mysql]
    6 763881d9d0dfb554f8121cefbb8f5a48750886930e231de638503ebf90db8705 /bin/sh -c rpmkeys --import http://repo.mysql.com/RPM-GPG-KEY-mysql && yum install -y $PACKAGE_URL   && rm -rf /var/cache/yum/*
    7 fa93025f4880175c972c3815d91db8b8f829d1c7b8b3e91d6763f914ae0aaa0b /bin/sh -c #(nop) ENV PACKAGE_URL=https://repo.mysql.com/yum/mysql-5.6-community/docker/x86_64/mysql-community-server-minimal-5.6.24-2.el7.x86_64.rpm
    8 c94023a2b9196d6d07566f496f1cb6c735be16b0012e0da7afd3937ec19d800e /bin/sh -c #(nop) CMD [/bin/bash]
    9 8f0a27825a9abd867f33b50e6583827d9727fcfa5d46ff5fda36c4c92dcfe2b9 /bin/sh -c #(nop) ADD file:919fbca9692a1c56f625d30a09e2727549c53a8a14859cca5b706b266cca9e46 in /
    10 ad98bd7101f267fed07cd488f28ec82e2059a492933d32679c3b7320479ecef0 /bin/sh -c #(nop) MAINTAINER Oracle Linux Product Team <ol-ovm-info_ww@oracle.com>
    11 511136ea3c5a64f264b78b5433614aec563103b4d4702f3ba7d4d2698e22c158

Pull:

     ./docker_download pull mysql/mysql-server --tag 5.6.23
     Downloading https://registry.hub.docker.com/v2/mysql/mysql-server/blobs/sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4...
     Layer sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4 with size 32B
     Downloading https://registry.hub.docker.com/v2/mysql/mysql-server/blobs/sha256:c45bb0e60064f776fdbee153e8c57ee9781273427dfc045a9e4f26fe230a33bb...
     Layer sha256:c45bb0e60064f776fdbee153e8c57ee9781273427dfc045a9e4f26fe230a33bb with size 1300B
     Downloading https://registry.hub.docker.com/v2/mysql/mysql-server/blobs/sha256:adbc50742107a7745478e1840f45670164826d8c6c1e9f50e2e6e645d8386a11...
     Layer sha256:adbc50742107a7745478e1840f45670164826d8c6c1e9f50e2e6e645d8386a11 with size 37027670B
     Downloading https://registry.hub.docker.com/v2/mysql/mysql-server/blobs/sha256:e15c386793c2cc0229c9bc91595b248f2702f633d72213caedb8b119019508c4...
     Layer sha256:e15c386793c2cc0229c9bc91595b248f2702f633d72213caedb8b119019508c4 with size 70978368B
