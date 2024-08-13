while true; do
    timeout -k 20m 20m kubectl logs -n riab rimedo-ts-7bf45bc9b7-4b67q -c rimedo-ts -f | tee -a logs.txt 
    sleep 5
done