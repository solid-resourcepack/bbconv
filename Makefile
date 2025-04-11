
.PHONY: test-plugin
test-plugin:
	cd plugin && ./gradlew shadowJar && cp build/libs/plugin-0.1.0-all.jar ../.run/plugins/bbconv.jar

.PHONY: run-server
run-server:
	cd .run && java -jar paper-1.21.4-222.jar --nogui