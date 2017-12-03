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

All commands accept `[image]` as positional parameter, in the form of:

    somedockerregistry.com/repository/image_name

or

    repository/image_name

In the second case `registry.hub.docker.com` is chosen as default (the same as with commands triggered for docker daemon).

---

Layers:

    ./docker_download layers quay.io/coreos/etcd --tag latest
    Image coreos/etcd:latest exists at registry quay.io
    SchemaVersion: 1
    FsLayers
     0 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
     1 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
     2 {sha256:397fadf4d50046ddb5d6ba3cdce413681a40345e370eaa689efe63c263402565}
     3 {sha256:a723000a4a8e3e5d41f6c2519f84d75fe62d204733d3ae4cfa3be983d4a9787a}
     4 {sha256:553c2784b0f79337a9277b22532562b64896c57c5658ef5d34b9c8c7e011b3fa}
     5 {sha256:69c64fa48ad985954ef8891d4a372625c5dcc5b6c7cb6f949bd891f95a5ab501}
     6 {sha256:1017c4c6e41e617f6bbfdd005a932f2ffa4a676a6d9958cfd5839671c49c793d}
     7 {sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4}
     8 {sha256:b56ae66c29370df48e7377c8f9baa744a3958058a766793f821dadcb144a4647}
    History
     0 06f1b00796d10fd42f13161f53bf5c3e0d6c65266ec1a7d9b08f9c16bf83b678 /bin/sh -c #(nop)  CMD ["/usr/local/bin/etcd"]
     1 281b6f8340feb5624c11b49c54684577405917c0e78fbb1bb93f9e303ca01029 /bin/sh -c #(nop)  EXPOSE 2379/tcp 2380/tcp
     2 2c12ad11ad1de8746250edfdee95f54754813ec0f263b12256b273957659d2a2 /bin/sh -c echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf
     3 ec2325819c7d7293b973133fc91fbad90e2f051f90577e9a2c016ce53edfb81a /bin/sh -c mkdir -p /var/lib/etcd/
     4 90176d21a318506c7147ee1f979d84f933a8dfde63bfda926a1e44bf33aa6120 /bin/sh -c mkdir -p /var/etcd/
     5 ddcd78fbf54aa751b66b2f3541da527196441d689843952bbfa7ff7011c0a2de /bin/sh -c #(nop) ADD file:fe89be5cd4104c14ddcb89f5ce88ccf58aa7a6b21f624b0b8c6e8978cb897f11 in /usr/local/bin/
     6 05624af864f6b3192fb41040e72aefd2436845b214611d45b35ae527d855196f /bin/sh -c #(nop) ADD file:c0e1301982cd75f701b9265f97d3f45f5631edcf1f8eff0ca1a229b06df20051 in /usr/local/bin/
     7 448b569aff2a8160ad9bb22fd3d85437d4951928897e3a68a706fe9650ef7fbf /bin/sh -c #(nop)  CMD ["/bin/sh"]
     8 b5815a31a59b66c909dbf6c670de78690d4b52649b8e283fc2bfd2594f61cca3 /bin/sh -c #(nop) ADD file:1e87ff33d1b6765b793888cd50e01b2bd0dfe152b7dbb4048008bfc2658faea7 in /

 ---

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
