apply:
	kubectl apply -f ./ns
	kubectl apply -f ./netpol -f ./app -n ns1
	kubectl apply -f ./netpol -f ./app -n ns2
	kubectl apply -f ./ing

delete:
	kubectl delete -f ./ns
