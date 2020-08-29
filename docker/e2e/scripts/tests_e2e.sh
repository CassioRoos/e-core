exit_status=$?
sleep 90
for i in `seq 1 100`; do
    echo "Waiting for system to start up. Healthcheck #$i"
    sleep 1
    curl -s --fail "http://$TEST_SERVER_HOST$PORT/healthcheck" > /dev/null
    exit_status=$?
    if [ $exit_status -eq 0 ]; then
        echo "Server is healthy. Starting tests"
        break
    fi
done

if [ $exit_status -ne 0 ]; then
        echo "Server not healthy after 100 attempts"
        exit 1
fi

if [ $exit_status -eq 0 ]; then
    TEST_ENV=local go test -mod=vendor -v -race -failfast ./tests
    ret=$?
fi
exit $ret