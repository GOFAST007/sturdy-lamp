test_secrets_k8s() {
	if [ "$(skip 'test_secrets_k8s')" ]; then
		echo "==> TEST SKIPPED: test_secrets_k8s tests"
		return
	fi

	set_verbosity

	echo "==> Checking for dependencies"
	check_dependencies juju

	file="${TEST_DIR}/test-secrets-k8s.log"

	bootstrap "test-secrets-k8s" "${file}"

	test_secrets
	test_secret_drain
	test_user_secrets
	test_user_secret_drain

	# Takes too long to tear down, so forcibly destroy it
	export KILL_CONTROLLER=true
	destroy_controller "test-secrets-k8s"
}
