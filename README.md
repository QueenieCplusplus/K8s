# K8s
Golang & Linux

# Why use K8s? Decentralization.

因為 IT 是由新技術驅動的行業。運端技術使 Docker 容器化技術讓單機走向叢集，K8s 幫助推進分散式系統的管理，無論系統是採用企業內部執行抑或是被託管在雲上。

K8s 幫助企業精簡不必要的人力，部署和維護 K8s 的維運工程師（不用 7 天 * 24 小時） + 注重於服務元件的系統工程師 + 開發業務程式的開發工程師。

# Why use K8s? MicroService.

微服務架構能將巨大的單體應用程式分解為很多很小相互連結的微服務，一個微服務背後可能有數個實例抄本支撐，抄本 copy 的數量隨系統負擔變化可進行調整，embeded 的 LB 在此發揮作用，對於系統升級或擴充的快速反饋，並且仍能保持穩定性很有幫助。

K8s 能包裝這些微服務，並將系統隨時隨地的搬遷 (GCE 公有雲或是 OpenStack 私有雲)，做更為彈性地管理。

# shift System without modifing Config File

服務的叢集 IP 無須維運人員改變執行期的設定檔案（免除修改程式碼的麻煩），即可將系統從實體機遷移至雲端。亦可在執行高峰時期，將部分的服務對應的 Pod Copy 分流方式地放入雲端，不但節省企業的實體機成本投入，也能不中斷不影響系統運作下作搬遷動作，不會影響用戶體驗。

像是交通、線上預訂口罩、餐廳預定、電商、外賣外送，都會讓機構或企業租用雲端進行分流管理（流量控制）。

# Cluster

叢集中的 Node 數量可以從數台擴展成數百個的規模，因此系統能應付大量用戶同時存取系統的壓力。
同 Pod 上的 containers 彼此溝通依靠 Link。

例如：

    $kubectl get services

    Name          IPs              PORTs         

    master        10.254.2.3      6379/TCP       與 slave 同 Port 不同 IP ，是指 Pod 不同，但服務相同
    slave         10.254.80.6     6379/TCP       (M/S pods 彼此靠 vSW 溝通)
    backend       10.254.170.5    80/TCP         與 nodes 們靠 port 300001 溝通，對外訪問則用 port 80，
                                                 同 IP 不同 port 為 bind（）繫結的應用。
                                                 // valid: spec.ports[0].nodePort: The range of valid ports is 30000-32767


# 系統部署架構圖

實現讀寫分離。
如下可能是同一台 Node 中的同一台 VM。（也可能是不同台實體機。）




                     post|update|delete
               Write ------------------ API/ Back-End ---------- Master Pod

                                                                   ｜

                                                                Sync 同步

                                                                   ｜
                            get
                    Read  ------ Front-End ------------------- [Slave Pod 1, ...]
                    
                    
# M/S 主從資料同步

為了使 slave pod 知道 master pod的 位址，可將 slave image 啟動命令 /run.sh 中寫入：

    // 6379 port
    server --slaveof ${MASTER_SERVICE_HOST} 6379 
    
建立 slave pod 時，系統自動在 container 中生成與 master 相關的環境變數如取得 IP 位置。
                                 
                    
# Linux 系統

建議安裝 VirtualBox 或是 VMware 在本機上虛擬一個 CentOS7 的虛擬機，虛擬機有自己的網路模式連到外部。
關閉虛擬機內的防火牆，並且利用 yum 安裝 etcd 和 kubernets 軟體。修改設定檔案後，啟動服務！
      

# Activate Serivices

       $systemctl start etcd
       $systemctl start docker
       $systemctl start kube-apiserver
       #systemctl start kube-controller-manager
       $systemctl start kube-scheduler
       $systemctl start kubelet
       $systemctl start kube-proxy

# Backend

see comment-poster.php
its backend-service.yaml includes:

    type: NodePort
    ports:
    -port: 80 // Node上的物理機阜提供對外訪問能力
     nodePort: 30001 // 預設為 30000 ~ 32767 為特定範圍
     
# Pod 組態檔案

路徑：/etc/kubernetes/manifests/kube-apiserver.yaml 

https://qhh.me/2019/08/15/Kubernetes-调整-nodePort-端口范围/

 




