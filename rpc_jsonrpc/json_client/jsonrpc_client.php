<?php


class JsonRPC
{

    private $conn;

    function __construct($host, $port)
    {
        $this->conn = fsockopen($host, $port, $errno, $errstr, 3);
        if (!$this->conn) {
            return false;
        }
    }

    public function Call($method, $params)
    {
        if (!$this->conn) {
            return false;
        }
        $err = fwrite($this->conn, json_encode(array(
                'method' => $method,
                'params' => array($params),
                'id' => 0,
            )) . "\n");
        if ($err === false) {
            return false;
        }
        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            return NULL;
        }
        return json_decode($line, true);
    }
}

$client = new JsonRPC("127.0.0.1", 8096);
$args = array('A' => 9, 'B' => 2);
$r = $client->Call("Arith.Multiply", $args);
printf("%d * %d = %d\n", $args['A'], $args['B'], $r['result']['Pro']);
$r = $client->Call("Arith.Divide", array('A' => 9, 'B' => 2));
printf("%d / %d, Quo is %d, Rem is %d\n", $args['A'], $args['B'], $r['result']['Quo'], $r['result']['Rem']);
