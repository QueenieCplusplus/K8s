// 2020, 7/25, pm 2:30
// post-comment.php


<?
    require 'Predis/Autoloader.php';
    
    Predis\Autoloader::register();

    if(isset($_GET['cmd']) === true){

        
        $host = 'master';
        header('');

          // to set or write to master;
        if($_GET['cmd'] == 'set'){

            $host = 'host';

            header('Content-Type: application/json');

            if(getenv('GET_HOSTS_FROM'=='env')){

                $host = getenv('MASTER_SERVICE_HOST');
                
            }

            $client = new Predis\Client([
                'schema' => 'tcp',
                'host' => $host,
                'port' => 6379

            ]);

            $client->set($_GET['key'], $_GET['value']);
            print('{"message": "update"}');

        }else{

            // to get or read from slave;
            $hostS = 'slave';

            if(getenv('GET_HOSTS_FROM'=='env'){

                $hostS = 'SLAVE_SERVICE_HOST'

            }

            $clientS = new Predis\Client([
                'schema' => 'tcp',
                'host' => $hostS,
                'port' => 6379

            ]);

            $v = $clientS->get($_GET['key']);
            print(' { "show data" : " ' . $v . ' "}');

    }else{
        phpinfo();
    }
?>


<!-- 建立 rc/backend for this php -->
<!-- $kubectl create -f backend-controller.yaml -->