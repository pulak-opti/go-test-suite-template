timeout=10
interval=1
url="http://localhost:8080/hello"

end_time=$((SECONDS + timeout))

while [ $SECONDS -lt $end_time ]; do
    if curl -X GET -Is "$url" | head -n 1 | grep "200 OK"; then
        echo "Server is ready!"
        exit 0
    fi
    sleep "$interval"
done

echo "Server did not become ready within the timeout."
exit 1