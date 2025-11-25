docker run -d --name rabbitmq   -p 5672:5672   -p 15672:15672   -v rabbitmq_data:/var/lib/rabbitmq rabbitmq:4-management

dmujuzi@dmujuzi:~$ history | grep "rabbitmqctl"
  318  sudo rabbitmqctl list_queues
  319  docker exec -it rabbitmq rabbitmqctl list_queues name messages_ready messages_unacknowledged
  321  docker exec -it rabbitmq rabbitmqctl list_exchanges
  322  docker exec -it rabbitmq rabbitmqctl list_bindings
  863  history | grep "rabbitmqctl"
dmujuzi@dmujuzi:~$ 
