# K8s
Golang

# Why use K8s? Decentralization.

因為 IT 是由新技術驅動的行業。運端技術讓 Docker 容器化技術讓單機走向叢集，K8s 幫助推進分散式系統的管理，無論系統是採用企業內部執行抑或是被託管在雲上。

K8s 幫助企業精簡不必要的人力，部署和維護 K8s 的維運工程師（不用 7 天 * 24 小時） + 注重於服務元件的系統工程師 + 開發業務程式的開發工程師。

# Why use K8s? MicroService.

微服務架構能將巨大的單體應用程式分解為很多很小相互連結的微服務，一個微服務背後可能有數個實例抄本支撐，抄本 copy 的數量隨系統負擔變化可進行調整，embeded 的 LB 在此發揮作用，對於系統升級或擴充的快速反饋，並且仍能保持穩定性很有幫助。

K8s 能包裝這些微服務，並將系統隨時隨地的搬遷 (GCE 公有雲或是 OpenStack 私有雲)，做更為彈性地管理。

# shift System without modifing Config File

服務的叢集 IP 無須維運人員改變執行期的設定檔案（免除修改程式碼的麻煩），即可將系統從實體機遷移至雲端。亦可在執行高峰時期，將部分的服務對應的 Pod Copy 分流方式地放入雲端，不但節省企業的實體機成本投入，也能不中斷不影響系統運作下作搬遷動作，不會影響用戶體驗。

像是交通、口罩、餐廳預定、電商、外賣外送，都會讓機構或企業租用雲端進行分流管理（流量控制）。

# Cluster

叢集中的 Node 數量可以從數台擴展成數百個的規模，因此系統能應付大量用戶同時存取系統的壓力。


# 系統部署架構圖

實現讀寫分離。




                     post|update|delete
               Write ------------------ API/ Back-End --------------- Master Node

                                                                   ｜

                                                                Sync 同步

                                                                   ｜
                            get
                    Read  ------ Front-End ------------------- [Slave Node 1, ...]
                    
                    
                    
# Linux 系統

建議安裝 VirtualBox 或是 VMware 在本機上虛擬一個 CentOS7 的虛擬機，虛擬機有自己的網路模式連到外部。
關閉虛擬機內的防火牆，並且利用 yum 安裝 etcd 和 kubernets 軟體。修改設定檔案後，啟動服務！
      
