<?php
//命名空间
namespace Helloworld;

//定义PHP客户端
class GreeterClient extends \gRPC\BaseStub
{

  //定义构造方法
  public function __construct($hostname, $opts, $channel = null)
  {
    parent::__construct($hostname, $opts, $channel);
  }

  /**
   * 实现proto文件中定义的SayHello()方法
   * Sends a greeting
   * @param \Helloworld\HelloRequest $argument input argument
   * @param array $metadata metadata
   * @param array $options call options
   * @return \gRPC\UnaryCall
   */
  public function SayHello(\Helloworld\HelloRequest $argument,
                           $metadata = [], $options = [])
  {
    return $this->_simpleRequest('/helloworld.Greeter/SayHello',
      $argument,
      ['\Helloworld\HelloReply', 'decode'],
      $metadata, $options);
  }

}
