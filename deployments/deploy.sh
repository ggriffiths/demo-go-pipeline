kubectl apply -f zookeeper.yml -n grant
sleep 10

kubectl apply -f kafka.yml -n grant
sleep 10

kubectl apply -f pipeline.yml -n grant